package nginx

// TODO: NGINX DNS Server Boot
// TODO: DNS Support
import (
	"github.com/DanielRustrum/Https-Go-Server/package/core"
)

//Config is ...
type Config struct{}

//Package is ...
type Package struct{}

//Use is ...
func Use(data core.Config) {}

//GetPackage is ...
func GetPackage() Package {
	return Package{}
}

//Run is ...
func Run() {}
