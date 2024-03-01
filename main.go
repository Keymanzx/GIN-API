package main

import (
	"flag"
	"fmt"
	"os"

	"api-gin/src/config"
	// "api-gin/src/db"
	"api-gin/src/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	// db.Init()
	server.Init()
}
