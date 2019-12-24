package conf

import (
	"log"
	"os"
)

type configuration struct {
	port string
}

var conf = configuration{}

func init() {
	log.Println("Loading configuration")
	conf.port = getValue("PORT", "8080")
	log.Println("Done loading configuration")
}

func GetPort() string {
	return conf.port
}

func getValue(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}
	log.Println("Using value ", value)
	return value
}
