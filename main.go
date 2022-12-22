package main

import (
	"csv_reader/db"
	"csv_reader/server"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	f, _ := os.OpenFile("Logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	db.Init()
	server.Init()
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	log.SetOutput(f)
}
