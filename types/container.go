package types

import "github.com/docker/docker/api/types"

type Container types.Container

type Collection struct {
	Name       string
	Containers []types.Container
}

type Namespace struct {
	Name        string
	Collections []Collection
}
