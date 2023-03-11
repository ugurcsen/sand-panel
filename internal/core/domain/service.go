package domain

type Service struct {
	Id           string `json:"Id"`
	Name         string
	CollectionId string
	Collection   *Collection
	Image        string
	Host         string
}

type ServiceStats struct {
	CPUUsage    uint64
	MemoryUsage uint64
}
