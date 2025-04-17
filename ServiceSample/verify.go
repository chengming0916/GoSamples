package main

import (
	"GoSamples/ServiceSample/internal/pkg/config"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(verifyCmd)
}

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify config is valid",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile == "" {
			logrus.Errorln("the config file is not specified")
			return nil
		}

		svrCfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			logrus.Errorln("load config error: ", err)
			os.Exit(1)
		}

		warning, err := config.ValidateConfig(svrCfg)
		if warning != nil {
			logrus.Warnf("WARNING: %v \n", warning)
		}

		if err != nil {
			logrus.Errorln("valid config error: ", err)
			os.Exit(1)
		}

		logrus.Infof("the config file %s syntax is ok \n", cfgFile)

		return nil
	},
}
