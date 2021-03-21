package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// 最简单的服务器， 返回字符串hello world
	simpleServe()

	// 附带装饰器
	//errorHandlerServe()
}

// 最简单的服务器
func simpleServe() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	err := http.ListenAndServe(":18999", nil)
	if err != nil {
		panic(err)
	}
}

// 装饰器
func errorHandlerServe() {
	// HandlerFunc的第二个参数需要一个函数func(ResponseWriter, *Request)
	// 我们传递的函数是HandlerFileList
	//
	// 我们定义一个appHandler 这个appHandler是一个函数func(ResponseWriter, *Request)的别名
	// 增加一个对应的函数 errWrapper,该函数会返回匿名函数func(ResponseWriter, *Request)
	//
	// 代码的执行流程是:  HandlerFileList -> errWrapper -> http.HandleFunc
	http.HandleFunc("/list/", errWrapper(HandlerFileList))

	err := http.ListenAndServe(":18999", nil)
	if err != nil {
		panic(err)
	}
}

// 自定义error
type userError interface {
	error
	Message() string
}

// 新定义的handlerFunc，多返回了一个error
type appHandler func(write http.ResponseWriter, request *http.Request) error

// 对handlerFunc抛出的数据进行处理
// 业务逻辑在参数handler里面
func errWrapper(handler appHandler) func(w http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, request *http.Request) {

		// 进行panic处理, recover截取panic写入日志，并返回http code
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(w,
					"msg:"+http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(w, request)

		// 对handler函数进行整体的错误处理
		if err != nil {
			if userErr, ok := err.(userError); ok {
				http.Error(w, "msg:"+userErr.Message(), http.StatusBadRequest)
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, "msg:"+http.StatusText(code), code)
		}
	}
}

func HandlerFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}

	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
