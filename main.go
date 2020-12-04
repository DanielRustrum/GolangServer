package main

import (
	server "github.com/DanielRustrum/Https-Go-Server/package/core"
)

func main() {

	//? Old Way
	// server.Setup(server.HTTPConfigData{
	// 	Host:       "localhost",
	// 	Port:       "9010",
	// 	PrivateDir: "private",
	// 	AppendWWW:  true,

	// 	OpenSSLCommand:    "C:\\Users\\danie\\Documents\\OpenSSL-Win64\\bin\\openssl.exe",
	// 	City:              "Flagstaff",
	// 	CountryCode:       "US",
	// 	StateOrProvidence: "Arizona",
	// })

	// routeMap := make(map[string]string)
	// routeMap["/"] = "./index.html"
	// routeMap["404"] = "./404.html"

	// fileServer := handler.FileServer("./public", routeMap)

	// server.AddDomain("", fileServer.Handler)
	// server.Run()

	_ := server.Use("http")

	_ := server.Setup(server.Configs{})
	modules := server.GetModules()

	_ := server.Run()
}
