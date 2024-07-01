package main_test

import (
	"fmt"
	"testing"

	"github.com/quii/go-specs-greet/adapters"
	httpserver "github.com/quii/go-specs-greet/adapters/httpserver"
	"github.com/quii/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	port := "8080"
	baseURL := fmt.Sprintf("http://localhost:%s", port)
	driver := httpserver.Driver{BaseURL: baseURL}

	adapters.StartDockerServer(t, port, "httpserver")
	specifications.GreetSpecification(t, driver)
}
