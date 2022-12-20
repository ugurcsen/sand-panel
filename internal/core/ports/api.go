package ports

type API interface {
	Listen() error
	RegisterServices(collection interface{}) error
}
