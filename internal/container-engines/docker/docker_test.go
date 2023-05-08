package docker

import (
	"github.com/pkg/errors"
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"os"
	"path"
	"testing"
)

var testPath = os.Getenv("HOME") + "/sand-panel-test-folder"

var d = New(testPath)
var testCollection = &domain.Collection{
	Id:     "test",
	Name:   "TestCollection",
	UserId: "TestUser",
	Services: []*domain.Service{
		{
			Id:    "test",
			Name:  "TestService",
			Image: "nginx",
			Hosts: []string{"test1.local", "test2.local"},
			Volumes: []domain.Volume{
				{Name: "html", Path: "/usr/share/nginx/html"},
			},
			Env:  domain.ServiceEnvStaging,
			Port: "80",
		},
	},
}

func TestDocker_GetService(t *testing.T) {

}

func TestDocker_ListServices(t *testing.T) {

}

func TestDocker_ServiceStats(t *testing.T) {

}

func TestDocker_CreateService(t *testing.T) {

}

func TestDocker_AddService(t *testing.T) {

}

func TestDocker_DeleteService(t *testing.T) {

}

func TestDocker_StartService(t *testing.T) {

}

func TestDocker_StopService(t *testing.T) {

}

func TestDocker_GetCollection(t *testing.T) {

}

func TestDocker_ListCollections(t *testing.T) {

}

func TestDocker_CreateCollection(t *testing.T) {
	collection, err := d.CreateCollection(testCollection)

	if err != nil {
		t.Fatal(err)
	}

	if collection.Id != "test" {
		t.Error("collection id not equal")
	}

	if collection.Name != "TestCollection" {
		t.Error("collection name not equal")
	}

	if collection.UserId != "TestUser" {
		t.Error("collection user id not equal")
	}

	if _, err = os.Stat(path.Join(testPath, "/TestUser/TestUser_test")); err != nil {
		t.Error("collection folder not created")
	}

	if _, err = os.Stat(path.Join(testPath, "/TestUser/TestUser_test/docker-compose.yml")); err != nil {
		t.Error("docker-compose.yml not created")
	}

	collection, err = d.CreateCollection(testCollection)

	if err == nil {
		t.Error("collection already exists but no error")
	}

	if err = os.RemoveAll(testPath); err != nil {
		t.Error("test folder not removed")
	}
}

func TestDocker_StartCollection(t *testing.T) {

}

func TestDocker_StopCollection(t *testing.T) {

}

func TestDocker_DeleteCollection(t *testing.T) {

}

func TestDocker_UpCollection(t *testing.T) {
	collection, err := d.CreateCollection(testCollection)
	if err != nil {
		if errors.Cause(err) != ErrorCollectionAlreadyExists {
			t.Fatal(err)
		}
		collection = testCollection
	}
	pipes, err := d.UpCollection(collection)
	if err != nil {
		t.Fatal(err)
	}
	_ = pipes
}

func TestDocker_DownCollection(t *testing.T) {

}

func TestDocker_New(t *testing.T) {

}
