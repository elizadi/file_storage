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
	conf, err := usecase.NewConfig("", log)
	if err != nil {
		log.Errorf("Error create config: %v\n", err)
		return
	}
	if conf.Port == 0 {
		conf.Port = 8000
	}
	if conf.DB_URL == "" {
		log.Errorf("Error get DB URL")
		return
	}
	if conf.LimitView <= 0 {
		conf.LimitView = 100
	}
	if conf.LimitDownload <= 0 {
		conf.LimitDownload = 10
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	ctx := context.Background()
	storage, err := storage.New(ctx, conf.DB_URL, log)
	if err != nil {
		logger.Fatalf("failed to connect db: %v", err)
	}

	if _, err := os.Stat(usecase.Directory); os.IsNotExist(err) {
		createDirErr := os.Mkdir(usecase.Directory, 0777)
		if err != nil {
			logger.Fatalf("failed to create dir: %v", createDirErr)
		}
	}

	uc, err := usecase.New(storage, log)
	if err != nil {
		log.WithError(err).Errorln("error creating usecase")
		return
	}
	s, err := grpc.NewServer(uc, conf, log)
	if err != nil {
		log.WithError(err).Errorln("error creating server")
		return
	}
	log.Infof("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.WithError(err).Errorln("failed to serve")
		return
	}
}
