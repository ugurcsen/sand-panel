package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"github.com/ugurcsen/sand-panel/internal/core/services"
	"net"
)

type rest struct {
	Options
}

type Options struct {
	Listener    net.Listener
	UserService ports.UserService
}

func (r rest) AddService(service services.Service) error {
	//TODO implement me
	panic("implement me")
}

func (r rest) Listen() error {
	app := iris.New()

	booksAPI := app.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)
	}

	err := app.Run(iris.Listener(r.Listener))

	return err
}

func list(ctx iris.Context) {
	ctx.Writef("Hello from list")
}

func create(ctx iris.Context) {
	ctx.Writef("Hello from list")
}

func NewRest(options Options) (ports.API, error) {
	if options.UserService == nil {
		return nil, fmt.Errorf("UserService is required")
	}
	return &rest{options}, nil
}
