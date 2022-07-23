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

func BenchmarkUUID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.New()
	}
}

func BenchmarkPoolUUID(b *testing.B) {
	uuid.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.New()
	}
}

func BenchmarkXID(b *testing.B) {
	uuid.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xid.New()
	}
}

func BenchmarkULID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ulid.Make()
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
		{"WithCrypoEntropy", crand.Reader},
		{"WithEntropy", rng},
		{"WithMonotonicEntropy_Inc0", ulid.Monotonic(rng, 0)},
		{"WithMonotonicEntropy_Inc1", ulid.Monotonic(rng, 1)},
		{"WithCryptoMonotonicEntropy_Inc0", ulid.Monotonic(crand.Reader, 0)},
		{"WithCryptoMonotonicEntropy_Inc1", ulid.Monotonic(crand.Reader, 1)},
		{"WithCryptoThreadSafe", CryptoEntropy()},
		{"WithoutEntropy", nil},
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
		_, _ = ulid.New(timestamp, entropy)
	})
}

func BenchmarkUUIDString(b *testing.B) {
	uuid.DisableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.NewString()
	}
}

func BenchmarkPoolUUIDString(b *testing.B) {
	uuid.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uuid.NewString()
	}
}

func BenchmarkPoolXIDString(b *testing.B) {
	uuid.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xid.New().String()
	}
}

func BenchmarkPoolULIDString(b *testing.B) {
	uuid.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ulid.Make().String()
	}
}

func BenchmarkParallellUUIDString(b *testing.B) {
	uuid.DisableRandPool()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = uuid.New().String()
		}
	})
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

func BenchmarkParallelULIDString(b *testing.B) {
	entropy := ulid.DefaultEntropy()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		ts := ulid.Now()
		for pb.Next() {
			_ = ulid.MustNew(ts, entropy).String()
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
		ts := ulid.Now()
		for pb.Next() {
			_ = ulid.MustNew(ts, entropy).String()
		}
	})
}
