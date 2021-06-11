package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"gogo_blueprint"`

	// MongoDB config
	MongoDBEndpoint         string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://touch:touchja@localhost:27017"`
	MongoDBName             string `env:"MONGODB_NAME" envDefault:"superhero"`
	MongoDBHeroTableName string `env:"MONGODB_HERO_TABLE_NAME" envDefault:"lists"`

	// Jaeger config
	JaegerAgentHost string `env:"JAEGER_AGENT_HOST" envDefault:"localhost"`
	JaegerAgentPort string `env:"JAEGER_AGENT_PORT" envDefault:"6831"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
