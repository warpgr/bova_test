package app

import (
	"context"
	"net/http"
	"time"

	"github.com/warpgr/bova_test/internal/configs"
	"github.com/warpgr/bova_test/internal/controller"
	"github.com/warpgr/bova_test/internal/repository"
	"github.com/warpgr/bova_test/internal/service"
	"github.com/warpgr/bova_test/pkg/daemons"
	"github.com/warpgr/bova_test/pkg/exchanges"
	"github.com/warpgr/bova_test/pkg/store"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type routeHandler interface {
	Register(r *gin.RouterGroup)
}

type Application interface {
	Init() error
	Run(ctx context.Context) error
	Shutdown(ctx context.Context, shutDownServer bool)
}

func NewApplication(config configs.Configs) Application {
	return &application{
		config:           config,
		handlers:         make([]routeHandler, 0, 1),
		priceListStorage: store.NewKVMapStorage[string, float64](100),
	}
}

type application struct {
	config   configs.Configs
	handlers []routeHandler
	server   *http.Server

	priceListStorage    store.KVStorage[string, float64]
	priceProvider       daemons.PriceProvider
	cancelPriceProvider context.CancelFunc
}

func (a *application) Init() error {
	log.Info("Initializing application.")
	router := gin.Default()
	api := router.Group("/api/v1")

	c := controller.NewPriceList(
		service.NewPriceList(
			repository.NewPriceList(a.priceListStorage)))

	a.handlers = append(a.handlers, c)

	for _, c := range a.handlers {
		c.Register(api)
	}

	a.priceProvider = daemons.NewPriceProvider(
		exchanges.NewKrakenExchange(a.config.KrakenEndpoint), a.priceListStorage)

	a.server = &http.Server{
		Addr:    a.config.Endpoint,
		Handler: router,
	}
	return nil
}

func (a *application) Run(ctx context.Context) error {
	log.Info("Running application.")

	daemonCtx, cancelDaemon := context.WithCancel(ctx)
	go a.priceProvider.Run(daemonCtx, time.Minute)
	a.cancelPriceProvider = cancelDaemon

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Errorf("Error occurs when trying start service. Error: %v.", err)
			a.Shutdown(ctx, false)
		}
	}()

	return nil
}

func (a *application) Shutdown(ctx context.Context, shutDownServer bool) {
	log.Info("Shutting down application.")

	a.cancelPriceProvider()

	if shutDownServer {
		log.Warnf("Shutting down server")
		if err := a.server.Shutdown(ctx); err != nil {
			log.Errorf("Error occurs when trying to shutdown server. Error: %v.", err)
		}
	}
}
