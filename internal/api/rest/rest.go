package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"net"
	"reflect"
)

type rest struct {
	options  *Options
	services *Services
}

type Options struct {
	Listener net.Listener
}

type Services struct {
	UserService ports.UserService
}

func (r *rest) RegisterServices(services interface{}) error {
	if s, ok := services.(*Services); ok {
		r.services = s
		return nil
	}

	return fmt.Errorf("RegisterServices: only support %v", reflect.TypeOf(&Services{}))
}

func (r *rest) Listen() error {
	app := iris.New()

	booksAPI := app.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", func(ctx iris.Context) {
			users, err := r.services.UserService.List()
			if err != nil {
				panic(err)
			}
			ctx.JSON(users)
		})
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)
	}

	err := app.Run(iris.Listener(r.options.Listener))

	return err
}

func create(ctx iris.Context) {
	ctx.Writef("Hello from list")
}

func NewRest(options *Options) (ports.API, error) {
	return &rest{options: options}, nil
}
