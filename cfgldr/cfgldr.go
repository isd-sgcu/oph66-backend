package cfgldr

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseConfig DatabaseConfig
	AppConfig      AppConfig
	RedisConfig    RedisConfig
	OAuth2Config   OAuth2Config
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

type OAuth2Config struct {
	RedirectURL  string   `mapstructure:"REDIRECT_URL"`
	ClientID     string   `mapstructure:"CLIENT_ID"`
	ClientSecret string   `mapstructure:"CLIENT_SECRET"`
	Scopes       []string `mapstructure:"SCOPES"`
	Endpoint     string   `mapstructure:"ENDPOINT"`
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

	oauth2CfgLdr := viper.New()
	oauth2CfgLdr.SetEnvPrefix("OAUTH2")
	oauth2CfgLdr.AutomaticEnv()
	oauth2CfgLdr.AllowEmptyEnv(false)
	oauth2Config := OAuth2Config{}
	if err := oauth2CfgLdr.Unmarshal(&oauth2Config); err != nil {
		return nil, err
	}

	return &Config{
		DatabaseConfig: dbConfig,
		AppConfig:      appConfig,
		RedisConfig:    redisConfig,
		OAuth2Config:   oauth2Config,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
