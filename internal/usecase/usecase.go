package usecase

import (
	"context"
	"file_storage/internal/repository/storage"
	"file_storage/internal/types"
	logger "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var (
	Directory = filepath.Dir("/downloadedFiles/")
)

type UseCase struct {
	storage storage.Storage
	log     *logger.Logger
}

func New(storage storage.Storage, log *logger.Logger) (*UseCase, error) {
	return &UseCase{
		storage: storage,
		log:     log,
	}, nil
}

func (s *UseCase) GetPhoto(ctx context.Context, id uint64) ([]byte, error) {
	s.log.Info("Star get photo")
	path, err := s.storage.GetPhotoPath(ctx, id)
	if err != nil {
		s.log.Errorf("Error get photo path: %v", err)
		return nil, err
	}
	photo, err := os.ReadFile(path)
	if err != nil {
		s.log.Errorf("Error get photo: %v", err)
		return nil, err
	}
	s.log.Info("Get photo")
	return photo, nil
}

func (s *UseCase) PostPhoto(ctx context.Context, name string, photo []byte) error {
	s.log.Info("Start save photo")
	path := filepath.Join(Directory, name)
	now := time.Now()
	err := os.WriteFile(path, photo, 0644)
	if err != nil {
		s.log.Errorf("Error save file: %v", err)
		return err
	}

	err = s.storage.SavePhoto(ctx, name, path, now)
	if err != nil {
		s.log.Errorf("Error create file: %v", err)
		removeErr := os.Remove(path)
		if removeErr != nil {
			s.log.WithError(err).Errorln("Error delete file")
			return removeErr
		}
		return err
	}

	s.log.Tracef("Save file: %s", name)
	return nil
}

func (s *UseCase) GetAllPhotosInfo(ctx context.Context) ([]*types.MetaData, error) {
	s.log.Info("Start get all photos info")
	photosInfo, err := s.storage.GetAllPhotosInfo(ctx)
	if err != nil {
		s.log.Errorf("Error get all photos info from DB: %v", err)
		return nil, err
	}

	s.log.Info("Get all photos info")
	return photosInfo, nil
}
