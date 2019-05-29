package context

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ConsulHost  string
	DBString    string
	ServiceName string
}

func ConfigInit(v *viper.Viper) *AppConfig {
	return &AppConfig{
		ConsulHost:  v.GetString("CONSUL_HOST"),
		DBString:    v.GetString("db"),
		ServiceName: v.GetString("service_name"),
	}
}
