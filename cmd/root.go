package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var RootCmd = &cobra.Command{
	Use:   "prometheus_clickhouse_exporter",
	Short: "prometheus_clickhouse_exporter",
	Long:  `prometheus_clickhouse_exporter`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.BindEnv("clickhouse.connection")
	viper.BindEnv("production")

	RootCmd.PersistentFlags().StringP("clickhouse-connection", "", "http://localhost:8123", "list of clickhouse connections")
	RootCmd.PersistentFlags().StringP("listen", "", "0.0.0.0:8000", "listen address")
	RootCmd.PersistentFlags().StringP("prometheus-namespace", "N", "obi8", "prometheus namespace")
	RootCmd.PersistentFlags().StringP("prometheus-subsystem", "S", "clickhouse", "prometheus subsystem")

	RootCmd.PersistentFlags().Bool("production", false, "Run in production mode")

	viper.BindPFlag("production", RootCmd.PersistentFlags().Lookup("production"))
	viper.BindPFlag("help", RootCmd.PersistentFlags().Lookup("help"))
	viper.BindPFlag("listen", RootCmd.PersistentFlags().Lookup("listen"))
	viper.BindPFlag("prometheus.namespace", RootCmd.PersistentFlags().Lookup("prometheus-namespace"))
	viper.BindPFlag("prometheus.subsystem", RootCmd.PersistentFlags().Lookup("prometheus-subsystem"))

	viper.BindPFlag("clickhouse.connection", RootCmd.PersistentFlags().Lookup("clickhouse-connection"))
	viper.SetDefault("clickhouse.connection", "http://localhost:8123")
	viper.SetDefault("production", false)
}
