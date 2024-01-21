package main

import (
	"context"
	"fmt"
	"tgrzimiar/go-scylla/config"
	"tgrzimiar/go-scylla/pkg/database/log"
	"tgrzimiar/go-scylla/pkg/database/scylla"
	"tgrzimiar/go-scylla/server"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func main() {

	logger := log.CreateLogger("info")
	ctx := context.Background()

	// cfg := config.LoadConfig(func() string {
	// 	if len(os.Args) < 2 {
	// 		logger.Fatal("Error: .env path is invalid")
	// 	}
	// 	return os.Args[1]
	// }())

	cluster := scylla.CreateCluster(gocql.Quorum, "users", "scylla-node1", "scylla-node2")
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		logger.Fatal("unable to connect to scylla", zap.Error(err))
	}
	defer session.Close()
	fmt.Println("start")
	server.Start(ctx, &config.Config{}, session, logger)

}
