package main

import (
	"log"
	service "github.com/madhu72/lowcode/service"
	example "github.com/madhu72/lowcode/example"
	example2 "github.com/madhu72/lowcode/example2"
)

func ShowInfo(obj service.LowcodeService) {
	obj.Start()
	obj.End()
}
func main() {
	log.Println("Bootstraping LOWCODE Generator")
	var test example.Example
	var test2 example2.Example2
	ShowInfo(test)
	ShowInfo(test2)
}
