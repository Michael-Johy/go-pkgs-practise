package main

import (
	"github.com/Michael-Johy/go-pkgs-practise/pprof/data"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("Hello, World"))
		}
	}()
	_ = http.ListenAndServe(":6060", nil)
	//http://localhost:6060/debug/pprof/
	//go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60
}
