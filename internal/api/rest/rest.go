package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"net"
	"reflect"
)

// rest is the rest api
type rest struct {
	options  *Options
	services *Services
}

// Options is the options for the rest api
type Options struct {
	Listener net.Listener
}

// Services is the services that the rest api will use
type Services struct {
	UserService ports.UserService
}

// RegisterServices registers the services to the rest api
func (r *rest) RegisterServices(services interface{}) error {
	if s, ok := services.(*Services); ok {
		r.services = s
		return nil
	}

	return fmt.Errorf("RegisterServices: only support %v", reflect.TypeOf(&Services{}))
}

// Listen starts the rest api
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

// NewRest creates a new rest api
func NewRest(options *Options) (*rest, error) {
	return &rest{options: options}, nil
}
