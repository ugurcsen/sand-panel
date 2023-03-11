package docker

import (
	"github.com/ugurcsen/sand-panel/internal/core/domain"
	"os"
	"testing"
)

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
	collection, err := New(os.Getenv("HOME") + "/sand-panel-test-folder").CreateCollection(&domain.Collection{
		Id:       "test",
		Name:     "TestCollection",
		UserId:   "TestUser",
		User:     nil,
		Services: nil,
	})

	if err != nil {
		t.Error(err)
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

	if _, err = os.Stat(os.Getenv("HOME") + "/sand-panel-test-folder/TestUser/TestUser_test"); err != nil {
		t.Error("collection folder not created")
	}

	if _, err = os.Stat(os.Getenv("HOME") + "/sand-panel-test-folder/TestUser/TestUser_test/docker-compose.yml"); err != nil {
		t.Error("docker-compose.yml not created")
	}
}

func TestDocker_StartCollection(t *testing.T) {

}

func TestDocker_StopCollection(t *testing.T) {

}

func TestDocker_DeleteCollection(t *testing.T) {

}

func TestDocker_UpCollection(t *testing.T) {

}

func TestDocker_DownCollection(t *testing.T) {

}

func TestDocker_New(t *testing.T) {

}
