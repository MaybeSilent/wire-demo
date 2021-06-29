package model

import "github.com/maybesilent/wire-demo/conf"

type Model struct {
	model string
}

func InitModel(Config *conf.Config) *Model {
	return &Model{model: Config.GetConf()}
}