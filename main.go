package main

import (
	"log"
	service "github.com/madhu72/lowcode/service"
	example "github.com/madhu72/lowcode/example"
)

func main() {
	log.Println("Bootstraping LOWCODE Generator")
	var test service.Service

	test = example.Example{}

	test.Start()
	test.Stop()
}
