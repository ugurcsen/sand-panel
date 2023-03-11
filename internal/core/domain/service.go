package domain

type Service struct {
	Id           string `json:"Id"`
	Name         string
	CollectionId string
	Collection   *Collection
}

type ServiceStats struct {
	CPUUsage    uint64
	MemoryUsage uint64
}
