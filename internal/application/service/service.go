package application

import (
	ip_ebus "github.com/octoposprime/op-be-graphql/internal/application/infrastructure/port/ebus"
	ip_repo "github.com/octoposprime/op-be-graphql/internal/application/infrastructure/port/repository"
	ip_service "github.com/octoposprime/op-be-graphql/internal/application/infrastructure/port/service"
	ds "github.com/octoposprime/op-be-graphql/internal/domain/service"
)

// Service is an application service.
// It manages the business logic of the application.
type Service struct {
	*ds.Service
	ip_repo.RedisPort
	ip_service.ServicePort
	ip_ebus.EBusPort
}

// NewService creates a new *Service.
func NewService(domainService *ds.Service,
	redisRepository ip_repo.RedisPort,
	microService ip_service.ServicePort,
	eBus ip_ebus.EBusPort) *Service {
	service := &Service{
		domainService,
		redisRepository,
		microService,
		eBus,
	}
	service.EBusPort.SetLogger(service.Log)
	//service.EventListen()
	return service
}

//// This is the event listener handler of the application layer.
//func (a *Service) EventListen() *Service {
//	go a.Listen(context.Background(), smodel.ChannelGraphqlNotification, a.EventListenerCallBack)
//	return a
//}
//
//// This is a call-back function of the event listener handler of the application layer.
//func (a *Service) EventListenerCallBack(channelName string, data any) {
//	//TODO: #19 Implement Notification Event
//}
