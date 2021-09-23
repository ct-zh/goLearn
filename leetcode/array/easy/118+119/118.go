package _118

func generate(numRows int) (arr [][]int) {
	if numRows <= 0 {
		return
	}

	for i := 1; i <= numRows; i++ {
		child := make([]int, i)
		for j := 0; j < i; j++ {
			value := 1
			// 先判断上一级是否存在，再判断 左上角 和 右上角是否合法
			// i=3 j=1 左上为：arr[i-2][0] 右上为arr[i-2][1]
			// 左上通项： arr[i-2][j-1] 右上通项arr[i-2][j]
			if i-2 > 0 && j-1 >= 0 && j < i-1 {
				//fmt.Printf("该值是第%d层第%d个数, 左上角的值为arr[%d][%d] 右上角数为arr[%d][%d]\n", i, j+1, i-2, j-1, i-2, j)
				value = arr[i-2][j-1] + arr[i-2][j]
			}
			child[j] = value
		}
		//fmt.Println(child)
		arr = append(arr, child)
	}

	return
}
