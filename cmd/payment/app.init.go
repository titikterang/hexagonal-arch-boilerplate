package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/payment/adapter/handler"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/payment/adapter/repository"
	"github.com/titikterang/hexagonal-arch-boilerplate/internal/payment/core/services"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/config"
)

func initConfig(configName string) (*config.Config, error) {
	cfg := &config.Config{}
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	fmt.Printf("%#v\n", cfg)
	return cfg, err
}

func initHandler(cfg *config.Config) (*handler.Handler, error) {
	repo := repository.NewRepository(cfg)
	hdl, err := handler.NewHandler(cfg, services.NewService(cfg, repo))
	if err != nil {
		return nil, err
	}

	return hdl, nil
}
