package cmd

import (
	"time"

	"github.com/blackhorseya/monorepo-go/adapter/redpacket/cmd/cronjob"
	"github.com/blackhorseya/monorepo-go/adapter/redpacket/cmd/restful"
	"github.com/blackhorseya/monorepo-go/internal/pkg/cmdx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultInterval = 5 * time.Second
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a service",
}

func init() {
	startCmd.AddCommand(cmdx.NewServiceCmd("api", "start a api service", restful.New))

	cronjobCmd := cmdx.NewServiceCmd("cronjob", "start a cronjob service", cronjob.New)
	cronjobCmd.Flags().Duration("interval", defaultInterval, "the interval of cronjob")
	_ = viper.BindPFlag("interval", cronjobCmd.Flags().Lookup("interval"))
	startCmd.AddCommand(cronjobCmd)

	rootCmd.AddCommand(startCmd)
}
