package network_service

type NetworkService interface {
	Run()
	Close()
}
