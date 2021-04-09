package service

type LowcodeService interface {
	Start()
	Stop()
	ConnectDb()
	DisconnectDb()
	ConnectMQ()
	DisconnectMQ()
}
