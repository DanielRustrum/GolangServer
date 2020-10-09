package server

// TODO: Ready Message
// TODO: Certs for Mobile and non-local devices
// TODO: Microsoft Auth
// TODO: Domain Paths
// TODO: Specify Auth Redirect
// TODO: NGINX DNS Server Boot
// TODO: Error Support
// TODO: CORS Support
// TODO: File Handler > Inprogress
// TODO: REST Handler
// TODO: Debug Handler

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/DanielRustrum/Https-Go-Server/package/handler"
)

//* Server Logic
type subdomainHandler map[string]http.Handler
type domainMap map[string]func(http.ResponseWriter, *http.Request)

var configData ConfigData = ConfigData{}
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

func genCert() {
	certCommand := "-Command mkcert -cert-file " +
		configData.PrivateDir +
		"/server.cert -key-file " +
		configData.PrivateDir +
		"/server.key"

	domainCert := ""

	for key := range domains {
		if key == "" {
			domainCert = domainCert + " " + configData.Host
		} else {
			domainCert = domainCert + " " + key + "." + configData.Host
		}
	}

	commandList := strings.Fields(certCommand + domainCert)

	cmd1 := exec.Command("./scripts/create-cert-authority.bat")
	err := cmd1.Run()
	if err != nil {
		log.Fatalf("cert failed to be created... mkcert command failed with %s\n", err)
	}

	cmd := exec.Command("powershell.exe", commandList...)

	err = cmd.Run()
	if err != nil {
		log.Fatalf("cert failed to be created... mkcert command failed with %s\n", err)
	}
}

//* Public

//ConfigData is ...
type ConfigData struct {
	Host       string
	Port       string
	PrivateDir string
	AppendWWW  bool
}

//AddDomain is ...
func AddDomain(key string, handler func(http.ResponseWriter, *http.Request)) {
	domains[key] = handler

	if configData.AppendWWW {
		if key == "" {
			domains["www"] = handler
		} else {
			domains["www."+key] = handler
		}
	}
}

//Setup is ...
func Setup(data ConfigData) {
	if !ranSetup {
		configData = data
		handler.Setup()
		domains = make(domainMap)
	}
	ranSetup = true
}

//Run is ...
func Run() {
	genCert()

	subdomains := make(subdomainHandler)

	for key, value := range domains {
		tempMux := http.NewServeMux()
		tempMux.HandleFunc("/", value)
		subdomains[key] = tempMux
	}

	fmt.Printf("Server Ready\n")
	fmt.Printf("Website available on https://%s:%s\n", configData.Host, configData.Port)

	cfg := &tls.Config{
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

	log.Fatal(srv.ListenAndServeTLS("private/server.cert", "private/server.key"))
}
