package view

import (
	"context"

	"github.com/mr-torgue/coredns/plugin/metadata"
	"github.com/mr-torgue/coredns/request"
)

// Metadata implements the metadata.Provider interface.
func (v *View) Metadata(ctx context.Context, _state request.Request) context.Context {
	metadata.SetValueFunc(ctx, "view/name", func() string {
		return v.viewName
	})
	return ctx
}
