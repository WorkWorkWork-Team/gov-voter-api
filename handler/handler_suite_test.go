package handler_test

import (
	"context"
	"fmt"
	"os"
	"path"
	"regexp"
	"testing"
	"time"

	"github.com/WorkWorkWork-Team/common-go/databasemysql"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

var (
	MySQLConnection *sqlx.DB
)

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
	mysql, err := databasemysql.NewDbConnection(databasemysql.Config{
		Hostname:     fmt.Sprint(MySQLContainer.Host, ":", MySQLContainer.Port),
		Username:     "root",
		Password:     "my-secret-pw",
		DatabaseName: "devDB",
	})
	Expect(err).To(BeNil())
	MySQLConnection = mysql
})

var _ = AfterSuite(func() {
	fmt.Println("⛔️ AfterSuite Integration test")
	MySQLContainer.Terminate()
})

func setupMySQL() ContainerAddress {
	name := "mongo"
	ctx := context.Background()
	currentDirectory, err := os.Getwd()
	Expect(err).To(BeNil())
	exp, err := regexp.Compile(`(.*)\/handler`)
	Expect(err).To(BeNil())
	result := exp.FindAllStringSubmatch(currentDirectory, -1)

	containerReq := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "my-secret-pw",
			"MYSQL_ROOT_HOST":     "%",
			"MYSQL_DATABASE":      "devDB",
		},
		Mounts: testcontainers.Mounts(testcontainers.ContainerMount{
			Source: testcontainers.GenericBindMountSource{
				HostPath: path.Join(result[0][1], "database"),
			},
			Target: testcontainers.ContainerMountTarget("/docker-entrypoint-initdb.d"),
		}),
		WaitingFor: wait.ForLog("3306").WithStartupTimeout(time.Second * 20),
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
