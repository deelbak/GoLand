package mathmode

func Fac(n int) int {
	res := 1
	for n >= 1 {
		res *= n
		n--
	}
	return res

}
