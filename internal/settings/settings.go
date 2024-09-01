package settings

import (
	"log"
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"Host"`
	Addr     Addr   `yaml:"Addr"`
	Gotify   Gotify `yaml:"Gotify"`
	StartMsg bool   `yaml:"StartMsg"`
}

type Addr struct {
	Ip   string `yaml:"Ip"`
	Port string `yaml:"Port"`
}

type Gotify struct {
	URL    string `yaml:"URL"`
	ApiKey string `yaml:"ApiKey"`
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
	u, err := url.ParseRequestURI(cfg.Gotify.URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("URL is valid: %s", u)

	return cfg
}

func (c *Config) PrintLogo() {
	log.Printf(`
=====================================================
 ███████  ██████ ██   ██ ██ ██████  ███    ██  █████ 
██       ██      ██   ██ ██ ██   ██ ████   ██ ██   ██
█████    ██      ███████ ██ ██   ██ ██ ██  ██ ███████
██       ██      ██   ██ ██ ██   ██ ██  ██ ██ ██   ██
 ███████  ██████ ██   ██ ██ ██████  ██   ████ ██   ██
=====================================================
Catcher started in: %s:%s
Gotify server: %s
=====================================================
`, c.Addr.Ip, c.Addr.Port, c.Gotify.URL)
}
