package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	TokenSymetricKey     string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	AwsRegion            string        `mapstructure:"AWS_REGION"`
	AwsS3Bucket          string        `mapstructure:"AWS_S3_BUCKET"`
	AwsAccessKeyID       string        `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretKeyID       string        `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
