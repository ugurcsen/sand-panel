package ports

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
)

type ContainerEngine interface {
	ServiceStats(service *domain.Service) (*domain.ServiceStats, error)
	CreateService(service *domain.Service) error
	DeleteService(service *domain.Service) error
	StartService(service *domain.Service) (*domain.ServiceOperationResponse, error)
	StopService(service *domain.Service) (*domain.ServiceOperationResponse, error)

	//----------------------------------------------------------------------

	CreateCollection(collection *domain.Collection) (*domain.Collection, error)
	StartCollection(collection *domain.Collection) (*domain.ServiceOperationResponse, error)
	StopCollection(collection *domain.Collection) (*domain.ServiceOperationResponse, error)
	DeleteCollection(collection *domain.Collection) error
	UpCollection(collection *domain.Collection) (*domain.ServiceOperationResponse, error)
	DownCollection(collection *domain.Collection) (*domain.ServiceOperationResponse, error)
}
