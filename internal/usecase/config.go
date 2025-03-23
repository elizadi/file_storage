package usecase

import (
	"errors"
	logger "github.com/sirupsen/logrus"
	"strings"

	"github.com/spf13/viper"
)

var incorrectParameterError = errors.New("incorrect parameter")

func NewConfig(filePath string, log *logger.Logger) (*Config, error) {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("err loading: %v", err)
	//}
	//
	//dbUrl := os.Getenv("DB_URL")
	//if dbUrl == "" {
	//	log.WithError(incorrectParameterError).Errorln("Empty DB URL")
	//	return nil, incorrectParameterError
	//}
	//
	//limitView, err := strconv.Atoi(os.Getenv("LimitView"))
	//if err != nil {
	//	log.WithError(incorrectParameterError).Errorln("Cannot convert LimitView to int")
	//}
	//if limitView <= 0 {
	//	log.Info("Limit view is less than zero")
	//	limitView = 100
	//}
	//
	//limitDownload, err := strconv.Atoi(os.Getenv("LimitDownload"))
	//if err != nil {
	//	log.WithError(incorrectParameterError).Errorln("Cannot convert limitDownload to int")
	//}
	//if limitDownload <= 0 {
	//	log.Info("Limit download is less than zero")
	//	limitDownload = 10
	//}
	//
	//port, err := strconv.Atoi(os.Getenv("Port"))
	//if err != nil {
	//	log.WithError(incorrectParameterError).Errorln("Cannot convert port to int")
	//}
	//if port <= 0 {
	//	log.Info("Limit view is incorrect")
	//	limitDownload = 8000
	//}

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

	//C = Config{
	//	LimitView:     limitView,
	//	LimitDownload: limitDownload,
	//	Port:          port,
	//	DB_URL:        dbUrl,
	//}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Errorf("unable to decode into struct, %v\n", err)
		return nil, err
	}

	return &C, nil
}

type Config struct {
	LimitView     int
	LimitDownload int
	Port          int
	DB_URL        string
}
