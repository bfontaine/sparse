package sparse

func minOrder(a, b int64) (int64, int64) {
	if b < a {
		a, b = b, a
	}
	return a, b
}
