package ports

import "github.com/ugurcsen/sand-panel/internal/core/domain"

type ContainerEngine interface {
	Get(id string) (*domain.Container, error)
	List() (domain.ContainerCollection, error)
	Stats(id string) (*domain.ContainerStats, error)
}
