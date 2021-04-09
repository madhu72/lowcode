package github.com/madhu72/lowcode/service

type LowcodeService interface {
	Start()
	Stop()
	ConnectDb()
	DisconnectDb()
	ConnectMQ()
	DisconnectMQ()
}
