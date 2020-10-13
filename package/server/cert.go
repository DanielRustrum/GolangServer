package server

import (
	"crypto/tls"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"golang.org/x/crypto/acme/autocert"
)

// TODO: Cert Caching
// TODO: Gen Cert Though OpenSSL Library
// TODO: Get CertificateFunc with OpenSSL Library instead of file
// TODO: DNS Cert from Lets Encrypt; use Autocert > Inprogress
// TODO: Automatically Gen Certs > Inprogress
// TODO: Gen Cert Through OpenSSL Command > InProgress

//* OpenSSL Functions

func downloadOpenSSLSource(privateDir string) error {
	return errors.New("Not Implemented")
}

func getOpenSSLLib() (OpenSSLLib, error) {
	return OpenSSLLib{}, errors.New("Not Implemented")
}

func getOpenSSLCertThroughCommand(openSSLCommand string, privateDir string, providedDomains string) error {
	runCommands := func(commandStrings map[string]string) error {
		for message, commandString := range commandStrings {
			commandArgs := strings.Fields(commandString)
			command := exec.Command(openSSLCommand, commandArgs...)
			fmt.Printf(message + "\n")
			err := command.Run()
			if err != nil {
				return err
			}
		}
		return nil
	}

	commandStrings := make(map[string]string)

	// * Key = Message, Value = Command
	// TODO: Create Commands
	commandStrings["Generating Private Key"] = ""

	commandStrings["Generating Public Key"] = ""

	commandStrings["Generating CSR"] = ""

	commandStrings["Generating Cert"] = ""

	commandStrings["Signing Public Key"] = ""

	return runCommands(commandStrings)
}

//* Cert Functions

func getDNSCert(providedDomains string) (CertificateFunc, error) {
	domainsList := strings.Fields(providedDomains)
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domainsList...),
	}
	return certManager.GetCertificate, errors.New("Not Fully Implemented")
}

func getLocalCert(providedDomains string) (CertificateFunc, error) {
	//openSSLlib, err := FetchOpenSSLLib(privateDir)
	return nil, errors.New("Not Implemented")
}

//* Public

// OpenSSLLib is ...
type OpenSSLLib struct{}

// CertificateFunc is ...
type CertificateFunc func(*tls.ClientHelloInfo) (*tls.Certificate, error)

// GetCert is ...
func GetCert(openSSLCommand string, privateDir string, host string, domains string) (string, string, CertificateFunc) {
	var certFunc CertificateFunc = nil
	var cert, key string = "", ""

	if openSSLCommand == "" {
		var err error

		if host != "localhost" {
			certFunc, err = getDNSCert(domains)
		} else {
			certFunc, err = getLocalCert(domains)
		}

		if err != nil {
			//! Temporary Panic Until non-openSSLCommand Method is Implmented
			panic("Function Not Implemented without OpenSSLCommand Yet")
		}

	} else {
		err := getOpenSSLCertThroughCommand(openSSLCommand, privateDir, domains)

		if err != nil {
			panic("Failed to Generate Certificate/Key Files")
		}

		cert = privateDir + "\\server.cert"
		key = privateDir + "\\server.key"
	}

	return cert, key, certFunc
}

// FetchOpenSSLLib is ...
func FetchOpenSSLLib(privateDir string) (OpenSSLLib, error) {
	// downloadOpenSSLSource(privateDir)
	// return getOpenSSLLib(privateDir)
	return OpenSSLLib{}, errors.New("Function Not Implemented")
}
