package storage

import (
	"context"
	"file_storage/internal/types"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	logger "github.com/sirupsen/logrus"
	"time"
)

type Storage struct {
	//connect *pgx.Conn
	pool   *pgxpool.Pool
	logger *logger.Entry
}

func (s *Storage) Migrations(ctx context.Context) error {
	connection, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.WithError(err).Error("Error whole acquiring connection from the database pool!")
		return err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		s.logger.WithError(err).Error("Could not ping the db!")
		return err
	}

	_, err = connection.Exec(ctx, createTablePhotosTemplate)
	if err != nil {
		s.logger.WithError(err).Error("Error writing photo info to table")
		return err
	}
	return nil
}

func New(ctx context.Context, dbUrl string, log *logger.Logger) (Storage, error) {
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return Storage{}, err
	}

	logEntry := log.WithField("package", "storage")

	r := Storage{pool: pool, logger: logEntry}
	err = r.Migrations(ctx)
	if err != nil {
		return Storage{}, err
	}
	return r, nil
}

// SavePhoto - save one photo to DB
func (s *Storage) SavePhoto(ctx context.Context, name, path string, created time.Time) error {
	connection, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Error whole acquiring connection from the database pool!")
		return err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Could not ping the db!", err)
		return err
	}

	_, err = connection.Exec(ctx, uploadPhotoTemplate, name, created, path)
	if err != nil {
		s.logger.WithError(err).Errorln()
		return err
	}
	return nil
}

//func (r *Storage) SaveBatchPhotos(ctx context.Context, args [][]interface{}) error {
//	copyCount, err := r.connect.CopyFrom(
//		ctx,
//		pgx.Identifier{"Photos"},
//		[]string{"name", "create", "edite", "file_path"},
//		pgx.CopyFromRows(args),
//	)
//	if err != nil {
//		return err
//	}
//
//}

// GetAllPhotosInfo - get photo`s info (such as id, name, created date and edited date) from table photo for all photos
func (s *Storage) GetAllPhotosInfo(ctx context.Context) ([]*types.MetaData, error) {
	connection, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Error whole acquiring connection from the database pool!")
		return nil, err
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Could not ping the db!")
		return nil, err
	}

	rows, err := connection.Query(ctx, getAllPhotosInfoTemplate)
	if err != nil {
		s.logger.WithError(err).Errorln("Failed to get metadata from table.")
		return nil, err
	}
	defer rows.Close()

	metadata := []*types.MetaData{}
	for rows.Next() {
		data := &types.MetaData{}
		err = rows.Scan(&data.Id, &data.Name, &data.Created, &data.Edited)
		if err != nil {
			s.logger.WithError(err).Errorln("Failed to parse metadata to result")
			return nil, err
		}
		metadata = append(metadata, data)
	}
	return metadata, nil
}

// GetPhotoInfo - get photo`s info (such as id, name, created date and edited date) from table photo for one photo by id
func (s *Storage) GetPhotoInfo(ctx context.Context, id uint64) (*types.MetaData, error) {
	connection, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Error whole acquiring connection from the database pool!")
		return nil, err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Could not ping the db!")
		return nil, err
	}
	row := connection.QueryRow(ctx, getPhotoInfoTemplate, id)
	res := &types.MetaData{}
	err = row.Scan(&res.Id, &res.Name, &res.Created, &res.Edited)
	if err != nil {
		s.logger.WithError(err).Errorln("Failed to parse metadata to result")
		return nil, err
	}
	return res, nil
}

// GetPhotosListInfo - get photo`s info (such as id, name, created date and edited date) from table photo for some photo by ids
func (s *Storage) GetPhotosListInfo(ctx context.Context, ids []uint64) ([]*types.MetaData, error) {
	connection, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Error whole acquiring connection from the database pool!")
		return nil, err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Could not ping the db!")
		return nil, err
	}
	rows, err := connection.Query(ctx, getPhotosListTemplate, ids)
	if err != nil {
		s.logger.WithError(err).Errorln("Failed to get metadata from table")
		return nil, err
	}
	res := make([]*types.MetaData, len(ids))
	for rows.Next() {
		data := &types.MetaData{}
		err = rows.Scan(&data.Id, &data.Name, &data.Created, &data.Edited)
		if err != nil {
			s.logger.WithError(err).Errorln("Failed to parse metadata to result")
			return nil, err
		}
		res = append(res, data)
	}
	return res, nil
}

// GetPhotoPath - get file path and name from table for one photo by id
func (s *Storage) GetPhotoPath(ctx context.Context, id uint64) (path string, name string, err error) {
	connection, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Error whole acquiring connection from the database pool!")
		return "", "", err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Could not ping the db!")
		return "", "", err
	}

	row := connection.QueryRow(ctx, getPhotoPathTemplate, id)
	err = row.Scan(&path, &name)
	if err != nil {
		s.logger.WithError(err).Errorln("Failed to get file path from table")
		return "", "", err
	}
	return path, name, nil
}

// GetBatchPhotosPaths - get file path from table for some photos by ids
func (s *Storage) GetBatchPhotosPaths(ctx context.Context, ids []uint64) ([]string, error) {
	connection, err := s.pool.Acquire(context.Background())
	if err != nil {
		s.logger.WithError(err).Error("Error whole acquiring connection from the database pool!")
		return nil, err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		s.logger.WithError(err).Errorln("Could not ping the db!")
		return nil, err
	}
	rows, err := connection.Query(ctx, getPhotosPathsTemplate, ids)
	res := make([]string, len(ids))
	for rows.Next() {
		var path string
		err = rows.Scan(&path)
		if err != nil {
			s.logger.WithError(err).Errorln("Failed to get file path from table")
			return nil, err
		}
		res = append(res, path)
	}
	return res, nil
}
