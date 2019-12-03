package main

import "log"

var config *Config

func main() {
	start()
}

func start() {
	config = NewConfig("")
	log.Println(startServer())
}
