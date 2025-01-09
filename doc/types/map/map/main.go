package main

func main() {

	// runtime.makemap
	_ = make(map[int64]struct{}, 100)

	// runtime.mapassign_fast64
	_ = map[int64]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}
}
