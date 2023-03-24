package config

import (
	"github.com/jinzhu/configor"
)

func Initialize(fileName string) {
	err := configor.Load(&Config, fileName)
	if err != nil {
		panic("Unable to read configurations")
	}
}

var Config = struct {
	Clickhouse struct {
		Dsn      string `env:"CLICK_HOUSE_URL" yaml:"dsn" default:"clickhouse:9000"`
		Database string `env:"CLICKHOUSE_DATABASE" yaml:"database" default:"default"`
	}
	Kafka struct {
		BootstrapServers    string `env:"KAFKA_BOOTSTRAP_SERVERS" yaml:"bootstrapServers" default:"localhost:9094"`
		KAFKA_METRICS_TOPIC string `env:"KAFKA_METRICS_TOPIC" yaml:"metricsTopic" default:"metrics"`
	}
	Database struct {
		Name string `env:"DATABASE_NAME" yaml:"name" defaults:"clickhouse"`
	}
}{}
