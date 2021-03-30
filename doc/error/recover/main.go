package main

import "fmt"

func tryRecover() {
	// 内建函数recover允许程序管理恐慌过程中的Go程。
	// 在defer的函数中，执行recover调用会取回传至panic调用的错误值，恢复正常执行，停止恐慌过程。
	// 若recover在defer的函数之外被调用，它将不会停止恐慌过程序列。
	// 在此情况下，或当该Go程不在恐慌过程中时，或提供给panic的实参为nil时，recover就会返回nil。
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("Nothing to recover")
			return
		}

		if err, ok := r.(error); ok {
			fmt.Println("Error: ", err)
		} else {
			panic(fmt.Sprintf("Panic: %v", r))
		}
	}()

	//panic(123)
}

func main() {
	tryRecover()
}
