package docker

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ugurcsen/go-docker-compose-client/client"
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"github.com/ugurcsen/sand-panel/internal/core/ports"
	"log"
	"os"
)

var _ ports.ContainerEngine = &docker{}

type docker struct {
	BaseDir string
	log     *log.Logger
	ctx     context.Context
}

func (d docker) StartCollection(collection *domain.Collection) (*domain.ServiceOperationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) StopCollection(collection *domain.Collection) (*domain.ServiceOperationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) DeleteCollection(collection *domain.Collection) error {
	//TODO implement me
	panic("implement me")
}

func (d docker) UpCollection(collection *domain.Collection) (*domain.ServiceOperationResponse, error) {
	return d.UpCollectionWithContext(d.ctx, collection)
}

func (d docker) UpCollectionWithContext(ctx context.Context, collection *domain.Collection) (*domain.ServiceOperationResponse, error) {
	cli, _ := client.NewClientWithContext(ctx, d.getCollectionPath(collection))
	pipes, err := cli.Up()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to start collection %s", collection.Name))
	}

	return &domain.ServiceOperationResponse{
		Stdout: pipes.Stdout,
		Stderr: pipes.Stderr,
		Stdin:  pipes.Stdin,
		Wait:   cli.Wait,
	}, nil
}

func (d docker) DownCollection(collection *domain.Collection) (*domain.ServiceOperationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) CreateService(service *domain.Service) error {
	//TODO implement me
	panic("implement me")
}

func (d docker) ServiceStats(service *domain.Service) (*domain.ServiceStats, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) DeleteService(service *domain.Service) error {
	//TODO implement me
	panic("implement me")
}

func (d docker) StartService(service *domain.Service) (*domain.ServiceOperationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) StopService(service *domain.Service) (*domain.ServiceOperationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d docker) CreateCollection(c *domain.Collection) (*domain.Collection, error) {
	var collectionPath = d.getCollectionPath(c)

	if _, err := os.Stat(collectionPath); err == nil {
		return nil, errors.Wrap(ErrorCollectionAlreadyExists, collectionPath)
	}

	if err := os.MkdirAll(collectionPath, 0755); err != nil {
		return nil, errors.Wrap(ErrorCollectionNotCreated, collectionPath)
	}

	composePath := d.getComposePath(c)
	f, err := os.Create(composePath)
	if err != nil {
		return nil, errors.Wrap(ErrorCollectionFileNotCreated, composePath)
	}
	defer d.handleDefer(f.Close)

	err = composeYamlBuilder(c, f)
	if err != nil {
		return nil, errors.Wrap(ErrorCollectionFileNotCreated, err.Error())
	}
	return c, nil
}

func New(baseDir string) ports.ContainerEngine {
	return NewWithLogger(baseDir, log.Default())
}

func NewWithLogger(baseDir string, logger *log.Logger) ports.ContainerEngine {
	return &docker{BaseDir: baseDir, log: logger, ctx: context.Background()}
}

func (d docker) handleDefer(f func() error) {
	if err := f(); err != nil {
		d.log.Println(err.Error())
	}
}
