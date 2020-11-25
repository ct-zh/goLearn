package _17

// 时间复杂度 3^n= O(2^n)
func letterCombinations(digits string) (res []string) {
	if digits == "" {
		return res
	}

	findCombination([]byte(digits), 0, "", &res)

	return
}

// 手机上数字与字母的映射关系
func getLetterMap() map[int]string {
	return map[int]string{
		0: " ",
		1: "",
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
}

// digits: 处理的数字字符串; index: 当前处理的数字在digits里的索引
// s: 保存了此时digits[0...index-1]翻译得到的一个字母字符串
// 寻找和digits[index]匹配的字母，获得digits[0...index]翻译得到的解
func findCombination(digits []byte, index int, s string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, s)
		return
	}

	c := digits[index] // 当前数字只能在2-9范围内或者等于0
	if !(c >= '0' && c <= '9' && c != '1') {
		panic("error")
	}

	letters := []byte(getLetterMap()[int(c-'0')]) // 当前数字对应的字母byte数组
	for i := 0; i < len(letters); i++ {
		findCombination(digits, index+1, s+string(letters[i]), result)
	}

	return
}
