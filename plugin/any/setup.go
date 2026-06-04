package any

import (
	"github.com/coredns/caddy"
	"github.com/mr-torgue/coredns/core/dnsserver"
	"github.com/mr-torgue/coredns/plugin"
)

func init() { plugin.Register("any", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'any'
	if c.NextArg() {
		return plugin.Error("any", c.ArgErr())
	}
	if c.NextBlock() {
		return plugin.Error("any", c.Errf("unknown property '%s'", c.Val()))
	}

	a := Any{}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		a.Next = next
		return a
	})

	return nil
}
