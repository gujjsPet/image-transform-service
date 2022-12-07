package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		return errors.New("filename " + fileName + " contains illegel character")
	}

	if !fileExists(fileName) {
		return errors.New("file " + fileName + " does not exist")
	}

	return nil
}

func GetFilePath(fileName string) (string, error) {
	fmt.Println(filepath.Abs(filepath.Join(FilesDirectory, fileName)))
	if err := validateExistingFile(fileName); err != nil {
		return "", err
	}

	return filepath.Join(FilesDirectory, fileName), nil
}
