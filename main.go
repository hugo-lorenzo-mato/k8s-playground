package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"k8s-playground/cmd"
	"k8s-playground/conf"
	timeUtil "k8s-playground/util/time"
	"k8s-playground/util/version"
	"os"
	"time"
)

var debug bool

func init() {
	log.SetOutput(os.Stdout)
	if os.Getenv("DEBUG") != "" {
		log.SetFormatter(&log.TextFormatter{
			ForceColors:               false,
			DisableColors:             false,
			ForceQuote:                false,
			DisableQuote:              true,
			EnvironmentOverrideColors: false,
			DisableTimestamp:          false,
			FullTimestamp:             false,
			TimestampFormat:           "",
			DisableSorting:            false,
			SortingFunc:               nil,
			DisableLevelTruncation:    false,
			PadLevelText:              false,
			QuoteEmptyFields:          false,
			FieldMap:                  nil,
			CallerPrettyfier:          nil,
		})
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	defer timeUtil.FromStart(time.Now(), "Cli Command Execution")
	version.Lookup()
	app := cli.NewApp()
	app.Version = conf.VERSION
	app.Usage = "This CLI brings together some personal tests using the K8S and OCP libraries."
	app.Commands = cmd.Cmds
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Main:  %s", err)
	}
}
