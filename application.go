package main

import (
	"log"
	"net/http"

	server "github.com/Arning627/clamav-proxy/web"
	"github.com/go-ini/ini"
)

func main() {
	log.Println("server start listening 8088")
	http.HandleFunc("/ping", server.Ping)
	err := http.ListenAndServe(":22", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err.Error())
	}
}

type Config struct {
	System System `ini:"system"`
	Clamav Clamav `ini:"clamav"`
}
type System struct {
	Port string `ini:"port"`
}
type Clamav struct {
	Port string `ini:"port"`
	Host string `ini:"host"`
}

var (
	CONFIG = new(Config)
)

func init() {
	log.Println("Application initializing...")
	err := ini.MapTo(CONFIG, "config.ini")
	if err != nil {
		log.Fatalf("Failed to read configuration file")
	}
}
