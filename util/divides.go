package util

func Divides(n int) []int {
	as := make([]int, 0)

	for m := 1; m <= n; m++ {
		if GCD(n, m) == m {
			as = append(as, m)
		}
	}

	return as
}
