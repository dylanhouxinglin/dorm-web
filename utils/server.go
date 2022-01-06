package utils

type Server interface {
	RegisterUrl()
	Run() error
	EtcdRegister()	error
}