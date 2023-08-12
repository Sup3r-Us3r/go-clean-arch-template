package config

import (
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type Conf struct {
	Environment                   string `mapstructure:"ENVIRONMENT"`
	ServerPort                    string `mapstructure:"SERVER_PORT"`
	LogOutput                     string `mapstructure:"LOG_OUTPUT"`
	MongoURI                      string `mapstructure:"MONGO_URI"`
	MongoDBName                   string `mapstructure:"MONGO_DB_NAME"`
	JWTSecretKey                  string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationInMinutes        int    `mapstructure:"JWT_EXPIRATION_IN_MINUTES"`
	RefreshTokenExpirationInHours int    `mapstructure:"REFRESH_TOKEN_EXPIRATION_IN_HOURS"`
}

func LoadConfig(path string) *Conf {
	var cfg *Conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(filepath.Join(path, ".env"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func GetEnvAsString(key string) string {
	return viper.GetString(key)
}

func GetEnvAsStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func GetEnvAsStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}

func GetEnvAsStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func GetEnvAsBool(key string) bool {
	return viper.GetBool(key)
}

func GetEnvAsDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func GetEnvAsFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetEnvAsInt(key string) int {
	return viper.GetInt(key)
}

func GetEnvAsInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetEnvAsInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetEnvAsIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func GetEnvAsSizeInBytes(key string) uint {
	return viper.GetSizeInBytes(key)
}

func GetEnvAsStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func GetEnvAsTime(key string) time.Time {
	return viper.GetTime(key)
}

func GetEnvAsUint(key string) uint {
	return viper.GetUint(key)
}

func GetEnvAsUint16(key string) uint16 {
	return viper.GetUint16(key)
}

func GetEnvAsUint32(key string) uint32 {
	return viper.GetUint32(key)
}

func GetEnvAsUint64(key string) uint64 {
	return viper.GetUint64(key)
}
