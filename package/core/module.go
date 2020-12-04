package core

import (
	"strings"
	httpcore "github.com/DanielRustrum/Https-Go-Server/package/modules/http"
)

//! Need to update when adding new module

// Config is ...
type Config struct {
	HttpConfig httpcore.Config
}

//! Need to update when adding new module

//Package is ...
type Package struct{
	HttpPackage httpcore.Package
}

//! Need to update when adding new module

//Use is ...
func Use(module string) (err error)
	switch fmtMod := strings.ToLower(module); fmtMod{

	case "http":
		useList[module] = httpcore.Use
		builtpack.HttpPackage = httpcore.GetPackage
	default:
		return errors.new("module not found")

	}

	return nil
}

//GetModules is ...
func GetModules() Package {
	return builtpack
}


//* Internal Operations
var useList map[string]func(Config) = make(map[string]func(Config))
var builtpack Package = Package{}

//! Need to update when adding new module

func runUse(module string, callback func(Config)) {
	switch fmtMod := strings.ToLower(module); fmtMod {

	case "http":
		callback(Config{
			HttpConfig: configData.HttpConfig,
		})

	}
}

//! Need to update when adding new module

func runRun(module string) {
	switch fmtMod := strings.ToLower(module); fmtMod {

	case "http":
		httpcore.Run()

	}
}