// +build wireinject

package init

import (
	"context"
	"github.com/google/wire"
	"github.com/maybesilent/wire-demo/conf"
	"github.com/maybesilent/wire-demo/consumer"
)

func InitWithDi(ctx context.Context) (*consumer.Consumer, error) {
	wire.Build(conf.InitConfig, consumer.InitConsumer)
	return &consumer.Consumer{}, nil
}




