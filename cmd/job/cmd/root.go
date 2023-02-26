package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/cd-home/Goooooo/internal/pkg/errno"
	"github.com/cd-home/Goooooo/pkg/config"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var configModule = fx.Provide(NewViper)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root",
	Long:  "root",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Exeute() {
	cobra.CheckErr(rootCmd.Execute())
}

func NewViper() *viper.Viper {
	_differentAppAndEnvironment(&app, &mode)
	return config.NewViper(app, mode, configs...)
}

func _differentAppAndEnvironment(app, mode *string) {
	_app, _mode := os.Getenv("APP_NAME"), os.Getenv("APP_MODE")
	if len(strings.TrimSpace(_app)) != 0 && len(strings.TrimSpace(_mode)) != 0 {
		*app = _app
		*mode = _mode
	}

	if len(strings.TrimSpace(*app)) == 0 || len(strings.TrimSpace(*mode)) == 0 {
		log.Fatalln(errno.ErrorNotFoundService)
	}
	// just for testing prod mode
	if *mode == "prod" {
		godotenv.Load()
	}
}
