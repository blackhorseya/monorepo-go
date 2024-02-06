package configx

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	C = new(Config)

	A = new(Application)
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

	err := bindEnv(v)
	if err != nil {
		return err
	}

	err = v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return err
	}

	err = v.Unmarshal(&C)
	if err != nil {
		return err
	}

	return nil
}

// ReplaceApplication replaces the application.
func ReplaceApplication(app Application) {
	A = &app
}

func bindEnv(v *viper.Viper) error {
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var err error
	err = v.BindEnv("log.level", "LOG_LEVEL")
	if err != nil {
		return err
	}

	err = v.BindEnv("log.format", "LOG_FORMAT")
	if err != nil {
		return err
	}

	err = v.BindEnv("finmind.http.url", "FINMIND_HTTP_URL")
	if err != nil {
		return err
	}

	err = v.BindEnv("finmind.token", "FINMIND_TOKEN")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.storage.mongodb.dsn", "ORIANNA_STORAGE_MONGODB_DSN")
	if err != nil {
		return err
	}

	err = v.BindEnv("lineNotify.endpoint", "LINE_NOTIFY_ENDPOINT")
	if err != nil {
		return err
	}

	err = v.BindEnv("lineNotify.accessToken", "LINE_NOTIFY_ACCESS_TOKEN")
	if err != nil {
		return err
	}

	return nil
}
