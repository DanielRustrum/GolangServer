package core

import (
	"errors"
)

//Setup is ...
func Setup(config Config) (err error) {
	if !configSet {
		configData = config

		for key, value := range useList {
			runUse(key, value)
		}
	}
	return errors.New("setup already ran")
}

//Run is ...
func Run() (err error) {
	for key, _ := range useList {
		runRun(key)
	}

	return nil
}

var configData Config = Config{}
var configSet bool = false
