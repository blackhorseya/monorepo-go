package configx

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	C = new(Config)
)

// Load loads config from path.
func Load(path string, name string) error {
	v := viper.GetViper()

	if path != "" {
		v.SetConfigFile(path)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		v.AddConfigPath(home)
		v.SetConfigType("yaml")
		v.SetConfigName("." + name)
	}

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())

	err = v.Unmarshal(&C)
	if err != nil {
		return err
	}

	return nil
}
