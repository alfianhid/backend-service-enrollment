// Package main is the entry point to start the cerebrum server
package main

import (
	"backend-service/pkg/api"
	"backend-service/pkg/utl/config"
	"backend-service/pkg/utl/support"
)

// main server
func main() {
	cfgPath, err := support.ExtractPathFromFlags()
	if err != nil {
		panic(err.Error())
	}

	cfg, err := config.LoadConfigFrom(cfgPath)
	if err != nil {
		panic(err.Error())
	}

	if cfg == nil {
		panic("unknown error loading yaml file")
	}

	if err = api.Start(cfg); err != nil {
		panic(err.Error())
	}
}
