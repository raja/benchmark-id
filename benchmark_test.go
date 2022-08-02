package main

import (
	"io"
	"math/rand"
	"sync"
	"testing"
	"time"

	crand "crypto/rand"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
)

func BenchmarkPoolUUID(b *testing.B) {
	uuid.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uuid.New()
	}
}

func BenchmarkXIDNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xid.New()
	}
}

func BenchmarkULIDMake(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ulid.Make()
	}
}

var (
	entropy     io.Reader
	entropyOnce sync.Once
)

func CryptoEntropy() io.Reader {
	entropyOnce.Do(func() {
		crand := crand.Reader
		entropy = &ulid.LockedMonotonicReader{
			MonotonicReader: ulid.Monotonic(crand, 0),
		}
	})
	return entropy
}

func benchmarkMakeULID(b *testing.B, f func(uint64, io.Reader)) {
	b.ReportAllocs()
	b.SetBytes(int64(len(ulid.ULID{})))

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, tc := range []struct {
		name    string
		entropy io.Reader
	}{

		{"WithMonotonicEntropy", ulid.Monotonic(rng, 0)},
		{"WithCryptoMonotonicEntropy", ulid.Monotonic(crand.Reader, 0)},
		{"WithCryptoThreadSafe", CryptoEntropy()},
	} {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				f(ulid.Now(), tc.entropy)
			}
		})
	}
}

func BenchmarkULIDNew(b *testing.B) {
	benchmarkMakeULID(b, func(timestamp uint64, entropy io.Reader) {
		_ = ulid.MustNew(timestamp, entropy)
	})
}

func BenchmarkPoolUUIDString(b *testing.B) {
	uuid.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uuid.NewString()
	}
}

func BenchmarXIDString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xid.New().String()
	}
}

func BenchmarkULIDString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ulid.Make().String()
	}
}

func BenchmarkULIDCryptoString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ulid.MustNew(ulid.Now(), CryptoEntropy()).String()
	}
}

func BenchmarkParallelPoolUUIDString(b *testing.B) {
	uuid.EnableRandPool()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = uuid.New().String()
		}
	})
}

func BenchmarkParallelXIDString(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = xid.New().String()
		}
	})
}

func BenchmarkParallelULIDMakeString(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = ulid.Make().String()
		}
	})
}

func BenchmarkParallelULIDCryptoThreadSafeString(b *testing.B) {
	entropy := CryptoEntropy()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = ulid.MustNew(ulid.Now(), entropy).String()
		}
	})
}
