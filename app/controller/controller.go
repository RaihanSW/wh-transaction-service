package controller

import (
	"dc-transaction-service/app/controller/handler"
	"dc-transaction-service/app/controller/handler/middleware"
	"dc-transaction-service/app/model/constant"
	"dc-transaction-service/config"
	"dc-transaction-service/docs"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Controller struct {
	transactionHandler       handler.TransactionHandlerInterface
	transactionStatusHandler handler.TransactionStatusHandlerInterface
	goodsHandler             handler.GoodsHandlerInterface
	goodsStatusHandler       handler.GoodsStatusHandlerInterface
	goodsCategoryHandler     handler.GoodsCategoryHandlerInterface
	generalHandler           handler.GeneralHandlerInterface
	cfg                      *config.ApplicationConfig
	env                      string
}

func NewController(
	transactionHandler handler.TransactionHandlerInterface,
	transactionStatusHandler handler.TransactionStatusHandlerInterface,
	goodsHandler handler.GoodsHandlerInterface,
	goodsStatusHandler handler.GoodsStatusHandlerInterface,
	goodsCategoryHandler handler.GoodsCategoryHandlerInterface,
	generalHandler handler.GeneralHandlerInterface,
	cfg *config.ApplicationConfig,
	env string) *Controller {

	return &Controller{
		transactionHandler:       transactionHandler,
		transactionStatusHandler: transactionStatusHandler,
		goodsHandler:             goodsHandler,
		goodsStatusHandler:       goodsStatusHandler,
		goodsCategoryHandler:     goodsCategoryHandler,
		generalHandler:           generalHandler,
		cfg:                      cfg,
		env:                      env,
	}
}

func (c *Controller) StartRoute() *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.Middleware())

	r.Get("/ping", middleware.SetRequestID(c.generalHandler.Ping))
	r.Get("/health/check", middleware.SetRequestID(c.generalHandler.HealthCheck))
	r.Get("/health/check-db", middleware.SetRequestID(c.generalHandler.HealthCheckDB))

	// Init Swagger with swaggo
	docs.SwaggerInfo.Title = "DC Transaction Service"
	docs.SwaggerInfo.Description = fmt.Sprintf("DC Transaction Service with Golang - Environtment: %s", strings.ToLower(c.env))
	docs.SwaggerInfo.Version = "1.1"
	switch strings.ToLower(c.env) {
	case constant.EnvLocal:
		docs.SwaggerInfo.Schemes = []string{"http"}
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", c.cfg.Port.Service)
		docs.SwaggerInfo.BasePath = "/"
	case constant.EnvDev:
		docs.SwaggerInfo.Schemes = []string{"https"}
		docs.SwaggerInfo.Host = "dc21-dev-api.cinema21.co.id"
		docs.SwaggerInfo.BasePath = "/transaction/"
	case constant.EnvStaging:
		docs.SwaggerInfo.Schemes = []string{"https"}
		docs.SwaggerInfo.Host = "dc21-staging-api.cinema21.co.id"
		docs.SwaggerInfo.BasePath = "/transaction/"
	}

	if strings.ToLower(c.env) != constant.EnvProd {
		r.Mount("/swagger", httpSwagger.WrapHandler)
	}

	r.Route("/transaction", func(r chi.Router) {
		r.Use(middleware.GetLogMiddleware(*c.cfg))
		r.Post("/", c.transactionHandler.TransactionCreate)
		r.Get("/", c.transactionHandler.TransactionGetAll)
		r.Get("/{uid}", c.transactionHandler.TransactionGetByID)
		r.Put("/{uid}", c.transactionHandler.TransactionUpdate)
		r.Delete("/{uid}", c.transactionHandler.TransactionDelete)
		r.Route("/status", func(r chi.Router) {
			r.Post("/", c.transactionStatusHandler.TransactionStatusCreate)
			r.Get("/", c.transactionStatusHandler.TransactionStatusGetAll)
			r.Get("/{id}", c.transactionStatusHandler.TransactionStatusGetByID)
			r.Put("/{id}", c.transactionStatusHandler.TransactionStatusUpdate)
			r.Delete("/{id}", c.transactionStatusHandler.TransactionStatusDelete)
		})
	})

	r.Route("/goods", func(r chi.Router) {
		r.Use(middleware.GetLogMiddleware(*c.cfg))
		r.Post("/", c.goodsHandler.GoodsCreate)
		r.Get("/", c.goodsHandler.GoodsGetAll)
		r.Get("/{uid}", c.goodsHandler.GoodsGetByID)
		r.Put("/{uid}", c.goodsHandler.GoodsUpdate)
		r.Delete("/{uid}", c.goodsHandler.GoodsDelete)
		r.Route("/status", func(r chi.Router) {
			r.Post("/", c.goodsStatusHandler.GoodsStatusCreate)
			r.Get("/", c.goodsStatusHandler.GoodsStatusGetAll)
			r.Get("/{id}", c.goodsStatusHandler.GoodsStatusGetByID)
			r.Put("/{id}", c.goodsStatusHandler.GoodsStatusUpdate)
			r.Delete("/{id}", c.goodsStatusHandler.GoodsStatusDelete)
		})
		r.Route("/category", func(r chi.Router) {
			r.Post("/", c.goodsCategoryHandler.GoodsCategoryCreate)
			r.Get("/", c.goodsCategoryHandler.GoodsCategoryGetAll)
			r.Get("/{id}", c.goodsCategoryHandler.GoodsCategoryGetByID)
			r.Put("/{id}", c.goodsCategoryHandler.GoodsCategoryUpdate)
			r.Delete("/{id}", c.goodsCategoryHandler.GoodsCategoryDelete)
		})
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", c.cfg.Port.Service),
		WriteTimeout: c.cfg.Port.ServiceTimeout * time.Second,
		ReadTimeout:  c.cfg.Port.ServiceTimeout * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv
}
