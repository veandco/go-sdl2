package sdl

func btoi(b bool) int {
	if b == true {
		return 1
	}

	return 0
}

func itob(i int) bool {
	if i > 0 {
		return true
	}

	return false
}
