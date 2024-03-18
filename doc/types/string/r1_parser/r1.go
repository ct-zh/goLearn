package r1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 4 秒处理 10 亿行数据！ Go 语言的 9 大代码方案，一个比一个快
// https://mp.weixin.qq.com/s/iylZAKZfxLL6SYruww_8zA

// 2024 年开年，Java “十亿行挑战”（1BRC）火爆外网。
// 该挑战赛要求开发者编写一个 Java 程序，
// 从一个包含十亿行信息的文本文件中检索温度测量值，
// 并计算每个气象站的最小、平均值和最高温度。
// “十亿行挑战”的目标是为这项任务创建最快的实现，同时探索现代 Java 的优势。
// 日前，从业 20 年的软件工程师 Ben Hoyt 用 Go 语言参与该挑战，
// 他一共想出了 9 种解决方案，完成 10 亿行数据处理的时间最快只需 4 秒，
// 最慢需要 1 分 45 秒。
// Ben Hoyt 还给自己提了点限制条件：
// 每种方法都仅使用 Go 标准库以保证可移植性，
// 不涉及程序集、不涉及 unsafe、不涉及内存映射文件。
// 跟其他作者的发现相比，Ben Hoyt 的解决方案不是最慢的、但也没能占据榜首。
// 不过最重要的是，他的解跟其他参赛者的思路都不一样，这种独立性可能更具价值。

// r1 1分45秒的版本
func r1(inputPath string, output io.Writer) error {
	type stats struct { // 格式为  地区名;温度
		min, max, sum float64 // 该地区温度最小值、最大值、总和(计算平均温度)
		count         int64
	}

	// 获取文件句柄
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	stationStats := make(map[string]stats)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// 解析每行数据，地区;温度 如果没有;符号代表是非法数据，跳过
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}
		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			return err
		}

		s, ok := stationStats[station]
		if ok {
			s.min = min(s.min, temp)
			s.max = max(s.max, temp)
			s.count++
			s.sum += temp
		} else {
			s.min = temp
			s.max = temp
			s.sum = temp
			s.count = 1
		}
		stationStats[station] = s
	}

	// 给这些气象站地区按照名称排序
	stations := make([]string, 0, len(stationStats))
	for s := range stationStats {
		stations = append(stations, s)
	}
	sort.Strings(stations)

	// 输出数据
	fmt.Fprint(output, "{")
	for i, station := range stations {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}
		s := stationStats[station]
		mean := s.sum / float64(s.count)
		fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f", station, s.min, mean, s.max)
	}
	fmt.Fprint(output, "}\n")
	return nil
}

func r2(inputPath string, output io.Writer) error {
	type stats struct { // 格式为  地区名;温度
		min, max, sum float64 // 该地区温度最小值、最大值、总和(计算平均温度)
		count         int64
	}

	// 获取文件句柄
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	stationStats := make(map[string]*stats) // 使用指针来替代原始变量

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tempBytes := scanner.Bytes() // 不再使用Text()，而是使用bytes获取字节切片

		// 循环获取地区名
		index := 0
		stationBytes := make([]byte, 0)
		for i, tempByte := range tempBytes {
			if tempByte == ';' {
				index = i + 1
				break
			}
			stationBytes = append(stationBytes, tempByte)
		}
		station := string(stationBytes)

		// 使用逐字节解析温度来代替 strconv.ParseFloat
		negative := false
		if tempBytes[index] == '-' {
			index++
			negative = true
		}
		temp := float64(tempBytes[index] - '0') // parse first digit
		index++
		if tempBytes[index] != '.' {
			temp = temp*10 + float64(tempBytes[index]-'0') // parse optional second digit
			index++
		}
		index++                                    // skip '.'
		temp += float64(tempBytes[index]-'0') / 10 // parse decimal digit
		if negative {
			temp = -temp
		}

		s := stationStats[station]
		if s != nil {
			s.min = min(s.min, temp)
			s.max = max(s.max, temp)
			s.count++
			s.sum += temp
		} else {
			s = &stats{
				min:   temp,
				max:   temp,
				sum:   temp,
				count: 1,
			}
		}
		stationStats[station] = s
	}

	// 给这些气象站地区按照名称排序
	stations := make([]string, 0, len(stationStats))
	for s := range stationStats {
		stations = append(stations, s)
	}
	sort.Strings(stations)

	// 输出数据
	fmt.Fprint(output, "{")
	for i, station := range stations {
		if i > 0 {
			fmt.Fprint(output, ", ")
		}
		s := stationStats[station]
		mean := s.sum / float64(s.count)
		fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f", station, s.min, mean, s.max)
	}
	fmt.Fprint(output, "}\n")
	return nil
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
