package storage

const (
	getAllPhotosInfoTemplate = `SELECT id, name, created, edited FROM Photos;`

	getPhotosListTemplate = `SELECT id, name, created, edited FROM Photos WHERE id IN ($1);`

	getPhotoInfoTemplate = `SELECT id, name, created, edited FROM Photos WHERE id = $1;`

	uploadPhotoTemplate = `INSERT INTO Photos(name, created, edited, file_path) VALUES ($1, $2, $2, $3) RETURNING id;`

	uploadPhotosListTemplate = `INSERT INTO Photos(name, create, file_path) VALUES $1 ;`

	updatePhotoTemplate = `UPDATE Photos SET name = $1, edite = $2, file_path = $3 WHERE id = $4 RETURNING id;`

	getPhotosPathsTemplate = `SELECT file_path FROM Photos WHERE id IN ($1);`

	getPhotoPathTemplate = `SELECT file_path, name FROM Photos WHERE id = $1`

	getAllPhotosPathsTemplate = `SELECT file_path FROM Photos;`
)
