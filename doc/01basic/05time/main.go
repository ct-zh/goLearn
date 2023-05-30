package main

import (
	"fmt"
	"time"
)

const myTimeFormat = "2006-01-02 15:04:05"

// time包的用法demo
func main() {
	//fmt.Println("=======time高频用法====")
	//timeUsual()
	//fmt.Println()

	fmt.Println("=======time demo====")
	timeDemo()
	fmt.Println()

	//fmt.Println("=======计时器Timer 和 ticker====")
	//timerTicker()
	//fmt.Println()

	//fmt.Println("=======time 其他types====")
	//weekday()
	//month()
	//location()
	//fmt.Println()
}

// time包高频用法
func timeUsual() {
	// 获取当前时间戳/获取某个时间的时间戳
}

// type time
func timeDemo() {
	// 程序中应使用Time类型值来保存和传递时间，而不能用指针。
	// 就是说，表示时间的变量和字段，应为time.Time类型，而不是*time.Time.类型。
	// 一个Time类型值可以被多个go程同时使用。
	// 时间点可以使用Before、After和Equal方法进行比较。
	// Sub方法让两个时间点相减，生成一个Duration类型值（代表时间段）。
	// Add方法给一个时间点加上一个时间段，生成一个新的Time类型时间点。

	// 每一个时间都具有一个地点信息（及对应地点的时区信息），
	// 当计算时间的表示格式时，如Format、Hour和Year等方法，都会考虑该信息。
	// Local、UTC和In方法返回一个指定时区（但指向同一时间点）的Time。
	// 修改地点/时区信息只是会改变其表示；不会修改被表示的时间点，因此也不会影响其计算。

	fmt.Println("********  init 一个 time.Time")

	t1 := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Date: ", t1)

	t2 := time.Now()
	fmt.Println("Now: ", t2)

	// 解析一个时间
	// layout是模板时间，value是和模板时间*相同格式*的某个时间点
	// 返回的是时间为value的 time.Time类型
	t3, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println("Parse1: ", t3)
	t4, _ := time.Parse("2006-01-02 15:04:05", "2021-02-01 23:29:00")
	fmt.Println("Parse2: ", t4)

	// ParseInLocation类似Parse但有两个重要的不同之处。
	// 第一，当缺少时区信息时，Parse将时间解释为UTC时间，而ParseInLocation将返回值的Location设置为loc；
	// 第二，当时间字符串提供了时区偏移量信息时，Parse会尝试去匹配本地时区，而ParseInLocation会去匹配loc。
	t5, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-02-01 23:29:00", time.Local)
	fmt.Println("ParseInLocation: ", t5)

	// 用时间戳获取一个time实例
	t6 := time.Unix(0, 0)
	fmt.Println("Unix: ", t6)

	fmt.Println()
	fmt.Println("***** time实例的对应函数")

	fmt.Println("t6所在时区: ", t6.Location())

	zoneName, offset := t6.Zone()
	fmt.Println("时区名: ", zoneName, " 偏移: ", offset)
}

// type Timer and type ticker
func timerTicker() {
	fmt.Println("开始时间：", time.Now().Format(myTimeFormat))

	// 每秒提醒1次
	ticker1 := time.NewTicker(1 * time.Second)
	go func() {
		for value := range ticker1.C {
			fmt.Println("心跳: ", value)
		}
	}()

	// 创建一个7秒后执行的函数
	time.AfterFunc(7*time.Second, func() {
		fmt.Println("计时函数执行， 当前时间：", time.Now().Format(myTimeFormat))
	})

	// 计时10秒， 10秒后向通道返回当前时间
	timer2 := time.NewTimer(10 * time.Second)
	t := <-timer2.C
	fmt.Println("计时完毕，当前时间：", t.Format(myTimeFormat))

}

// type time.Weekday. 返回某日是星期几，0代表周日,以此类推;
func weekday() {
	fmt.Println(time.Weekday.String(0))
}

// type time.Month.月份从1开始
func month() {
	m := time.Month(1)
	fmt.Println(m.String())
}

// type time.Location
func location() {
	// func LoadLocation
	// 参数是某个时区(如： Asia/Shanghai)或者Local
	// 此时LoadLocation会查找环境变量ZONEINFO指定目录或解压该变量指定的zip文件
	// 然后查找Unix系统的惯例时区数据安装位置，最后查找$GOROOT/lib/time/zoneinfo.zip。
	loc, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	fmt.Println(loc) // Local

	// 或者直接
	fmt.Println(time.Local)
}
