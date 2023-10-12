package cmd

import (
	"github.com/blackhorseya/monorepo-go/adapter/user/cmd/cronjob"
	"github.com/blackhorseya/monorepo-go/adapter/user/cmd/grpc"
	"github.com/blackhorseya/monorepo-go/adapter/user/cmd/restful"
	"github.com/spf13/cobra"
)

// startCmd represents the start command.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a user service",
}

func init() {
	startCmd.AddCommand(startAPICmd)
	startCmd.AddCommand(startGrpcCmd)
	startCmd.AddCommand(startCronjobCmd)

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var startAPICmd = &cobra.Command{
	Use:   "api",
	Short: "start a user api service",
	Run: func(cmd *cobra.Command, args []string) {
		service, err := restful.New()
		cobra.CheckErr(err)

		err = service.Start()
		cobra.CheckErr(err)

		err = service.AwaitSignal()
		cobra.CheckErr(err)
	},
}

var startGrpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "start a user grpc service",
	Run: func(cmd *cobra.Command, args []string) {
		service, err := grpc.New()
		cobra.CheckErr(err)

		err = service.Start()
		cobra.CheckErr(err)

		err = service.AwaitSignal()
		cobra.CheckErr(err)
	},
}

var startCronjobCmd = &cobra.Command{
	Use:   "cronjob",
	Short: "start a user cronjob service",
	Run: func(cmd *cobra.Command, args []string) {
		service, err := cronjob.New()
		cobra.CheckErr(err)

		err = service.Start()
		cobra.CheckErr(err)

		err = service.AwaitSignal()
		cobra.CheckErr(err)
	},
}
