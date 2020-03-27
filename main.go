package main

import (
	"./protoserver"
	"./ticker"
)

func main() {
	protoserver.Serve()
	ticker.Ticker()
}
