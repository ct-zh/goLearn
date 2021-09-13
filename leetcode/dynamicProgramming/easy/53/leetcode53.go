package leetcode53

//链接：https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/

// 动态规划解题思路：
// 假设，求f(x)是第x个数结尾的最大子序和；
// 那么问题可以转变为比较f(x)大还是f(x-1)大；并且f(x) = num[x] + f(x-1);
// 状态转移方程为 f(x) = max(f(x), f(x-1)) = max(f(x-1) + num[x], f(x-1))
// 边界条件为 f(0) = num[0]

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

// 分治法解题思路：
// see: readme.md

type Status struct {
	lSum int // 以 l 为左端点的最大子段和
	rSum int // 以 r 为右端点的最大子段和
	mSum int // 最大子段和（要求的值）
	iSum int // 区间和
}

func maxSubArray2(nums []int) int {
	return getMax(nums, 0, len(nums)-1).mSum
}

func getMax(nums []int, l int, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := getMax(nums, l, m)
	rSub := getMax(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func pushUp(l Status, r Status) Status {
	return Status{
		iSum: l.iSum + r.iSum, // 区间和 = 左子区间和 + 右子区间和

		lSum: max(l.lSum, l.iSum+r.lSum),
		rSum: max(r.rSum, r.iSum+l.rSum),

		// 如果最大子序和不跨越m，则取 max(l.mSum, r.mSum)
		// 如果跨越m，则取 l.rSum+r.lSum ()
		mSum: max(max(l.mSum, r.mSum), l.rSum+r.lSum),
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
