package main

//go:generate go run directives_generate.go
//go:generate go run owners_generate.go

import (
	"crypto/tls"

	_ "unsafe"

	_ "github.com/mr-torgue/coredns/core/plugin" // Plug in CoreDNS.
	"github.com/mr-torgue/coredns/coremain"
)

// BAD CODE WARNING!
// linking not recommmended but it is the easiest way to enforce the use of AES256.
// AES128 should be quantum-safe, but AES256 might be more secure.

//go:linkname defaultCipherSuitesTLS13 crypto/tls.defaultCipherSuitesTLS13
var defaultCipherSuitesTLS13 = []uint16{
	tls.TLS_AES_256_GCM_SHA384,
	//tls.TLS_AES_128_GCM_SHA256,
	tls.TLS_CHACHA20_POLY1305_SHA256,
}

//go:linkname defaultCipherSuitesTLS13NoAES crypto/tls.defaultCipherSuitesTLS13NoAES
var defaultCipherSuitesTLS13NoAES = []uint16{
	tls.TLS_CHACHA20_POLY1305_SHA256,
	tls.TLS_AES_256_GCM_SHA384,
	//tls.TLS_AES_128_GCM_SHA256,
}

func main() {
	coremain.Run()
}
