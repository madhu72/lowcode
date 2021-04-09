package example2

import (
	"log"
)
type Example2 struct {

}

func (e *Example2) Start() {
	log.Println("Starting Example2 Service")
}

func (e *Example2) Stop() {
	log.Println("Ending Example2 Service")
}

func (e *Example2) ConnectDb() {
	log.Println("Connecting to Database")
}

func (e *Example2) DisconnectDb() {
	log.Println("Disconnecting Database")
}

func (e *Example2) ConnectMQ() {
	log.Println("Connecting to MQ")
}

func (e *Example2) DisconnectMQ() {
	log.Println("Disconnecting MQ")

}

