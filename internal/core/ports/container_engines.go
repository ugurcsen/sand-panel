package ports

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
)

type ContainerEngine interface {
	GetService(serviceId string) (*domain.Service, error)
	ListServices(collectionId string) ([]*domain.Service, error)
	ServiceStats(serviceId string) (*domain.ServiceStats, error)
	CreateService(service *domain.Service) (*domain.Service, error)
	AddService(collectionId string, service domain.Service) error
	DeleteService(serviceId string) error
	StartService(serviceId string) (domain.Pipes, error)
	StopService(serviceId string) (domain.Pipes, error)

	//----------------------------------------------------------------------

	GetCollection(collectionId string) (*domain.Collection, error)
	ListCollections(userId string) ([]*domain.Collection, error)
	CreateCollection(collection *domain.Collection) (*domain.Collection, error)
	StartCollection(collectionId string) (domain.Pipes, error)
	StopCollection(collectionId string) (domain.Pipes, error)
	DeleteCollection(collectionId string) error
	UpCollection(collectionId string) (domain.Pipes, error)
	DownCollection(collectionId string) (domain.Pipes, error)
}
