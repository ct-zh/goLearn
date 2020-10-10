package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 获取当前目录
func Pwd() string {
	dir, _ := os.Getwd()
	return dir
}

// 判断路径是否存在
// path 绝对路径
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 判断路径是否是一个文件夹
func IsDir(path string) bool {
	dir, err := os.Stat(path)
	if err != nil {
		return false
	}
	return dir.IsDir()
}

// 判断路径是否是一个文件
func IsFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

// 获取目录列表
func ReadPath(path string) ([]string, error) {
	dir, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !dir.IsDir() {
		return nil, fmt.Errorf("this path is a 02file, not dir")
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
