package main_test

import (
	"fmt"
	"testing"

	"github.com/SamMotta/go-specs-greet/adapters"
	"github.com/SamMotta/go-specs-greet/adapters/grpcserver"
	"github.com/SamMotta/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port       = "50051"
		binToBuild = "grpcserver"
		driver     = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	t.Cleanup(driver.Close)
	adapters.StartDockerServer(t, port, binToBuild)
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
