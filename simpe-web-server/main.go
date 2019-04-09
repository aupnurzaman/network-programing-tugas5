package main

import (
	"net/http"
	"log"
	"flag"
	"fmt"
)

var (
	cnf        *Config
	configPath string
)

func initConfig() error {
	flag.StringVar(&configPath, "c", "config.yaml", "Configuration File")
	flag.Parse()

	c, err := NewCfg(configPath)
	if err != nil {
		return err
	}
	cnf = c

	return err
}

func main() {
  err := initConfig()
  if err != nil {
	log.Fatal(err)
  }
	
  fs := http.FileServer(http.Dir(cnf.HttpCfg().Dir))
  http.Handle("/", fs)

  log.Println("Listening...")
  port := fmt.Sprintf(":%d", cnf.HttpCfg().Port)
  http.ListenAndServe(port, nil)
}