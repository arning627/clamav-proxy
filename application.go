package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arning627/clamav-proxy/config"
	"github.com/arning627/clamav-proxy/web"
)

func main() {
	log.Println("Server start listening", config.InitializeConfig.System.Port)
	http.HandleFunc("/ping", web.Ping)
	http.HandleFunc("/scan", web.Scan)
	http.HandleFunc("/", web.Execute)
	err := http.ListenAndServe(fmt.Sprint(":", config.InitializeConfig.System.Port), nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err.Error())
	}
}
