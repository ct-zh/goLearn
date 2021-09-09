package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/yanyiwu/gojieba"
)

var (
	cpuProfile = flag.String("cpu", "", "性能测试")
	str        = flag.String("s", "", "计算分词的字符串")
)

func main() {
	flag.Parse()

	if len(*str) == 0 {
		os.Exit(0)
	}

	if len(*cpuProfile) > 0 {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			panic(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	x := gojieba.NewJieba()
	defer x.Free()

	words := x.Cut(*str, true)

	data, err := json.Marshal(words)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
