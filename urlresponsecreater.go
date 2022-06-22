package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// var URL_RESPONSE_PATH = os.Getenv("URL_RESPONSE_PATH")
var URL_RESPONSE_PATH = "/Users/purnachand/Desktop/examples"

func main() {
	urlAsString := os.Args[1]
	fmt.Println(urlAsString)
	theURL, _ := url.Parse(urlAsString)
	values := theURL.Query()
	sb := strings.Builder{}
	for k, v := range values {
		sb.WriteString(fmt.Sprintf("%s : \n", k))
		for _, value := range v {
			sb.WriteString(fmt.Sprintf("\t %s \n", value))
		}
	}
	theURL.RawQuery = "" //Removes all the queries from url
	folderPath, err := createFolder(createFolderName(theURL))
	if err != nil {
		panic(err)
	}
	if err := createQueryParamsFile(folderPath, "queryParams.txt", &sb); err != nil {
		panic(err)
	}
	if err := createEmptyFile(folderPath, "response.json"); err != nil {
		panic(err)
	}
	if err := createEmptyFile(folderPath,"Lambda.txt");err!=nil {
		panic(err)
	}
}

func createFolder(folderName string) (string, error) {
	path := fmt.Sprintf("%s/%s", URL_RESPONSE_PATH, folderName)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

func createFolderName(theURL *url.URL) string {
	theReplacer := strings.NewReplacer(".", "_", "/", "_")
	folderName := fmt.Sprintf("%s", theReplacer.Replace(theURL.String()[8:]))
	return folderName
}

func createEmptyFile(path string, fileName string) error {
	rFile, err := createFile(path, fileName)
	rFile.Close()
	return err
}

func createFile(path string, fileName string) (*os.File, error) {
	absFilePath := fmt.Sprintf("%s/%s", path, fileName)
	_, err := os.ReadFile(absFilePath)
	if err == nil {
		os.Truncate(absFilePath, 0) //Erase all contents in the file
	}
	return os.Create(absFilePath)
}

func createQueryParamsFile(path string, fileName string, content *strings.Builder) error {
	qFile, err := createFile(path, fileName)
	x := content.String()
	fmt.Println(x)
	fmt.Println(qFile.Name())
	qFile.Write(bytes.NewBufferString(x).Bytes())
	qFile.Close()
	return err
}
