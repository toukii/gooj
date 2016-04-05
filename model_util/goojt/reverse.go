package goojt

func reverse(in []int) []int {
	leng := len(in)
	l := leng / 2
	for i := 0; i < l; i++ {
		in[i], in[leng-1-i] = in[leng-1-i], in[i]
	}
	return in
}