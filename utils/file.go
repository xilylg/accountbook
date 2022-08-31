package utils

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

func CreateFile(c *gin.Context, filename string, pri string) error {
	if filename == "" {
		return errors.New("filename empty")
	}
	_, err := os.Stat(filename)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func WriteFile(c *gin.Context, filename string, data string) error {
	if filename == "" || data == "" {
		return errors.New("filename or data is empty")
	}
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.WriteString(data)
	return err
}

func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return file
}

func CurrentRootPath() string {
	fl := CurrentFile()
	return filepath.Dir(filepath.Dir(fl))
}
