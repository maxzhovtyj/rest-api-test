package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"rest-api-test/pkg/logging"
	"sync"
)

const (
	dbDatabase     = "db.database"
	dbCollection   = "db.collection"
	dbUsername     = "db.username"
	portName       = "PORT"
	ipName         = "BIND_IP"
	dbPasswordName = "PASSWORD_DB"
	redisHostName  = "redis.host"
	redisPortName  = "redis.port"
)

type Redis struct {
	Host string `yml:"host"`
	Port string `yml:"port"`
}

type Config struct {
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Port       string `env:"PORT"`
	BindIp     string `env:"BIND_IP"`
	Redis
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()

		logger.Info("initializing .yml file")
		if err := initConfig(); err != nil {
			logger.Panic("panic while initializing .yml file")
			panic(err)
		}

		logger.Info("initializing .env file")
		if err := godotenv.Load(); err != nil {
			logger.Panic("panic while initializing .env file")
			panic(err)
		}

		redisInstance := Redis{
			Host: viper.GetString(redisHostName),
			Port: viper.GetString(redisPortName),
		}
		instance = &Config{
			Database:   viper.GetString(dbDatabase),
			Collection: viper.GetString(dbCollection),
			Username:   viper.GetString(dbUsername),
			Password:   os.Getenv(dbPasswordName),
			BindIp:     os.Getenv(ipName),
			Port:       os.Getenv(portName),
			Redis:      redisInstance,
		}
	})

	return instance
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
