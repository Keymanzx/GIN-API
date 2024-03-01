package main

import (
	"flag"
	"fmt"
	"os"

	// "api-gin/src/db"
	"api-gin/src/server"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	// db.Init()
	server.Init()
}
