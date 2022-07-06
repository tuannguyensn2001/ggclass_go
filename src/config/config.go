package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type Config struct {
	db        *gorm.DB
	port      string
	dbUrl     string
	secretKey string
}

var Cfg Config

func (c *Config) GetPort() string {
	return c.port
}

func (c *Config) GetDB() *gorm.DB {
	return c.db
}

func (c *Config) SecretKey() string {
	return c.secretKey
}

func (c *Config) GetDBUrl() string {
	return c.dbUrl
}

func Load() error {
	path, _ := os.Getwd()

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	dbUrl := viper.GetString("DB_URL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	result := &Config{
		dbUrl: dbUrl,
		port:  viper.GetString("PORT"),
		db:    db,
	}

	Cfg = *result

	return nil

}
