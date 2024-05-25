package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/warpgr/bova_test/internal/app"
	"github.com/warpgr/bova_test/internal/configs"

	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := configs.LoadConfigs()
	if err != nil {
		log.Fatalf("Error occurs when trying to load configurations. Error: %v.", err)
	}

	application := app.NewApplication(*config)

	if err := application.Init(); err != nil {
		log.Fatalf("Error occurs when trying to initialize application. Error: %v.", err)
	}

	if err := application.Run(context.Background()); err != nil {
		log.Fatalf("Error occurs when trying to run application. Error: %v.", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	log.Warnf("Signal received. Shutting down all.")

	application.Shutdown(context.Background(), true)
}
