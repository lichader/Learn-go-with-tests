package main_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/quii/go-specs-greet/adapters"
	httpserver "github.com/quii/go-specs-greet/adapters/httpserver"
	"github.com/quii/go-specs-greet/specifications"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../../.",
			Dockerfile:    "./cmd/httpserver/Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	port := "8080"
	dockerFilePath := "./cmd/httpserver/Dockerfile"
	baseURL := fmt.Sprintf("http://localhost:%s", port)
	driver := httpserver.Driver{BaseURL: baseURL}

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
