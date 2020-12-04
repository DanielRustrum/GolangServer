package http

// TODO: Ready Message
// TODO: Non-Localhost server boot
// TODO: Http to Https redirect

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/DanielRustrum/Https-Go-Server/package/core"
)

//* Server Logic
type subdomainHandler map[string]http.Handler
type domainMap map[string]func(http.ResponseWriter, *http.Request)

var configData HTTPConfigData = HTTPConfigData{}
var domains map[string]func(http.ResponseWriter, *http.Request)
var ranSetup bool = false

func (subdomains subdomainHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	//* Static Headers
	response.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	//* Subdomain Management
	var mux http.Handler

	fullHost := configData.Host + ":" + configData.Port
	domainParts := strings.Split(request.Host, ".")[0]

	if domainParts == fullHost {
		mux = subdomains[""]
	} else {
		mux = subdomains[domainParts]
	}

	//* Serve Management
	if mux != nil {
		mux.ServeHTTP(response, request)
	} else {
		http.Error(response, "Not found", 404)
	}
}

func genDomainString() string {
	domainString := ""

	for key := range domains {
		if key == "" {
			domainString = domainString + " " + configData.Host
		} else {
			domainString = domainString + " " + key + "." + configData.Host
		}
	}

	return domainString
}

func setup(data HTTPConfigData) {
	if !ranSetup {
		configData = data
		domains = make(domainMap)
	}
	ranSetup = true
}

func addDomain(key string, handler func(http.ResponseWriter, *http.Request)) {
	domains[key] = handler

	if configData.AppendWWW {
		if key == "" {
			domains["www"] = handler
		} else {
			domains["www."+key] = handler
		}
	}
}

//* Public

//Config is ...
type Config struct {
	Host       string `default:"localhost"`
	Port       string `default:"8000"`
	PrivateDir string `default:".private"`
	AppendWWW  bool   `default:"false"`

	//* OpenSSL Information
	OpenSSLCommand    string `default:""`
	CountryCode       string `default:"."`
	City              string `default:"."`
	StateOrProvidence string `default:"."`
	Organization      string `default:"."`
	OrganizationUnit  string `default:"."`
}

//Package is ...
type Package struct {
	AddDomain func(key string, handler func(http.ResponseWriter, *http.Request))
}

//Use is ...
func Use(data core.Config) {

}

//GetPackage is ...
func GetPackage() Package {
	return Package{
		AddDomain: addDomain,
	}
}

//Run is ...
func Run() {

	subdomains := make(subdomainHandler)

	for key, value := range domains {
		tempMux := http.NewServeMux()
		tempMux.HandleFunc("/", value)
		subdomains[key] = tempMux
	}

	fmt.Printf("Server Ready\n")
	fmt.Printf("Website available on https://%s:%s\n", configData.Host, configData.Port)

	cert := GetCert(
		configData.Host,
		genDomainString(),
	)

	if cert == nil {
		//TODO: Default to http if Cert is not created
	} else {

		cfg := &tls.Config{
			GetCertificate:           cert,
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}

		srv := &http.Server{
			Addr:         ":" + configData.Port,
			Handler:      subdomains,
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}
}
