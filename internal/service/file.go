package service

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FileService struct {
}

func (s *FileService) createFolder() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed tracking current directory :%w", err)
	}

	rootDir := filepath.Join(currentDir, "..")
	dataDir := filepath.Join(rootDir, "data")

	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed create folder :%w", err)
	}
	return dataDir, nil
}

func (s *FileService) validationFile(fileHeader *multipart.FileHeader, file multipart.File) error {
	// check size file, max 2mb can be accepted
	if fileHeader.Size > 2<<20 {
		return fmt.Errorf("file %s is too large (max 2MB)", fileHeader.Filename)
	}

	// validation ext
	allowedExtension := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	fileExtension := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExtension[fileExtension] {
		return fmt.Errorf("invalid ext file, only images ext can be accepted")
	}

	// validation mime
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Reset file pointer after reading
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("error resetting file pointer: %w", err)
	}

	mimeType := http.DetectContentType(buffer)
	if !strings.HasPrefix(mimeType, "image/") {
		return fmt.Errorf("invalid mime type, only image/ can be accepted")
	}

	return nil
}

func (s *FileService) saveFile(dir string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	if err := s.validationFile(fileHeader, file); err != nil {
		return err
	}

	dst, err := os.Create(filepath.Join(dir, fileHeader.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	return nil
}

func (s *FileService) Upload(files []*multipart.FileHeader) error {
	// create folder
	dir, err := s.createFolder()
	if err != nil {
		return err
	}
	log.Println(dir)
	for _, file := range files {
		if err := s.saveFile(dir, file); err != nil {
			return err
		}
	}

	return nil
}
