package fileInput

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type FileInput interface {
	GetContentOfFile(string) []string
}

type fileInput struct {
}

func NewFileInput() FileInput {
	return fileInput{}
}

func (f fileInput) GetContentOfFile(filepath string) []string {
	fileContent := readFile(filepath)
	return fileContent
}

func readFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var readFile []string
	for {
		line, err := reader.ReadString('\n')
		if io.EOF == err {
			break
		}
		if err != nil {
			log.Fatal("Error while reading file: ", err)
		}

		readFile = append(readFile, strings.TrimSpace(strings.TrimSuffix(line, "\n")))
	}

	return readFile
}
