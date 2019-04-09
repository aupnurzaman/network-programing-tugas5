package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"log"

	"github.com/spf13/viper"
)

const (
	HTTP_PORT           string = "http.port"
	HTTP_DIR            string = "http.dir"
)

type HttpConfig struct {
	Port       int
	Dir        string
}

type Config struct {
	*viper.Viper
}

func setDefaults(v *viper.Viper) {
	v.SetDefault(HTTP_PORT, 8000)
	v.SetDefault(HTTP_DIR, "static")
}

func loadConfPath(v *viper.Viper, path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return v.ReadConfig(bytes.NewBuffer(f))
}

func (c *Config) HttpCfg() *HttpConfig {
	return &HttpConfig{
		Port:      c.GetInt(HTTP_PORT),
		Dir:       c.GetString(HTTP_DIR),
	}
}

func NewCfg(path ...string) (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	setDefaults(v)

	if len(path) > 0 {
		if err := loadConfPath(v, path[0]); err != nil {
			log.Println("Failed load from file. Using default configuration")
		}
	}

	return &Config{v}, nil
}
