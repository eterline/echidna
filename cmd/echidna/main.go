package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eterline/echidna/internal/gotify"
	"github.com/eterline/echidna/internal/server"
	"github.com/eterline/echidna/internal/settings"
)

func main() {
	var srv server.Server
	cfg := settings.Parse()
	srv.New(cfg)
	gotify.StartMessage(cfg)
	addr := fmt.Sprintf("%s:%s", cfg.Addr.Ip, cfg.Addr.Port)
	fmt.Println(cfg.Addr)
	err := http.ListenAndServe(addr, srv.Router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
