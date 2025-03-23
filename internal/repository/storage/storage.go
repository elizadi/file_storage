package storage

import (
	"context"
	"file_storage/internal/types"
	"fmt"
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

func (r *Storage) Migrations(ctx context.Context) error {
	connection, err := r.pool.Acquire(context.Background())
	if err != nil {
		fmt.Println("Error whole acquiring connection from the database pool!", err)
		return err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		fmt.Println("Could not ping the db!", err)
		return err
	}

	_, err = connection.Exec(ctx, createTablePhotosTemplate)
	if err != nil {
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
func (r *Storage) SavePhoto(ctx context.Context, name, path string, created time.Time) error {
	connection, err := r.pool.Acquire(context.Background())
	if err != nil {
		r.logger.WithError(err).Errorln("Error whole acquiring connection from the database pool!")
		return err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		r.logger.WithError(err).Errorln("Could not ping the db!", err)
		return err
	}

	_, err = connection.Exec(ctx, uploadPhotoTemplate, name, created, path)
	if err != nil {
		r.logger.WithError(err).Errorln()
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

func (r *Storage) GetAllPhotosInfo(ctx context.Context) ([]*types.MetaData, error) {
	connection, err := r.pool.Acquire(context.Background())
	if err != nil {
		fmt.Println("Error whole acquiring connection from the database pool!", err)
		return nil, err
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		fmt.Println("Could not ping the db!", err)
		return nil, err
	}

	rows, err := connection.Query(ctx, getAllPhotosInfoTemplates)
	if err != nil {
		r.logger.WithError(err).Errorln()
		return nil, err
	}
	defer rows.Close()

	metadata := []*types.MetaData{}
	for rows.Next() {
		data := &types.MetaData{}
		var err = rows.Scan(&data.Id, &data.Name, &data.Created, &data.Edited)
		if err != nil {
			return nil, err
		}
		metadata = append(metadata, data)
	}
	return metadata, nil
}

func (r *Storage) GetPhotoInfo(ctx context.Context, id uint64) (*types.MetaData, error) {
	connection, err := r.pool.Acquire(context.Background())
	if err != nil {
		fmt.Println("Error whole acquiring connection from the database pool!", err)
		return nil, err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		fmt.Println("Could not ping the db!", err)
		return nil, err
	}
	row := connection.QueryRow(ctx, getPhotosListTemplates, []uint64{id})
	res := &types.MetaData{}
	err = row.Scan(&res.Id, &res.Name, &res.Created, &res.Edited)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Storage) GetPhotosListInfo(ctx context.Context, ids []uint64) ([]*types.MetaData, error) {
	connection, err := r.pool.Acquire(context.Background())
	if err != nil {
		fmt.Println("Error whole acquiring connection from the database pool!", err)
		return nil, err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		fmt.Println("Could not ping the db!", err)
		return nil, err
	}
	rows, err := connection.Query(ctx, getPhotosListTemplates, ids)
	if err != nil {
		return nil, err
	}
	res := make([]*types.MetaData, len(ids))
	for rows.Next() {
		data := &types.MetaData{}
		err = rows.Scan(&data.Id, &data.Name, &data.Created, &data.Edited)
		if err != nil {
			return nil, err
		}
		res = append(res, data)
	}
	return res, nil
}

func (r *Storage) GetPhotoPath(ctx context.Context, id uint64) (string, error) {
	connection, err := r.pool.Acquire(context.Background())
	if err != nil {
		fmt.Println("Error whole acquiring connection from the database pool!", err)
		return "", err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		fmt.Println("Could not ping the db!", err)
		return "", err
	}

	row := connection.QueryRow(ctx, getPhotoPathTemplate, id)
	var res string
	err = row.Scan(&res)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (r *Storage) GetBatchPhotosPaths(ctx context.Context, ids []uint64) ([]string, error) {
	connection, err := r.pool.Acquire(context.Background())
	if err != nil {
		fmt.Println("Error whole acquiring connection from the database pool!", err)
		return nil, err
	}

	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		fmt.Println("Could not ping the db!", err)
		return nil, err
	}
	rows, err := connection.Query(ctx, getPhotosPathsTemplate, ids)
	res := make([]string, len(ids))
	for rows.Next() {
		var path string
		err = rows.Scan(&path)
		if err != nil {
			return nil, err
		}
		res = append(res, path)
	}
	return res, nil
}
