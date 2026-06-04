// Package core registers the server and all plugins we support.
package core

import (
	// plug in the server
	_ "github.com/mr-torgue/coredns/core/dnsserver"
)
