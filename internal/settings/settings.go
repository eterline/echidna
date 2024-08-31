package settings

import (
	"log"
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string
	Addr     Addr
	Gotify   Gotify
	StartMsg bool
}

type Addr struct {
	Ip   string
	Port string
}

type Gotify struct {
	URL    string
	ApiKey string
}

func Parse() Config {
	file, err := os.Open("init/settings.yml")
	if err != nil {
		log.Fatal(err.Error())
	}

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	u, err := url.ParseRequestURI("http://google.com/")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("URL is valid: %s", u)

	return cfg
}
