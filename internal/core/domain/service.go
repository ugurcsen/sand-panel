package domain

type Service struct {
	Id           string `json:"Id"`
	Name         string
	CollectionId string
	Collection   *Collection
	Image        string
	Hosts        []string
	Volumes      []Volume
	Port         string
	Env          ServiceEnv
}

type ServiceStats struct {
	CPUUsage    uint64
	MemoryUsage uint64
}

type Volume struct {
	Name string
	Path string
}
