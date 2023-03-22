package config

import (
	"log"

	"github.com/go-ini/ini"
)

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
	InitializeConfig = new(Config)
)

func init() {
	log.Println("Application initializing...")
	err := ini.MapTo(InitializeConfig, "config.ini")
	if err != nil {
		log.Fatalf("Failed to read configuration file")
	}
}
