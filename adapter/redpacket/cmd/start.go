package cmd

import (
	"github.com/blackhorseya/monorepo-go/adapter/redpacket/cmd/restful"
	"github.com/blackhorseya/monorepo-go/internal/pkg/cmdx"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a service",
}

func init() {
	startCmd.AddCommand(cmdx.NewServiceCmd("api", "start a api service", restful.New))

	rootCmd.AddCommand(startCmd)
}
