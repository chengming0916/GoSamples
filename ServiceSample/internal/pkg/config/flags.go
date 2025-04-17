package config

import (
	"GoSamples/ServiceSample/internal/pkg/config/types"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

func RegisterConfigFlags(cmd *cobra.Command, c *types.ServiceConfig) {
	cmd.PersistentFlags().StringVarP(&c.TcpCfg.Host, "tcp_host", "", "0.0.0.0", "tcp service bind address")
	cmd.PersistentFlags().IntVarP(&c.TcpCfg.Port, "tcp_port", "", 8000, "tcp service bind port")
}
