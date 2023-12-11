package pkg

func contains[T comparable](s []T, i T) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func indexOf[T comparable](s []T, i T) int {
	for k, v := range s {
		if v == i {
			return k
		}
	}
	return -1
}
