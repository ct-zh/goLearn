package _842

import (
	"math"
)

func splitIntoFibonacci(s string) (res []int) {
	backtrack(0, s, &res)
	return
}

func backtrack(index int, s string, res *[]int) bool {
	// 达到字符串末尾，如果结果集数量小于3个，说明不存在斐波拉契序列
	if len(s) == index {
		return len(*res) > 2
	}

	cur := 0
	// 从index开始循环s
	for i := index; i < len(s); i++ {
		// 如果当前数的开头是0，则不允许后面还有数
		if s[index] == '0' && i > index {
			break
		}

		// s[i]是byte类型，减去byte的0就转换为真正的数字(这里举两个例测试一下就知道了)
		cur = cur*10 + int(s[i]-'0')
		if cur > math.MaxInt32 { // 不能超过最大值
			break
		}

		// 前面已经有至少两个数，说明可以开始判断当前数是否满足F(i)+F(i+1)=F(i+2)
		l := len(*res)
		if l >= 2 {
			sum := (*res)[l-1] + (*res)[l-2]
			if cur < sum { // 小于则继续往cur后面添加数
				continue
			} else if cur > sum { // 大于则说明不符合条件了, 直接return false
				break
			}
		}

		// 将cur加入结果集中
		*res = append(*res, cur)
		// 继续往后判断， 如果全部符合条件则return true，得到答案
		if backtrack(i+1, s, res) {
			return true
		}
		// 否则 则 回溯掉cur
		*res = (*res)[:len(*res)-1]
	}

	return false
}
