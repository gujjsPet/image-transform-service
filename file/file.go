package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const FilesDirectory string = "storage/"
const IllegalCharacters string = `/\/?<>:*|"`

func validFilename(fileName string) bool {
	return !strings.ContainsAny(fileName, IllegalCharacters) && !strings.Contains(fileName, "..")
}

func fileExists(fileName string) bool {
	_, err := os.Stat(filepath.Join(FilesDirectory, fileName))
	return err == nil
}

func validateExistingFile(fileName string) error {
	if !validFilename(fileName) {
		return errors.New("filename " + fileName + " contains illegal character")
	}

	if !fileExists(fileName) {
		return errors.New("file " + fileName + " does not exist")
	}

	return nil
}

func constructFilePath(fileName string) string {
	return filepath.Join(FilesDirectory, fileName)
}

func GetFilePath(fileName string) (string, error) {
	if err := validateExistingFile(fileName); err != nil {
		return "", err
	}

	return constructFilePath(fileName), nil
}

func GenerateFilename(n string) string {
	nSlices := strings.Split(n, ".")
	fExt := nSlices[len(nSlices)-1]
	return uuid.New().String() + "." + fExt
}
