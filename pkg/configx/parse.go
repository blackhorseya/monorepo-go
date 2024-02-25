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
		home, _ := os.UserHomeDir()
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

	err = v.BindEnv("lineNotify.endpoint", "LINE_NOTIFY_ENDPOINT")
	if err != nil {
		return err
	}

	err = v.BindEnv("lineNotify.accessToken", "LINE_NOTIFY_ACCESS_TOKEN")
	if err != nil {
		return err
	}

	err = bindEnvForOrianna(v)
	if err != nil {
		return err
	}

	err = bindEnvForSion(v)
	if err != nil {
		return err
	}

	err = bindEnvForReURL(v)
	if err != nil {
		return err
	}

	err = bindEnvForIRent(v)
	if err != nil {
		return err
	}

	return nil
}

func bindEnvForOrianna(v *viper.Viper) (err error) {
	err = v.BindEnv("orianna.http.port", "ORIANNA_HTTP_PORT")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.http.model", "ORIANNA_HTTP_MODE")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.storage.mongodb.dsn", "ORIANNA_STORAGE_MONGODB_DSN")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.storage.influxdb.url", "ORIANNA_STORAGE_INFLUXDB_URL")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.storage.influxdb.token", "ORIANNA_STORAGE_INFLUXDB_TOKEN")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.mq.kafka.brokers", "ORIANNA_MQ_KAFKA_BROKERS")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.mq.kafka.username", "ORIANNA_MQ_KAFKA_USERNAME")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.mq.kafka.password", "ORIANNA_MQ_KAFKA_PASSWORD")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.linebot.secret", "ORIANNA_LINEBOT_SECRET")
	if err != nil {
		return err
	}

	err = v.BindEnv("orianna.linebot.token", "ORIANNA_LINEBOT_TOKEN")
	if err != nil {
		return err
	}

	return nil
}

func bindEnvForReURL(v *viper.Viper) (err error) {
	err = v.BindEnv("reurl.http.mode", "REURL_HTTP_MODE")
	if err != nil {
		return err
	}

	err = v.BindEnv("reurl.linebot.secret", "REURL_LINEBOT_SECRET")
	if err != nil {
		return err
	}

	err = v.BindEnv("reurl.linebot.token", "REURL_LINEBOT_TOKEN")
	if err != nil {
		return err
	}

	err = v.BindEnv("reurl.storage.redis.addr", "REURL_STORAGE_REDIS_ADDR")
	if err != nil {
		return err
	}

	err = v.BindEnv("reurl.storage.redis.password", "REURL_STORAGE_REDIS_PASSWORD")
	if err != nil {
		return err
	}

	err = v.BindEnv("reurl.storage.redis.db", "REURL_STORAGE_REDIS_DB")
	if err != nil {
		return err
	}

	return nil
}

func bindEnvForSion(v *viper.Viper) (err error) {
	err = v.BindEnv("sion.http.url", "SION_HTTP_URL")
	if err != nil {
		return err
	}

	err = v.BindEnv("sion.http.mode", "SION_HTTP_MODE")
	if err != nil {
		return err
	}

	err = v.BindEnv("sion.linebot.secret", "SION_LINEBOT_SECRET")
	if err != nil {
		return err
	}

	err = v.BindEnv("sion.linebot.token", "SION_LINEBOT_TOKEN")
	if err != nil {
		return err
	}

	return nil
}

func bindEnvForIRent(v *viper.Viper) (err error) {
	err = v.BindEnv("irent.version", "IRENT_VERSION")
	if err != nil {
		return err
	}

	err = v.BindEnv("irent.http.url", "IRENT_HTTP_URL")
	if err != nil {
		return err
	}

	return nil
}
