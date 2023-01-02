package mysql

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/domain"
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
)

func NewImageRepository(client ClientContext, filesDir string) domain.ImageRepository {
	return &imageRepository{
		client:   client,
		filesDir: filesDir,
	}
}

type imageRepository struct {
	client   ClientContext
	filesDir string
}

func (repo *imageRepository) NextID() domain.ImageID {
	return domain.ImageID(uuid.Generate())
}

func (repo *imageRepository) Store(image *domain.Image, file io.Reader) error {
	const sqlQuery = `INSERT INTO image (id, path)
		              VALUES (?, ?)
		              ON DUPLICATE KEY UPDATE path=VALUES(path)`

	args := []interface{}{
		binaryUUID(image.ID()),
		image.Path(),
	}

	_, err := repo.client.ExecContext(context.Background(), sqlQuery, args...)
	if err != nil {
		return errors.WithStack(err)
	}

	return repo.storeFile(image, file)
}

func (repo *imageRepository) GetByID(id domain.ImageID) (*domain.Image, error) {
	const sqlQuery = `SELECT id, path
		              FROM image
		              WHERE id = ?
		              LIMIT 1`

	var image sqlxImage
	err := repo.client.GetContext(context.Background(), &image, sqlQuery, binaryUUID(id))
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(domain.ErrImageNotFound)
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return domain.NewImage(
		domain.ImageID(image.ID),
		image.Path,
	)
}

func (repo *imageRepository) storeFile(image *domain.Image, file io.Reader) error {
	filePath := path.Join(repo.filesDir, image.Path())
	dstFile, err := os.Create(filePath)
	if err != nil {
		return errors.WithStack(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer dstFile.Close()

	_, err = io.Copy(dstFile, file)
	return errors.WithStack(err)
}
