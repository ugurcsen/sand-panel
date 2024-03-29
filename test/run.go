package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"os"
	"time"
)

func Run() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// List all container-engines
	args := filters.NewArgs()
	args.Add("label", "com.docker.compose.project=ugurcsencom")
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true, Filters: args})
	errorCheck(err)
	var stat types.ContainerStats
	a := types.Stats{}
	for _, container := range containers {
		stat, err = cli.ContainerStats(ctx, container.ID, true)
		errorCheck(err)
		break
	}
	decoder := json.NewDecoder(stat.Body)
	go func() {
		for {
			err := decoder.Decode(&a)
			fmt.Println(a.MemoryStats.MaxUsage)
			if err != nil {
				return
			}
		}
	}()
	fmt.Println("Hi")
	time.Sleep(time.Minute)
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
