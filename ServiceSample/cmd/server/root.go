package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	CfgFileTypeYaml = iota
	CfgFileTypeCmd
)

var (
	cfgFile     string
	showVersion bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file of sample service")
}

var rootCmd = &cobra.Command{
	Use:   "Sample service",
	Short: "Sample service with dashboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Println(version.Full())
			return nil
		}

		var cfg config.ServerConfig
		var err error
		if cfgFile != "" {
			var content []byte
			content, err = config.GetRenderedConfigFromFile(cfgFile)
			if err != nil {
				return err
			}
			cfg, err = parseServerConfig(CfgFileTypeYaml, content)
		} else {
			cfg, err = parseServerConfig(CfgFileTypeCmd, nil)
		}
		if err != nil {
			return err
		}

		err = runServer(cfg)
		if err != nil {
			fmt.Println(err)
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
