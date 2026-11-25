// Package replaycache implements a bounded, TTL-windowed duplicate detector
// for QUIC/TLS 0-RTT early data (RFC 9250 §4.5; RFC 8446 §8), built on
// plugin/pkg/cache.
//
// It is deliberately best-effort: quic-go does not expose the PSK/session
// identity of a resumed connection, so nothing above the TLS stack can
// implement strict single-use-ticket anti-replay. This cache detects the
// practically relevant vector — captured early data being re-sent (spoofed
// from the original client's address, or re-sent by the same client) — and
// flags duplicates so they can be dropped and counted.
package replaycache

import (
	"encoding/binary"
	"sync"
	"time"

	"github.com/coredns/coredns/plugin/pkg/cache"
)

// Cache is a bounded set of recently seen early-data keys with a TTL window.
type Cache struct {
	mu    sync.Mutex
	items *cache.Cache[time.Time] // key -> expiry
	ttl   time.Duration
}

// New returns a Cache that flags duplicate keys within ttl and holds at most
// size entries.
func New(ttl time.Duration, size int) *Cache {
	return &Cache{items: cache.New[time.Time](size), ttl: ttl}
}

// Key hashes the parts identifying one unit of early data. Each part is
// length-prefixed so different splits of the same bytes hash differently.
// Include the client address: identical queries from *different* clients are
// then not flagged, while replays of captured early data — which reuse the
// original client's address — are detected.
func Key(parts ...[]byte) uint64 {
	buf := make([]byte, 0, 64)
	for _, p := range parts {
		buf = binary.AppendUvarint(buf, uint64(len(p)))
		buf = append(buf, p...)
	}
	return cache.Hash(buf)
}

// Add reports whether key is fresh (true, may process) or a suspected replay
// within the TTL window (false). Safe for concurrent use.
func (c *Cache) Add(key uint64, now time.Time) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if exp, ok := c.items.Get(key); ok && now.Before(exp) {
		return false // duplicate within window
	}
	c.items.Add(key, now.Add(c.ttl))
	return true
}

// Len returns the number of keys currently held.
func (c *Cache) Len() int { return c.items.Len() }
