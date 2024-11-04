package main

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func main() {
	ref := "alpine"

	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)

	ctx := context.Background()

	rc, err := cli.ImagePull(ctx, ref, image.PullOptions{})
	defer func(rc io.ReadCloser) {
		if rc != nil {
			err := rc.Close()
			if err != nil {
				panic(err)
			}
		}
	}(rc)

	if err != nil {
		// if the error containers "docker login", print that out and end the program.
		if strings.Contains(err.Error(), "docker login") {
			fmt.Println(err.Error())
			return
		}
		panic(err)
	}
	if _, err := io.Copy(io.Discard, rc); err != nil {
		panic(err)
	}

	fmt.Println("Success!")
}
