package _1basic

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		filename string
	}{
		{filename: "./test.txt"},
	}

	for _, v := range tests {
		result, err := ReadFile(v.filename)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("读取文件：%s 成功！内容如下:\n", v.filename)
		for _, content := range result {
			fmt.Println(string(content))
		}

	}
}

// go14 test -v -run TestReadFileByBlock file_test.go file.go
func TestReadFileByBlock(t *testing.T) {
	tests := []struct {
		filename string
		bufSize  int
		hookFn   func([]byte)
	}{
		{filename: "./test.txt",
			bufSize: 1,
			hookFn: func(bytes []byte) {
				for _, v := range bytes {
					t.Log(string(v))
				}
			}},
	}

	for _, v := range tests {
		ReadFileByBlock(
			v.filename,
			v.bufSize,
			v.hookFn)
	}
}

// go14 test -v -run TestReadFileByLine file_test.go file.go
func TestReadFileByLine(t *testing.T) {
	tests := []struct {
		filename string
		hookFn   func([]byte)
	}{
		{filename: "./test.txt",
			hookFn: func(bytes []byte) {
				for _, v := range bytes {
					fmt.Printf(string(v))
				}
				fmt.Println()
			}},
	}

	for _, v := range tests {
		ReadFileByLine(
			v.filename,
			v.hookFn)
	}
}
