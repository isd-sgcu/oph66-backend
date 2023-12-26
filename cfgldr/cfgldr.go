package cfgldr

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseConfig DatabaseConfig
	AppConfig      AppConfig
	RedisConfig    RedisConfig
}

type DatabaseConfig struct {
	Url string `mapstructure:"URL"`
}

type AppConfig struct {
	Port string `mapstructure:"PORT"`
	Env  string `mapstructure:"ENV"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"ADDR"`
	Port     string `mapstructure:"PORT"`
	Password string `mapstructure:"PASSWORD"`
}

func LoadConfig() (*Config, error) {
	dbCfgLdr := viper.New()
	dbCfgLdr.SetEnvPrefix("DB")
	dbCfgLdr.AutomaticEnv()
	dbCfgLdr.AllowEmptyEnv(false)
	dbConfig := DatabaseConfig{}
	if err := dbCfgLdr.Unmarshal(&dbConfig); err != nil {
		return nil, err
	}

	appCfgLdr := viper.New()
	appCfgLdr.SetEnvPrefix("APP")
	appCfgLdr.AutomaticEnv()
	dbCfgLdr.AllowEmptyEnv(false)
	appConfig := AppConfig{}
	if err := appCfgLdr.Unmarshal(&appConfig); err != nil {
		return nil, err
	}

	redisCfgLdr := viper.New()
	redisCfgLdr.SetEnvPrefix("REDIS")
	redisCfgLdr.AutomaticEnv()
	dbCfgLdr.AllowEmptyEnv(false)
	redisConfig := RedisConfig{}
	if err := redisCfgLdr.Unmarshal(&redisConfig); err != nil {
		return nil, err
	}

	return &Config{
		DatabaseConfig: dbConfig,
		AppConfig:      appConfig,
		RedisConfig:    redisConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
