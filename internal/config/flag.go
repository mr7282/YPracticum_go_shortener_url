package config

import "flag"

var ServAddr *string
var BaseAddr *string


func ParseFlag() {
	ServAddr = flag.String("a", "localhost:8080", "server start address")
	BaseAddr = flag.String("b", "http://" + *ServAddr, "base response address")
	flag.Parse()
}