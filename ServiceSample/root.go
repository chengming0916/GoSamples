package main

import (
	"GoSamples/ServiceSample/internal/pkg/config"
	"GoSamples/ServiceSample/internal/pkg/config/types"
	"GoSamples/ServiceSample/internal/pkg/service"
	"GoSamples/ServiceSample/pkg/util/version"
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	showVersion bool
	cfg         types.ServiceConfig
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file of sample service")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "version of sample service")

	config.RegisterConfigFlags(rootCmd, &cfg)
}

var rootCmd = &cobra.Command{
	Use:   "Sample service",
	Short: "Sample service",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Println(version.Full())
			return nil
		}

		var (
			svrCfg *types.ServiceConfig
			err    error
		)

		if cfgFile != "" {
			svrCfg, err = config.LoadConfig(cfgFile)
			if err != nil {
				logrus.Errorln("faild to load config: ", err)
				os.Exit(1)
			}
		} else {
			cfg.Complete()
			svrCfg = &cfg
		}

		warning, err := config.ValidateConfig(svrCfg)
		if warning != nil {
			logrus.Warnf("WARNING: %v\n", warning)
		}
		if err != nil {
			logrus.Fatalln("faild to validate config: ", err)
			os.Exit(1)
		}

		if err := runServer(svrCfg); err != nil {
			logrus.Fatalln("faild to start service: ", err)
			os.Exit(1)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runServer(cfg *types.ServiceConfig) (err error) {

	if cfgFile != "" {
		logrus.Infof("uses config file: %s\n", cfgFile)
	} else {
		logrus.Infof("uses command line arguments for config")
	}

	svr, err := service.NewService(cfg)
	if err != nil {
		return err
	}
	svr.Run(context.Background())
	logrus.Infof("service started")
	return
}
