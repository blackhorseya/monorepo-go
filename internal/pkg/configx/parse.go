package configx

import (
	"github.com/spf13/viper"
)

var (
	C = new(Config)
)

// Load loads config from path.
func Load(path string) error {
	v := viper.GetViper()
	v.SetConfigFile(path)
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
