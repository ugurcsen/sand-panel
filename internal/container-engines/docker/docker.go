package docker

import (
	"github.com/pkg/errors"
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"os"
)

var _ ports.ContainerEngine = &docker{}

type docker struct {
	BaseDir string
}

func (d docker) GetService(serviceId string) (*domain.Service, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) ListServices(collectionId string) ([]*domain.Service, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) ServiceStats(serviceId string) (*domain.ServiceStats, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) CreateService(service *domain.Service) (*domain.Service, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) AddService(collectionId string, service domain.Service) error {
	//TODO implement me
	panic("implement me")
}

func (d docker) DeleteService(serviceId string) error {
	//TODO implement me
	panic("implement me")
}

func (d docker) StartService(serviceId string) (domain.Pipes, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) StopService(serviceId string) (domain.Pipes, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) GetCollection(collectionId string) (*domain.Collection, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) ListCollections(userId string) ([]*domain.Collection, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) CreateCollection(c *domain.Collection) (*domain.Collection, error) {
	var collectionPath = d.getCollectionPath(c)

	if _, err := os.Stat(collectionPath); err == nil {
		return nil, errors.Wrap(ErrorCollectionAlreadyExists, collectionPath)
	}

	if err := os.MkdirAll(collectionPath, 0755); err != nil {
		return nil, errors.Wrap(ErrorCollectionNotCreated, collectionPath)
	}

	composePath := d.getComposePath(c)
	f, err := os.Create(composePath)
	if err != nil {
		return nil, errors.Wrap(ErrorCollectionFileNotCreated, composePath)
	}

	err = composeYamlBuilder(c, f)
	if err != nil {
		return nil, errors.Wrap(ErrorCollectionFileNotCreated, err.Error())
	}
	return c, nil
}

func (d docker) StartCollection(collectionId string) (domain.Pipes, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) StopCollection(collectionId string) (domain.Pipes, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) DeleteCollection(collectionId string) error {
	//TODO implement me
	panic("implement me")
}

func (d docker) UpCollection(collectionId string) (domain.Pipes, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) DownCollection(collectionId string) (domain.Pipes, error) {
	//TODO implement me
	panic("implement me")
}

func New(baseDir string) ports.ContainerEngine {
	return &docker{BaseDir: baseDir}
}
