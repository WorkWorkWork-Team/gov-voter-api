package handler_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

type ContainerAddress struct {
	Host      string
	Port      string
	Terminate func()
}

var (
	MySQLContainer ContainerAddress
	Engine         *gin.Engine
)

var _ = BeforeSuite(func() {
	fmt.Println("🟢 BeforeSuite Integration test")
	MySQLContainer = setupMySQL()
})

var _ = AfterSuite(func() {
	fmt.Println("⛔️ AfterSuite Integration test")
	MySQLContainer.Terminate()
})

func setupMySQL() ContainerAddress {
	name := "mongo"
	ctx := context.Background()
	containerReq := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForLog("Database files initialized").WithStartupTimeout(time.Second * 10),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerReq,
		Started:          true,
	})

	if err != nil {
		logrus.Fatalf("error starting %s container: %s", name, err)
	}

	if err != nil {
		logrus.Fatalf("%s.Exec: %s", name, err)
	}

	containerHost, _ := container.Host(ctx)

	containerPort, err := container.MappedPort(ctx, "3306")
	if err != nil {
		logrus.Fatalf("pubsubContainer.MappedPort: %s", err)
	}

	terminateContainer := func() {
		logrus.Info("terminating container...")
		if err := container.Terminate(ctx); err != nil {
			logrus.Fatalf("error terminating pubsubContainer container: %v\n", err)
		}
	}
	return ContainerAddress{containerHost, containerPort.Port(), terminateContainer}
}
