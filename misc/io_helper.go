package misc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func TryClose(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

func IsFileExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err == nil && !fileInfo.IsDir() {
		return true
	}
	return false
}

func IsDirExists(dirPath string) bool {
	fileInfo, err := os.Stat(dirPath)
	if err == nil && fileInfo.IsDir() {
		return true
	}
	return false
}

func WriteFile(filePath string, content []byte) error {
	directoryPath, _ := path.Split(filePath)
	if directoryPath != "" && !IsDirExists(directoryPath) {
		err := os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm)
		if err != nil {
			return fmt.Errorf("os.MkdirAll(directoryPath, os.ModeDir|os.ModePerm): %v", err)
		}
	}
	err := ioutil.WriteFile(filePath, content, 0666)
	if err != nil {
		return fmt.Errorf("ioutil.WriteFile(filePath, content, 0666): %v", err)
	}
	return nil
}

func CopyFile(dst string, src string) error {
	srcReader, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("os.Open(src): %v", err)
	}
	defer TryClose(srcReader)
	dstWriter, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666): %v", err)
	}
	defer TryClose(dstWriter)
	_, err = io.Copy(dstWriter, srcReader)
	if err != nil {
		return fmt.Errorf("io.Copy(dstWriter, srcReader): %v", err)
	}
	return nil
}

func WriteTextFile(filePath string, text string) error {
	err := WriteFile(filePath, []byte(text))
	if err != nil {
		return fmt.Errorf("WriteFile(filePath, []byte(text): %v", err)
	}
	return nil
}

func WriteJsonFile(filePath string, v interface{}) error {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf(`json.MarshalIndent(v, "", "  "): %v`, err)
	}
	err = WriteFile(filePath, jsonData)
	if err != nil {
		return fmt.Errorf("WriteFile(filePath, jsonData): %v", err)
	}
	return nil
}

func WriteBase64File(filePath string, base64String string) error {
	fileData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString(base64String): %v", err)
	}
	return WriteFile(filePath, fileData)
}
