package cmd

import (
	"github.com/blackhorseya/monorepo-go/adapter/stringx/cmd/cronjob"
	"github.com/blackhorseya/monorepo-go/adapter/stringx/cmd/grpcserver"
	"github.com/blackhorseya/monorepo-go/adapter/stringx/cmd/restful"
	"github.com/blackhorseya/monorepo-go/internal/pkg/cmdx"
	"github.com/spf13/cobra"
)

// startCmd represents the start command.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a user service",
}

func init() {
	startCmd.AddCommand(cmdx.NewServiceCmd("api", "start a user api service", restful.New))
	startCmd.AddCommand(cmdx.NewServiceCmd("grpc", "start a user grpc service", grpcserver.New))
	startCmd.AddCommand(cmdx.NewServiceCmd("cronjob", "start a user cronjob service", cronjob.New))

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
