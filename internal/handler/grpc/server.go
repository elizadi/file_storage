package grpc

import (
	"context"
	"errors"
	"file_storage/internal/types"
	"file_storage/internal/usecase"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync"
)

type server struct {
	UnimplementedFileStorageServer
	uc      *usecase.UseCase
	mu      sync.Mutex
	clients map[string]*client
	cfg     *usecase.Config
	log     *logger.Entry
}

type client struct {
	Addr            string
	limiterDownload usecase.Limiter
	limiterView     usecase.Limiter
}

func NewServer(uc *usecase.UseCase, cfg *usecase.Config, log *logger.Logger) (*grpc.Server, error) {
	logg := log.WithField("package", "grpc")
	s := grpc.NewServer()

	RegisterFileStorageServer(s, &server{
		mu:      sync.Mutex{},
		clients: make(map[string]*client),
		cfg:     cfg,
		log:     logg,
		uc:      uc,
	})

	return s, nil
}

func (s *server) GetPhoto(ctx context.Context, in *GetPhotoRequest) (*GetPhotoResponse, error) {
	limiter, err := s.limiterDownload(ctx)
	if err != nil {
		s.log.WithError(err).Errorln("Limited connection")
		return nil, err
	}
	err = limiter.Acquire()
	if err != nil {
		return nil, err
	}
	defer limiter.Release()

	photo, name, err := s.uc.GetPhoto(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &GetPhotoResponse{
		Photo: photo,
		Name:  name,
	}, nil
}

func (s *server) PostPhoto(ctx context.Context, in *PostPhotoRequest) (*Empty, error) {
	limiter, err := s.limiterDownload(ctx)
	if err != nil {
		s.log.WithError(err).Errorln("Limited connection")
		return nil, err
	}
	err = limiter.Acquire()
	if err != nil {
		return nil, err
	}
	defer limiter.Release()

	err = s.uc.PostPhoto(ctx, in.Name, in.Photo)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *server) GetAllPhotosInfo(ctx context.Context, _ *Empty) (*GetPhotosInfoResponse, error) {
	limiter, err := s.limiterView(ctx)
	if err != nil {
		s.log.WithError(err).Errorln("Limited connection")
		return nil, err
	}
	err = limiter.Acquire()
	if err != nil {
		return nil, err
	}
	defer limiter.Release()

	photos, err := s.uc.GetAllPhotosInfo(ctx)
	if err != nil {
		return nil, err
	}

	return &GetPhotosInfoResponse{FileInfo: toGrpcMetadata(photos)}, nil
}

func toGrpcMetadata(metadata []*types.MetaData) []*MetaData {
	res := make([]*MetaData, 0, len(metadata))
	for _, m := range metadata {
		res = append(res, &MetaData{
			Id:        m.Id,
			Name:      m.Name,
			Created:   timestamppb.New(m.Created),
			Edited:    timestamppb.New(m.Edited),
			Extension: m.Extension,
		})
	}

	return res
}

func (s *server) limiterDownload(ctx context.Context) (*usecase.Limiter, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("no peer found")
	}
	s.mu.Lock()
	cl, ok := s.clients[p.Addr.String()]
	if !ok {
		cl = &client{
			Addr:            p.Addr.String(),
			limiterDownload: usecase.NewLimiter(s.cfg.LimitDownload),
			limiterView:     usecase.NewLimiter(s.cfg.LimitView),
		}
		s.clients[p.Addr.String()] = cl
	}
	s.mu.Unlock()

	return &cl.limiterDownload, nil
}

func (s *server) limiterView(ctx context.Context) (*usecase.Limiter, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("no peer found")
	}
	s.mu.Lock()
	cl, ok := s.clients[p.Addr.String()]
	if !ok {
		cl = &client{
			Addr:            p.Addr.String(),
			limiterDownload: usecase.NewLimiter(s.cfg.LimitDownload),
			limiterView:     usecase.NewLimiter(s.cfg.LimitView),
		}
		s.clients[p.Addr.String()] = cl
	}
	s.mu.Unlock()

	return &cl.limiterView, nil
}
