package cmd

import (
	"github.com/blackhorseya/monorepo-go/adapter/user/cmd/cronjob"
	"github.com/blackhorseya/monorepo-go/adapter/user/cmd/grpc"
	"github.com/blackhorseya/monorepo-go/adapter/user/cmd/restful"
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// ServiceCmd represents the service command.
type ServiceCmd struct {
	Use   string
	Short string
	Run   func(v *viper.Viper, logger *zap.Logger) adapterx.Servicer
}

// NewServiceCmd creates a new service command.
func NewServiceCmd(
	use string,
	short string,
	run func(v *viper.Viper, logger *zap.Logger) adapterx.Servicer,
) *cobra.Command {
	return (&ServiceCmd{Use: use, Short: short, Run: run}).NewCmd()
}

// NewCmd creates a new service command.
func (s *ServiceCmd) NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   s.Use,
		Short: s.Short,
		Run: func(cmd *cobra.Command, args []string) {
			v := viper.GetViper()
			logger := zap.NewExample()

			service := s.Run(v, logger)

			err := service.Start()
			cobra.CheckErr(err)

			err = service.AwaitSignal()
			cobra.CheckErr(err)
		},
	}

	return cmd
}

// startCmd represents the start command.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a user service",
}

func init() {
	startCmd.AddCommand(NewServiceCmd("api", "start a user api service", restful.New))
	startCmd.AddCommand(NewServiceCmd("grpc", "start a user grpc service", grpc.New))
	startCmd.AddCommand(NewServiceCmd("cronjob", "start a user cronjob service", cronjob.New))

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
