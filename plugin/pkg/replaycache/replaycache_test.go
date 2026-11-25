package replaycache

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestKeyLengthPrefix(t *testing.T) {
	if Key([]byte("ab"), []byte("c")) == Key([]byte("a"), []byte("bc")) {
		t.Error("different splits of the same bytes must hash differently")
	}
	if Key([]byte("query")) != Key([]byte("query")) {
		t.Error("identical parts must hash identically")
	}
}

func TestReplayCacheDuplicate(t *testing.T) {
	c := New(time.Minute, 1024)
	k := Key([]byte("query"))
	now := time.Now()
	if !c.Add(k, now) {
		t.Fatal("first Add must be fresh")
	}
	if c.Add(k, now.Add(time.Second)) {
		t.Fatal("second Add within window must be flagged as replay")
	}
}

func TestReplayCacheExpiry(t *testing.T) {
	c := New(time.Minute, 1024)
	k := Key([]byte("query"))
	now := time.Now()
	c.Add(k, now)
	if c.Add(k, now.Add(time.Second)) {
		t.Error("replay within TTL must be flagged")
	}
	if !c.Add(k, now.Add(2*time.Minute)) {
		t.Error("same key after TTL expiry must be fresh again")
	}
}

func TestReplayCacheDifferentClientsNotFlagged(t *testing.T) {
	c := New(time.Minute, 1024)
	now := time.Now()
	q := []byte("identical query bytes")
	c.Add(Key(q, []byte("192.0.2.1:53")), now)
	if !c.Add(Key(q, []byte("192.0.2.2:53")), now) {
		t.Error("identical query from a different client must not be flagged")
	}
}

func TestReplayCacheCapacity(t *testing.T) {
	const size = 1024
	c := New(time.Hour, size)
	now := time.Now()
	for i := 0; i < 100*size; i++ {
		c.Add(Key([]byte{byte(i), byte(i >> 8), byte(i >> 16)}), now)
	}
	if got := c.Len(); got > size {
		t.Errorf("cache exceeded capacity: %d > %d", got, size)
	}
}

func TestReplayCacheConcurrent(t *testing.T) {
	c := New(time.Minute, 1024)
	k := Key([]byte("query"))
	var fresh atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if c.Add(k, time.Now()) {
				fresh.Add(1)
			}
		}()
	}
	wg.Wait()
	if fresh.Load() != 1 {
		t.Errorf("exactly one goroutine must see the key as fresh, got %d", fresh.Load())
	}
}
