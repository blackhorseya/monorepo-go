package cmdx

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServiceCmd represents the service command.
type ServiceCmd struct {
	Use   string
	Short string
	Run   func(v *viper.Viper) (adapterx.Servicer, error)
}

func (s *ServiceCmd) NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   s.Use,
		Short: s.Short,
		Run: func(cmd *cobra.Command, args []string) {
			v := viper.GetViper()

			service, err := s.Run(v)
			cobra.CheckErr(err)

			err = service.Start()
			cobra.CheckErr(err)

			err = service.AwaitSignal()
			cobra.CheckErr(err)
		},
	}
}
