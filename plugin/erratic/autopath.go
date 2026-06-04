package erratic

import "github.com/mr-torgue/coredns/request"

// AutoPath implements the AutoPathFunc call from the autopath plugin.
func (e *Erratic) AutoPath(_state request.Request) []string {
	return []string{"a.example.org.", "b.example.org.", ""}
}
