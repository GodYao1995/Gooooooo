package xhttp

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/cd-home/Goooooo/pkg/config"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func TestNewServer(t *testing.T) {
	var newVp func() *viper.Viper = func() *viper.Viper {
		return config.NewViper("admin", "dev", "../../config/testdata/configs")
	}
	app := fx.New(
		fx.Provide(
			newVp,
		),
		fx.Invoke(New),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
