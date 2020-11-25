package main

import (
	"github.com/DanielRustrum/Https-Go-Server/package/handler"
	"github.com/DanielRustrum/Https-Go-Server/package/server"
)

func main() {
	server.Setup(server.ConfigData{
		Host:       "localhost",
		Port:       "9010",
		PrivateDir: "private",
		AppendWWW:  true,

		OpenSSLCommand:    "C:\\Users\\danie\\Documents\\OpenSSL-Win64\\bin\\openssl.exe",
		City:              "Flagstaff",
		CountryCode:       "US",
		StateOrProvidence: "Arizona",
	})

	routeMap := make(map[string]string)
	routeMap["/"] = "./index.html"
	routeMap["404"] = "./404.html"

	fileServer := handler.FileServer("./public", routeMap)

	server.AddDomain("", fileServer.Handler)
	server.Run()
}
