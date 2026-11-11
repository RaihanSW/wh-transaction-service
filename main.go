package main

import (
	"context"
	"database/sql"
	"dc-transaction-service/app/controller"
	apihandler "dc-transaction-service/app/controller/handler"
	"dc-transaction-service/app/repository"
	"dc-transaction-service/app/service"
	"dc-transaction-service/config"
	"dc-transaction-service/log"
	"dc-transaction-service/migration"
	"dc-transaction-service/seed"
	"flag"
	"fmt"
	syslog "log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	var env string
	flag.StringVar(&env, "env", env, "environment level")
	flag.Parse()

	cfg, err := config.ReadConfig(env)
	if err != nil {
		syslog.Fatalf("Error loading config: %v\n", err)
	}

	// Init log & load config
	log.InitLog()

	// Init monitoring
	config.InitializeProfiler(cfg)

	// Init external connections
	database, err := config.InitDBConnection(cfg)
	if err != nil {
		log.Logging.Fatalf("Error connecting to DB: %v\n", err)
	}

	// DB maintenance subcommands: migrate / seed
	if runDBCommand(database) {
		return
	}

	// Init repository database (single DB: warehouse_db for both write & read)
	generalRepository := repository.NewGeneralRepositoryImpl(cfg.Database["warehouse_db"], database["warehouse_db"])
	transactionRepository := repository.NewTransactionRepositoryImpl(cfg.ServiceName, cfg.Database["warehouse_db"], database["warehouse_db"], cfg.Database["warehouse_db"], database["warehouse_db"])
	transactionStatusRepository := repository.NewTransactionStatusRepositoryImpl(cfg.ServiceName, cfg.Database["warehouse_db"], database["warehouse_db"], cfg.Database["warehouse_db"], database["warehouse_db"])
	goodsRepository := repository.NewGoodsRepositoryImpl(cfg.ServiceName, cfg.Database["warehouse_db"], database["warehouse_db"], cfg.Database["warehouse_db"], database["warehouse_db"])
	goodsStatusRepository := repository.NewGoodsStatusRepositoryImpl(cfg.ServiceName, cfg.Database["warehouse_db"], database["warehouse_db"], cfg.Database["warehouse_db"], database["warehouse_db"])
	goodsCategoryRepository := repository.NewGoodsCategoryRepositoryImpl(cfg.ServiceName, cfg.Database["warehouse_db"], database["warehouse_db"], cfg.Database["warehouse_db"], database["warehouse_db"])

	// Init use case
	generalService := service.NewGeneralServiceImpl(cfg, generalRepository)
	transactionService := service.NewTransactionServiceImpl(cfg, transactionRepository)
	transactionStatusService := service.NewTransactionStatusServiceImpl(cfg, transactionStatusRepository)
	goodsService := service.NewGoodsServiceImpl(cfg, goodsRepository)
	goodsStatusService := service.NewGoodsStatusServiceImpl(cfg, goodsStatusRepository)
	goodsCategoryService := service.NewGoodsCategoryServiceImpl(cfg, goodsCategoryRepository)

	// Init handler
	generalHandler := apihandler.NewGeneralHandler(generalService)
	transactionHandler := apihandler.NewTransactionHandler(transactionService)
	transactionStatusHandler := apihandler.NewTransactionStatusHandler(transactionStatusService)
	goodsHandler := apihandler.NewGoodsHandler(goodsService)
	goodsStatusHandler := apihandler.NewGoodsStatusHandler(goodsStatusService)
	goodsCategoryHandler := apihandler.NewGoodsCategoryHandler(goodsCategoryService)

	// Run, service, run!!!!
	log.Logging.Infoln("Application is running, Port : " + fmt.Sprintf("%v", cfg.Port.Service))
	rest := controller.NewController(transactionHandler, transactionStatusHandler, goodsHandler, goodsStatusHandler, goodsCategoryHandler, generalHandler, cfg, env).StartRoute()

	// Wait for the interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Logging.Infoln("Shutdown Application ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// close everything in here
	// like a connection, channel, or anything
	defer func() {
		cancel()
		for _, db := range database {
			db.Close()
		}
	}()

	if err := rest.Shutdown(ctx); err != nil {
		log.Logging.Fatalf("Server Shutdown: %v", err)
	}
	log.Logging.Infoln("Application Stopped")
}

// runDBCommand handles `migrate` / `seed` subcommands passed after flags,
// e.g. `go run main.go -env local migrate up`. Returns true when a
// subcommand ran and the process should exit instead of serving.
func runDBCommand(database map[string]*sql.DB) bool {
	ctx := context.Background()

	switch flag.Arg(0) {
	case "migrate":
		var err error
		switch flag.Arg(1) {
		case "up":
			err = migration.RunUp(ctx, database)
		case "down":
			steps := 1
			if n, convErr := strconv.Atoi(flag.Arg(2)); convErr == nil {
				steps = n
			}
			err = migration.RunDown(ctx, database, steps)
		case "status":
			err = migration.RunStatus(ctx, database)
		default:
			syslog.Fatalln("usage: migrate up|down [steps]|status")
		}
		if err != nil {
			syslog.Fatalf("migrate: %v\n", err)
		}
		return true
	case "seed":
		if err := seed.Run(ctx, database, flag.Arg(1)); err != nil {
			syslog.Fatalf("seed: %v\n", err)
		}
		return true
	}

	return false
}
