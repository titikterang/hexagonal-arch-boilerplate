package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/viper"
	"github.com/titikterang/hexagonal-arch-boilerplate/lib/config"
)

func initConfig() (*Config, error) {
	cfg := &model.Config{}
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	return cfg, err
}

func InitRepo(cfg *config.Config) (usecase.LoanRepoInterface, error) {
	// init db conn
	log.Info("Initialize database connection")
	dbConn, err := database.NewConnection(cfg.DBConfig)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	db, err := repository.New(cfg, dbConn)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return db, nil
}

func InitUseCase(cfg *model.Config) (handler.LoanUseCaseInterface, error) {
	repo, err := InitRepo(cfg)
	if err != nil {
		log.Errorf("Failed to Initialize repository, %#v", err)
		return nil, err
	}
	return usecase.New(cfg, repo), nil
}
