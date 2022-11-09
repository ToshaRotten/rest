package main

import (
	"flag"
	"log"
	"rest/server/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config file")
}

func main() {
	flag.Parse()
	config, err := apiserver.TakeConfigFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
