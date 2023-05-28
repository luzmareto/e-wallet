package utils

import (
	"mime/multipart"

	"github.com/gabriel-vasile/mimetype"
)

func GetMIMEType(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buf := make([]byte, 512)
	if _, err := src.Read(buf); err != nil {
		return "", err
	}

	// Determine the MIME type of the file
	mime := mimetype.Detect(buf)
	return mime.String(), nil
}

func IsAllowedFileType(fileType string) bool {
	// Define the list of allowed file types
	allowedTypes := []string{
		"image/jpg",
		"image/jpeg",
		"image/png",
	}

	// Check if the file type is in the allowed list
	for _, allowedType := range allowedTypes {
		if fileType == allowedType {
			return true
		}
	}

	return false
}
