package whoami

import (
	"github.com/coredns/caddy"
	"github.com/mr-torgue/coredns/core/dnsserver"
	"github.com/mr-torgue/coredns/plugin"
)

func init() { plugin.Register("whoami", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'whoami'
	if c.NextArg() {
		return plugin.Error("whoami", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(_next plugin.Handler) plugin.Handler {
		return Whoami{}
	})

	return nil
}
