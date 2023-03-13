package domain

type Service struct {
	Id           string `json:"Id"`
	Name         string
	CollectionId string
	Collection   *Collection
	Image        string
	Hosts        []string
	Volumes      []Volume
}

type ServiceStats struct {
	CPUUsage    uint64
	MemoryUsage uint64
}

type Volume struct {
	From string
	To   string
}
