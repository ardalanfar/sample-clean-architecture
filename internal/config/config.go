package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type CnfPostgreSQL struct {
	Username string
	Password string
	Dbname   string
	Host     string
	Port     string
}

type SendEmail struct {
	Username string
	Password string
	SmtpHost string
	SmtpPort string
	From     string
}

type Logger struct {
	Level        string
	LogDirectory string
}

type Config struct {
	DB    *CnfPostgreSQL
	Email *SendEmail
	Log   *Logger
}

func InitConfig() (*Config, error) {
	cfg := &Config{}
	env := os.Getenv("APP_ENV")
	Configpath := os.Getenv("CONFIGPATH_ENV")

	if env == "" {
		env = "development"
	}

	if Configpath == "" {
		RootPath, err := getConfigRootPath()
		if err != nil {
			return nil, err
		}
		Configpath = RootPath
	}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(Configpath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	CfgPath := viper.Get(Configpath)
	if err := viper.Unmarshal(&CfgPath); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil
}

func getConfigRootPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	absCurrentDir, err := filepath.Abs(wd)
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(absCurrentDir, "internal/config/")

	return configPath, nil
}
