package arr

func AllAreEqual(vals ...int) bool {
	for _, v := range vals {
		if v != vals[0] {
			return false
		}
	}
	return true
}
