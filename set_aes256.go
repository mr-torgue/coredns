//go:build FORCE_AES256

package main


import (
	"crypto/tls"
	_ "unsafe"
)

// BAD CODE WARNING!
// linking not recommmended but it is the easiest way to enforce the use of AES256.
// AES128 should be quantum-safe, but AES256 might be more secure.

//go:linkname defaultCipherSuitesTLS13 crypto/tls.defaultCipherSuitesTLS13
var defaultCipherSuitesTLS13 []uint16

//go:linkname defaultCipherSuitesTLS13NoAES crypto/tls.defaultCipherSuitesTLS13NoAES
var defaultCipherSuitesTLS13NoAES []uint16

defaultCipherSuitesTLS13 = []uint16{
	tls.TLS_AES_256_GCM_SHA384,
	//tls.TLS_AES_128_GCM_SHA256,
	tls.TLS_CHACHA20_POLY1305_SHA256,
}
defaultCipherSuitesTLS13NoAES = []uint16{
	tls.TLS_CHACHA20_POLY1305_SHA256,
	tls.TLS_AES_256_GCM_SHA384,
	//tls.TLS_AES_128_GCM_SHA256,
}


