package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Pwd 获取当前目录
func Pwd() string {
	dir, _ := os.Getwd()
	return dir
}

// IsExist 判断路径是否存在
// path 绝对路径
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDir 判断路径是否是一个文件夹
func IsDir(path string) bool {
	dir, err := os.Stat(path)
	if err != nil {
		return false
	}
	return dir.IsDir()
}

// IsFile 判断路径是否是一个文件
func IsFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

// ReadPath 获取目录列表
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

// ReadFile 读取文件1,返回的是byte数组,使用range遍历
func ReadFile(filename string) ([]byte, error) {
	// ioutil.ReadAll() 这个函数需要 os.Open()传一个io.Reader
	// ioutil.ReadFile() 只需要path string
	return ioutil.ReadFile(filename)
}

// ReadFileByBlock 读取文件2 分块读取
// 如果文件过大可以考虑这种方法  可在速度和内存占用之间取得很好的平衡。
// filename: 文件路径地址 path string
// bufSize: 读取的块大小
// hookFn: 读取到对应数据后执行的函数
func ReadFileByBlock(filename string, bufSize int, hookFn func([]byte)) {
	// 使用os.Open 获得文件实例
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	buf := make([]byte, bufSize) // buf用来接受文件的数据
	bfRd := bufio.NewReader(f)   // 创建reader
	for {
		// reader根据buf的大小bufSize读取到对应大小的数据
		// n为读取的字节数，实际可能小于len(buf),此时文件可能已经读完了,下次再读取就会报EOF错误
		n, err := bfRd.Read(buf)
		if n <= 0 { // 此时文件已经读取完了，err 会报EOF错误，提前判断截断掉panic，直接return
			fmt.Println()
			return
		}
		if err != nil {
			panic(err)
		}
		hookFn(buf[:n])
	}
}

// ReadFileByLine 逐行读取文件
// 这个可能更加常用于业务逻辑 根据go语言bufio类的备注，使用scanner可能更好
func ReadFileByLine(filename string, hookFn func([]byte)) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		hookFn(scan.Bytes())
	}
}
