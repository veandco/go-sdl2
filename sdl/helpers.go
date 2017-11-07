package sdl

// Btoi returns 0 or 1 according to the value of b.
func Btoi(b bool) int {
	if b == true {
		return 1
	}

	return 0
}
