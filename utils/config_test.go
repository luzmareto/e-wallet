package utils

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	// Prepare the test environment
	viper.Reset()
	viper.SetEnvPrefix("TEST")
	viper.AutomaticEnv()

	// Set the environment variables for the test
	_ = viper.BindEnv("SERVER_ADDRESS")
	_ = viper.BindEnv("DB_DRIVER")
	_ = viper.BindEnv("DB_SOURCE")
	_ = viper.BindEnv("TOKEN_SYMETRIC_KEY")
	_ = viper.BindEnv("ACCESS_TOKEN_DURATION")
	_ = viper.BindEnv("REFRESH_TOKEN_DURATION")

	viper.Set("SERVER_ADDRESS", "localhost:8080")
	viper.Set("DB_DRIVER", "mysql")
	viper.Set("DB_SOURCE", "user:password@tcp(localhost:3306)/database")
	viper.Set("TOKEN_SYMETRIC_KEY", "your-secret-key")
	viper.Set("ACCESS_TOKEN_DURATION", "1h")
	viper.Set("REFRESH_TOKEN_DURATION", "24h")

	// Call the LoadConfig function
	config, err := LoadConfig("..")

	// Assert the result
	require.NoError(t, err)
	require.NotNil(t, config)
	require.Equal(t, "localhost:8080", config.ServerAddress)
	require.Equal(t, "mysql", config.DBDriver)
	require.Equal(t, "user:password@tcp(localhost:3306)/database", config.DBSource)
	require.Equal(t, "your-secret-key", config.TokenSymetricKey)
	require.Equal(t, time.Hour, config.AccessTokenDuration)
	require.Equal(t, 24*time.Hour, config.RefreshTokenDuration)

}
