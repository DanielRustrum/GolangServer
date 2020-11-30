package http

import (
	"crypto/tls"
	"errors"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

// TODO: Cert Caching
// TODO: Gen Cert Though OpenSSL Library
// TODO: Get CertificateFunc with OpenSSL Library instead of file
// TODO: DNS Cert from Lets Encrypt; use Autocert > Inprogress
// TODO: Automatically Gen Certs > Inprogress
// TODO: Gen Cert Through OpenSSL Command > InProgress

//* Local Cert Generation

func getLocalCert(providedDomains string) (CertificateFunc, error) {
	return nil, errors.New("Not Implemented")
}

//* Lets Encrypt Cert Generation

func getDNSCert(providedDomains string) (CertificateFunc, error) {
	domainsList := strings.Fields(providedDomains)
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domainsList...),
	}
	return certManager.GetCertificate, errors.New("Not Fully Implemented")
}

//* Public

// CertificateFunc is ...
type CertificateFunc func(*tls.ClientHelloInfo) (*tls.Certificate, error)

// GetCert is ...
func GetCert(host string, domains string) CertificateFunc {
	var cert CertificateFunc = nil
	var err error

	if host != "localhost" {
		cert, err = getDNSCert(domains)
	} else {
		cert, err = getLocalCert(domains)
	}

	if err != nil {
		return nil
	}

	return cert
}
