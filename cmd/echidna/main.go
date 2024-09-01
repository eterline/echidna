package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/eterline/echidna/internal/gotify"
	"github.com/eterline/echidna/internal/server"
	"github.com/eterline/echidna/internal/settings"
)

func main() {
	logPath := fmt.Sprintf(
		"logs/echidna_%v_%v_%v-%v_%v_%v.log",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second(),
	)
	file, _ := os.Create(logPath)
	log.SetOutput(file)
	defer file.Close()

	cfg := settings.Parse()
	gotify.StartMessage(cfg)
	srv := server.New(cfg)
	addr := fmt.Sprintf("%s:%s", cfg.Addr.Ip, cfg.Addr.Port)
	cfg.PrintLogo()
	err := http.ListenAndServe(addr, srv.Router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
