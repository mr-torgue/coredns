//go:build gofuzz

package whoami

import (
	"github.com/mr-torgue/coredns/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	w := Whoami{}
	return fuzz.Do(w, data)
}
