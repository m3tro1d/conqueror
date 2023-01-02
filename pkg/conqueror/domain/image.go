package domain

import (
	stderrors "errors"
	"io"
)

var (
	ErrImageNotFound = stderrors.New("image not found")
)

func NewImage(id ImageID, path string) (*Image, error) {
	return &Image{
		id:   id,
		path: path,
	}, nil
}

type Image struct {
	id   ImageID
	path string
}

type ImageRepository interface {
	NextID() ImageID
	Store(image *Image, file io.Reader) error
	GetByID(id ImageID) (*Image, error)
}

func (i *Image) ID() ImageID {
	return i.id
}

func (i *Image) Path() string {
	return i.path
}
