package config

import "github.com/kelseyhightower/envconfig"

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

func NewDBConfig() (DBConfig, error) {
	dbConfig := DBConfig{}
	err := envconfig.Process("POSTGRES", &dbConfig)
	if err != nil {
		return DBConfig{}, err
	}
	return dbConfig, nil
}
