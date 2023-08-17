package sliceutil

// RemoveString remove string from slice if function return true.
func RemoveString(slice []string, remove func(s string) bool) []string {
	for i := 0; i < len(slice); i++ {
		if remove(slice[i]) {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}

	return slice
}

// FindString return true if target in slice, return false if not.
func FindString(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}

	return false
}

// FindInt return true if target in slice, return false if not.
func FindInt(slice []int, target int) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}

	return false
}

// FindUint return true if target in slice, return false if not.
func FindUint(slice []uint, target uint) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}
