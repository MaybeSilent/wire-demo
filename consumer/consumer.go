package consumer

import "github.com/maybesilent/wire-demo/conf"

type Consumer struct {
	consumer string
}

func InitConsumer(Config *conf.Config) *Consumer {
	return &Consumer{consumer: Config.GetConf()}
}
