package usecase

import (
	logger "github.com/sirupsen/logrus"
	"strings"

	"github.com/spf13/viper"
)

func NewConfig(filePath string, log *logger.Logger) (*Config, error) {
	viper.SetConfigFile(filePath)
	if filePath == "" {
		viper.SetConfigFile("./config.yml")
	}
	viper.AutomaticEnv()
	viper.SetEnvPrefix("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("unable to decode into struct, %v\n", err)
		return nil, err
	}

	viper.BindEnv("DB_URL", "DB_URL")

	var C Config

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Errorf("unable to decode into struct, %v\n", err)
		return nil, err
	}

	return &C, nil
}

type Config struct {
	LimitView     int //limit on the number of simultaneous connections from a client to view the file list
	LimitDownload int //limit on the number of simultaneous connections from a client to download or upload files
	Port          int
	DB_URL        string
}
