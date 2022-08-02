package main

/*
import (
	"testing"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
)

const maxCount = 100_000_000 // 100M

func collisionError(t *testing.T, name string, i int) {
	t.Errorf("[%s] Collision is detected on loop #[%d]", name, i)
}

func logSuccess(t *testing.T, name string) {
	t.Logf("[%s] No collision occured through #[%d] iterations", name, maxCount)
}

func exampleOutput(t *testing.T, name string, m map[string]struct{}) {
	const size = 10
	list := make([]string, 0, size)
	count := 0
	for key := range m {
		if count >= size {
			break
		}
		list = append(list, key)
		count++
	}
	t.Logf("[%s] Example output => %#v", name, list)
}

func TestCollisionUUID(t *testing.T) {
	name := "uuid.NewString()"
	uuid.EnableRandPool()
	m := make(map[string]struct{}, maxCount)
	for i := 0; i < maxCount; i++ {
		v := uuid.New().String()
		if _, ok := m[v]; ok {
			collisionError(t, name, i)
			return
		}
		m[v] = struct{}{}
	}
	logSuccess(t, name)
	exampleOutput(t, name, m)
	m = nil
}

func TestCollisionXID(t *testing.T) {
	name := "xid.New().String()"
	m := make(map[string]struct{}, maxCount)
	for i := 0; i < maxCount; i++ {
		v := xid.New().String()
		if _, ok := m[v]; ok {
			collisionError(t, name, i)
			return
		}
		m[v] = struct{}{}
	}
	logSuccess(t, name)
	exampleOutput(t, name, m)
	m = nil
}

func TestCollisionULID(t *testing.T) {
	entropy := CryptoEntropy()
	name := "ulidCryptoEntropy"
	m := make(map[string]struct{}, maxCount)
	for i := 0; i < maxCount; i++ {
		v := ulid.MustNew(ulid.Now(), entropy).String()
		if _, ok := m[v]; ok {
			collisionError(t, name, i)
			return
		}
		m[v] = struct{}{}
	}
	logSuccess(t, name)
	exampleOutput(t, name, m)
	m = nil
}

/*func TestCollisionULIDForceCollision(t *testing.T) {
	name := "ulidNoEntropy"
	m := make(map[string]struct{}, maxCount)
	for i := 0; i < maxCount; i++ {
		v := ulid.MustNew(ulid.Now(), nil).String()
		if _, ok := m[v]; ok {
			collisionError(t, name, i)
			return
		}
		m[v] = struct{}{}
	}
	logSuccess(t, name)
	exampleOutput(t, name, m)
	m = nil
}*/
