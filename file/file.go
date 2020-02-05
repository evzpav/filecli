package file

import (
	"bufio"
	"fmt"
	"os"
)

type File struct{}

func New() (*File, error) {
	return &File{}, nil
}

func (f *File) ReadFile(path string) (string, error) {

	inFile, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("error to read file: %s", err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	var content string
	for scanner.Scan() {
		content += scanner.Text()
	}

	return content, nil
}

func (f *File) WriteFile(fileName, text string) (string, error) {

	if text == "" {
		return "", fmt.Errorf("Empty content")
	}
	filePointer, err := os.Create(fileName)
	if err != nil {
		return "", nil
	}

	_, err = filePointer.Write([]byte(text))
	if err != nil {
		return "", nil
	}

	filePointer.Close()
	return text, nil
}
