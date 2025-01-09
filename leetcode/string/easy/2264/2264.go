// 2264. Largest 3-Same-Digit Number in String
// https://leetcode.cn/problems/largest-3-same-digit-number-in-string
// 2025-01-08
package main

func largestGoodInteger(num string) string {
	max := byte(0)
	for i := 0; i <= len(num)-3; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			if num[i] > max {
				max = num[i]
			}
		}
	}
	if max == 0 {
		return ""
	}
	return string(max) + string(max) + string(max)
}
