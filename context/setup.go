package context

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

var (
	cfg = ConfigInit(LoadConfiguration())
)

type Context struct {
	Config *AppConfig
}

func Init() {
}

func GetContext() *Context {
	return &Context{
		Config: cfg,
	}
}

func LoadConfiguration() *viper.Viper {
	v := viper.New()

	v.AutomaticEnv()
	config := api.DefaultConfig()
	config.Address = v.GetString("CONSUL_HOST")
	config.Token = v.GetString("CONSUL_TOKEN")

	c, err := api.NewClient(config)
	if err != nil {
		fmt.Println("cant connect to consul")
		panic(err)
	}

	pair, _, err := c.KV().Get(v.GetString("CONSUL_PATH"), nil)
	if err != nil {
		fmt.Println("invalid consul config")
		panic(err)
	}

	if pair == nil || string(pair.Value) == "" {
		panic("received null configuration from consul")
	}

	v.SetConfigType("JSON")
	v.ReadConfig(bytes.NewBuffer(pair.Value))

	return v
}
