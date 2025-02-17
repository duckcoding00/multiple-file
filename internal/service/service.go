package service

import "mime/multipart"

type Service struct {
	File interface {
		Upload(files []*multipart.FileHeader) error
	}
}

func NewService() Service {
	return Service{
		File: &FileService{},
	}
}
