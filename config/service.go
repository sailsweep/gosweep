package config

import (
	"context"
	myconfig "entdemo/config/types"
)

func GetConfig(ctx context.Context) (context.Context, error) {
	var config myconfig.Config
	if err := config.InitSetting(); err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, "config", config), nil
}

func NewConfig() (myconfig.Config, error) {
	var config myconfig.Config
	if err := config.InitSetting(); err != nil {
		return config, err
	}
	return config, nil
}
