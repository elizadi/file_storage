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

//func (s *server) PostPhotos(ctx context.Context, in *kitty.PostPhotosRequest) (*kitty.Empty, error) {
//	paths := []string{}
//	argsForDB := [][]interface{}{}
//	for _, photo := range in.Photos {
//		path := fmt.Sprintf("%s/%s%s", Directory, photo.MetaData.Name, photo.MetaData.Extension)
//		err := os.WriteFile(path, photo.Photo, 0644)
//		if err != nil {
//			log.Printf("Error save file %s%s: %v", photo.MetaData.Name, photo.MetaData.Extension, err)
//			return nil, err
//		}
//		paths = append(paths, path)
//		arg := []interface{}{photo.MetaData.Name, photo.MetaData.Created, photo.MetaData.Edited, path}
//		argsForDB = append(argsForDB, arg)
//	}
//
//	err = s.Storage.S
//}

//func (s *server) GetPhotosList(ctx context.Context, in *kitty.GetPhotoRequest) (*kitty.GetPhotoResponse, error) {
//
//}

func main() {
	log := logger.New()
	//log.SetFormatter(&logger.TextFormatter{})
	//log.SetLevel(logger.TraceLevel)
	conf, err := usecase.NewConfig("", log)
	//log.Traceln("Directory", usecase.Directory)
	if err != nil {
		log.WithError(err).Errorln("Error create config")
		os.Exit(1)
	}
	if conf.Port == 0 {
		conf.Port = 8000
	}
	//log.Traceln("url", conf.DB_URL)
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
		log.Errorf("failed to listen: %v", err)
		os.Exit(1)
	}
	ctx := context.Background()
	storage, err := storage.New(ctx, conf.DB_URL, log)
	if err != nil {
		log.Errorf("failed to connect db: %v", err)
		os.Exit(1)
	}

	if _, err := os.Stat(usecase.Directory); os.IsNotExist(err) {
		createDirErr := os.Mkdir(usecase.Directory, 0777)
		if createDirErr != nil {
			log.Errorf("failed to create dir: %v", createDirErr)
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
