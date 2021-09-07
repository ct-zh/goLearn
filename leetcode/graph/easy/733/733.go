package _33

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	org := image[sr][sc]
	if newColor == org {
		return image
	}
	return _floodFill(image, sr, sc, newColor, org)
}

func _floodFill(image [][]int, sr int, sc int, newColor int, orgColor int) [][]int {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) {
		return image
	}
	if image[sr][sc] != orgColor {
		return image
	}
	image[sr][sc] = newColor

	_floodFill(image, sr+1, sc, newColor, orgColor)
	_floodFill(image, sr-1, sc, newColor, orgColor)
	_floodFill(image, sr, sc+1, newColor, orgColor)
	_floodFill(image, sr, sc-1, newColor, orgColor)
	return image
}
