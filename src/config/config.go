package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/pusher/pusher-http-go"
	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Config struct {
	db        *gorm.DB
	port      string
	dbUrl     string
	secretKey string
	pusher    pusher.Client
	rabbitMQ  *amqp091.Connection
	rds       *redis.Client
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

func (c *Config) GetPusher() pusher.Client {
	return c.pusher
}

func (c *Config) GetRabbitMQ() *amqp091.Connection {
	return c.rabbitMQ
}

func (c *Config) GetRedis() *redis.Client {
	return c.rds
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

	pusherClient := pusher.Client{
		AppID:   "1440558",
		Key:     "26bdb6fd186156c41fe6",
		Secret:  "708f13f675065ba00a92",
		Cluster: "ap1",
		Secure:  true,
	}

	result := &Config{
		dbUrl:     dbUrl,
		port:      viper.GetString("PORT"),
		db:        db,
		pusher:    pusherClient,
		rabbitMQ:  connectRabbitMq(),
		secretKey: viper.GetString("APP_KEY"),
		rds: redis.NewClient(&redis.Options{
			Addr:     "redis-17404.c299.asia-northeast1-1.gce.cloud.redislabs.com:17404",
			Password: "oVzG4E5NyOWCLaYU1II0021uR6rwj2yp",
		}),
	}

	Cfg = *result

	return nil

}

func connectRabbitMq() *amqp091.Connection {
	conn, err := amqp091.Dial(viper.GetString("RABBITMQ_URL"))
	if err != nil {
		log.Fatalln("connect failed to rabbit")
	}
	return conn
}
