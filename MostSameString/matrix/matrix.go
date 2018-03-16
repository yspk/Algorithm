package matrix

type Matrix struct {
	rows, columns int
	data          []int
}

func (A *Matrix) Set(r, c, val int) {
	A.data[findIndex(r, c, A)] = val
}

func (A *Matrix) Get(r, c int) int {
	return A.data[findIndex(r, c, A)]
}

func findIndex(r, c int, A *Matrix) int {
	return (r-1)*A.columns + (c - 1) // begin from 1
}

func NewMatrix(r, c int) *Matrix {
	return &Matrix{
		r,
		c,
		make([]int, r*c),
	}
}

type Matrix2 struct {
	rows, columns int
	data          []string
}

func (B *Matrix2) Set(r, c int, val string) {
	B.data[findIndex2(r, c, B)] = val
}

func (B *Matrix2) Get(r, c int) string {
	return B.data[findIndex2(r, c, B)]
}

func findIndex2(r, c int, B *Matrix2) int {
	return (r-1)*B.columns + (c - 1) // begin from 1
}

func NewMatrix2(r, c int) *Matrix2 {
	return &Matrix2{
		r,
		c,
		make([]string, r*c),
	}
}
