package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsDir(path string) bool {
	dir, err := os.Stat(path)
	if err != nil {
		return false
	}
	return dir.IsDir()
}

func IsFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

func ReadPath(path string) ([]string, error) {
	dir, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !dir.IsDir() {
		return nil, fmt.Errorf("this path is a file, not dir")
	}

	info, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var contents []string
	for _, v := range info {
		if v.IsDir() {
			contents = append(contents, v.Name()+"/")
		} else {
			contents = append(contents, v.Name())
		}
	}

	return contents, nil
}

//func GetFileContent(path string) ([]byte, error) {
//	myFile, err := os.Open(path)
//	if err != nil {
//		return nil, err
//	}
//
//}
