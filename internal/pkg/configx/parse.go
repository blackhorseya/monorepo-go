package configx

import (
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
		v.SetConfigName(name)
	}

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&C)
	if err != nil {
		return err
	}

	return nil
}

// LoadWithViper loads config from viper.
func LoadWithViper(v *viper.Viper) error {
	err := v.Unmarshal(&C)
	if err != nil {
		return err
	}

	return nil
}
