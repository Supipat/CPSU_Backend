package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort          string
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseSSLMode  string

	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	S3BucketName       string
}

func LoadConfig() (Config, error) {

	_ = godotenv.Load()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set default values
	viper.SetDefault("POSTGRES.HOST", "localhost")
	viper.SetDefault("POSTGRES.PORT", 5432)
	viper.SetDefault("POSTGRES.USER", "postgres")
	viper.SetDefault("POSTGRES.PASSWORD", "")
	viper.SetDefault("POSTGRES.DBNAME", "cpsu")
	viper.SetDefault("POSTGRES.SSLMODE", "disable")

	viper.SetDefault("AWS.REGION", "ap-southeast-2")
	viper.SetDefault("S3.BUCKET_NAME", "cpsu-website")

	// Set config values
	config := Config{
		AppPort:            viper.GetString("APP.PORT"),
		DatabaseHost:       viper.GetString("POSTGRES.HOST"),
		DatabasePort:       viper.GetInt("POSTGRES.PORT"),
		DatabaseUser:       viper.GetString("POSTGRES.USER"),
		DatabasePassword:   viper.GetString("POSTGRES.PASSWORD"),
		DatabaseName:       viper.GetString("POSTGRES.DBNAME"),
		DatabaseSSLMode:    viper.GetString("POSTGRES.SSLMODE"),
		AWSRegion:          viper.GetString("AWS.REGION"),
		AWSAccessKeyID:     viper.GetString("AWS.ACCESS_KEY_ID"),
		AWSSecretAccessKey: viper.GetString("AWS.SECRET_ACCESS_KEY"),
		S3BucketName:       viper.GetString("S3.BUCKET_NAME"),
	}

	return config, nil
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseName,
		c.DatabaseSSLMode)
}
