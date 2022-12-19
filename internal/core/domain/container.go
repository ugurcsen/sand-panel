package domain

type Container struct {
	ID   string `json:"Id"`
	Name string
}

type ContainerCollection struct {
	Name       string
	Containers []Container
}

type ContainerStats struct {
	CPUUsage    uint64
	MemoryUsage uint64
}
