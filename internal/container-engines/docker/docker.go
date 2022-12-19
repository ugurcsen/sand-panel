package docker

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
)

type docker struct {
}

func (d docker) Get(id string) (*domain.Container, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) List() (domain.ContainerCollection, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) Stats(id string) (*domain.ContainerStats, error) {
	//TODO implement me
	panic("implement me")
}

func New() ports.ContainerEngine {
	return &docker{}
}
