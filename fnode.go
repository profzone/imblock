package main

import (
	"github.com/profzone/imblock/core/castle"
	"github.com/spf13/viper"
	"fmt"
	"github.com/sirupsen/logrus"
	"flag"
	"github.com/profzone/imblock/global"
	"os"
	"github.com/profzone/imblock/persistence_service"
	"github.com/profzone/imblock/api_service"
	"os/signal"
	"github.com/profzone/imblock/network_service"
)

var (
	configFilePath string
	configGroup    string
)

func main() {
	flag.StringVar(&configGroup, "g", "", "modify the group of app")
	flag.StringVar(&configFilePath, "c", "", "the config file path")
	flag.Parse()

	initConfig()
	global.InitConfig(configGroup)

	stack := castle.NewStack("fnode")
	stack.RegisterService(persistence_service.NewPersistenceBoltDBServiceBootstrap)
	stack.RegisterService(network_service.NewNetworkDHTServiceBootstrap)
	stack.RegisterService(api_service.NewApiHttpServiceBootstrap)

	stack.Start()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	logrus.Info("system shutting down")
	if err := stack.Stop(); err == nil {
		logrus.Info("system shutdown.")
	}
}

func initConfig() {
	var configFileName = ".imblock"
	if configGroup != "" {
		configFileName += fmt.Sprint(".", configGroup)
	}
	if configFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFilePath)
	} else {
		// Search config in home directory with name ".chain" (without extension).
		viper.AddConfigPath("./")
		viper.SetConfigName(configFileName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	env := viper.GetString("ENV")
	switch env {
	case "LOCAL":
		logrus.SetLevel(logrus.DebugLevel)
		break
	case "PROD":
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetOutput(os.Stdout)
	logrus.AddHook(ContextHook{})
}
