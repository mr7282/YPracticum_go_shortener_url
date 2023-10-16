package config

import (
	"flag"
	"os"
)

var ServAddr *string
var BaseAddr *string


func ParseFlag() {
	ServAddr = flag.String("a", "localhost:8080", "server start address")
	BaseAddr = flag.String("b", "http://" + *ServAddr, "base response address")
	flag.Parse()

	if envServAddr := os.Getenv("SERVER_ADDRESS"); envServAddr != "" {
        ServAddr = &envServAddr
    }

	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
        BaseAddr = &envBaseAddr
    }
}