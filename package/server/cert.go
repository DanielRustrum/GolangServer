package server

import (
	"crypto/tls"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

// TODO: Cert Cacheing
// TODO: DNS Cert from Lets Encrypt; use Autocert

func getDNSCert() func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
	domainCert := genDomainString()
	domainsList := strings.Fields(domainCert)
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domainsList...),
	}
	return certManager.GetCertificate
}

func getLocalCert() (string, string) {

	return "", ""
}
