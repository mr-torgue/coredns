package request

import (
	"github.com/mr-torgue/coredns/plugin/pkg/edns"

	"github.com/mr-torgue/dns"
)

func supportedOptions(o []dns.EDNS0) []dns.EDNS0 {
	supported := make([]dns.EDNS0, 0, 3)
	for _, opt := range o {
		if edns.SupportedOption(opt.Option()) {
			supported = append(supported, opt)
		}
	}
	return supported
}
