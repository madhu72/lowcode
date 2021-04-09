package github.com/madhu72/lowcode/example

import (
	"log"
)
type Example struct {

}

func (e *Example) Start() {
	log.Println("Starting Example Service")
}

func (e *Example) Stop() {
	log.Println("Ending Example Service")
}

func (e *Example) ConnectDb() {
	log.Println("Connecting to Database")
}

func (e *Example) DisconnectDb() {
	log.Println("Disconnecting Database")
}

func (e *Example) ConnectMQ() {
	log.Println("Connecting to MQ")
}

func (e *Example) DisconnectMQ() {
	log.Println("Disconnecting MQ")

}

