package main

import (
	"context"
	"file_storage/internal/handler/grpc"
	"file_storage/internal/repository/storage"
	"file_storage/internal/usecase"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net"
	"os"
)

func main() {
	log := logger.New()

	conf, err := usecase.NewConfig("", log)
	if err != nil {
		log.WithError(err).Errorln("Error create config")
		os.Exit(1)
	}
	if conf.Port == 0 {
		conf.Port = 8000
	}

	if conf.DB_URL == "" {
		log.Errorf("Error get DB URL")
		os.Exit(1)
	}
	if conf.LimitView <= 0 {
		conf.LimitView = 100
	}
	if conf.LimitDownload <= 0 {
		conf.LimitDownload = 10
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.WithError(err).Errorln("failed to listen")
		os.Exit(1)
	}
	ctx := context.Background()
	storage, err := storage.New(ctx, conf.DB_URL, log)
	if err != nil {
		log.WithError(err).Errorln("failed to connect db")
		os.Exit(1)
	}

	if _, err = os.Stat(usecase.Directory); os.IsNotExist(err) {
		createDirErr := os.Mkdir(usecase.Directory, 0777)
		if createDirErr != nil {
			log.WithError(createDirErr).Errorln("failed to create dir")
			os.Exit(1)
		}
	}

	uc, err := usecase.New(storage, log)
	if err != nil {
		log.WithError(err).Errorln("error creating usecase")
		os.Exit(1)
	}
	s, err := grpc.NewServer(uc, conf, log)
	if err != nil {
		log.WithError(err).Errorln("error creating server")
		os.Exit(1)
	}
	log.Infof("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.WithError(err).Errorln("failed to serve")
		os.Exit(1)
	}
}
