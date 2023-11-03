package cmdx

import (
	"github.com/blackhorseya/monorepo-go/pkg/adapterx"
	"github.com/spf13/viper"
)

// ServiceCmd represents the service command.
type ServiceCmd struct {
	Use   string
	Short string
	Run   func(v *viper.Viper) (adapterx.Servicer, error)
}
