package main

import (
	"fmt"
)

type Temp struct {
	maybe   []int
	already map[int]bool
}

var input [][]int
var last int

func main() {
	input = make([][]int, 9)
	for i := 0; i < 9; i++ {
		input[i] = make([]int, 9)
	}
	var temp [][]*Temp
	//糯米
	input[0][0] = 1
	input[0][3] = 5

	input[1][2] = 4
	input[1][5] = 2
	input[1][6] = 3

	input[2][1] = 3
	input[2][4] = 1
	input[2][7] = 8

	input[3][0] = 2
	input[3][3] = 3
	input[3][5] = 8
	input[3][7] = 7

	input[4][2] = 9
	input[4][6] = 4

	input[5][1] = 1
	input[5][3] = 2
	input[5][5] = 9
	input[5][8] = 5

	input[6][1] = 8
	input[6][4] = 6
	input[6][7] = 1

	input[7][2] = 5
	input[7][3] = 4
	input[7][6] = 9

	input[8][5] = 3
	input[8][8] = 6

	input[3][1] = 4
	//guess
	input[7][5] = 1
	input[7][4] = 2
	input[7][1] = 6

	input[8][7] = 5

	//input[0][0] = 8
	//
	//input[1][2] = 7
	//input[1][3] = 5
	//input[1][8] = 9
	//
	//input[2][1] = 3
	//input[2][6] = 1
	//input[2][7] = 8
	//
	//input[3][1] = 6
	//input[3][5] = 1
	//input[3][7] = 5
	//
	//input[4][2] = 9
	//input[4][4] = 4
	//
	//input[5][3] = 7
	//input[5][4] = 5
	//
	//input[6][2] = 2
	//input[6][4] = 7
	//input[6][8] = 4
	//
	//input[7][5] = 3
	//input[7][6] = 6
	//input[7][7] = 1
	//
	//input[8][6] = 8
	////guess
	//input[6][7] = 9
	//input[6][6] = 3
	//input[1][6] = 2
	//input[0][6] = 4
	//input[1][1] = 4
	//input[8][7] = 2
	//input[8][8] = 5
	//input[7][2] = 4
	//input[7][0] = 5
	//input[7][1] = 8

	Calc(input, temp)
}

func initTemp(temp [][]*Temp) [][]*Temp {
	temp = make([][]*Temp, 9)
	for i := 0; i < 9; i++ {
		temp[i] = make([]*Temp, 9)
		for j := 0; j < 9; j++ {
			temp[i][j] = &Temp{
				maybe:   make([]int, 0),
				already: make(map[int]bool),
			}
		}
	}
	return temp
}

func Calc(in [][]int, temp [][]*Temp) {
	temp = CalcTemp(in, temp)
	var count int
	for i, v := range temp {
		for j, m := range v {
			if len(m.maybe) == 0 {
				count++
			}
			if len(m.maybe) == 1 {
				in[i][j] = m.maybe[0]
				count++
				fmt.Printf("input[%d][%d] is %d\n", i, j, in[i][j])
			}
			if len(m.maybe) == 2 {
				fmt.Printf("input[%d][%d] maybe %d\n", i, j, m.maybe)
			}

		}
	}
	fmt.Printf("alread set %d, total 81\n", count)
	if count == 81 {
		printInput(in)
		return
	} else if count == last {
		printInput(in)
		return
	} else {
		last = count
		Calc(in, temp)
	}
}

func CalcTemp(in [][]int, temp [][]*Temp) [][]*Temp {
	temp = initTemp(temp)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if in[i][j] == 0 {
				for m := 0; m < 9; m++ {
					if in[m][j] != 0 {
						temp[i][j].already[in[m][j]] = true
					}
				}
				for n := 0; n < 9; n++ {
					if in[i][n] != 0 {
						temp[i][j].already[in[i][n]] = true
					}
				}
				k, p := i/3, j/3
				for r := k * 3; r < k*3+3; r++ {
					for s := p * 3; s < p*3+3; s++ {
						if in[r][s] != 0 {
							temp[i][j].already[in[r][s]] = true
						}
					}
				}
				//对角线
				//if i == j {
				//	for r := 0; r < 9; r++ {
				//		if in[r][r] != 0 {
				//			temp[i][j].already[in[r][r]] = true
				//		}
				//	}
				//}
				//if i+j == 8 {
				//	for r := 0; r < 9; r++ {
				//		if in[r][8-r] != 0 {
				//			temp[i][j].already[in[r][8-r]] = true
				//		}
				//	}
				//}

				for l := 1; l <= 9; l++ {
					if !temp[i][j].already[l] {
						temp[i][j].maybe = append(temp[i][j].maybe, l)
					}
				}
				//fmt.Println(temp[i][j].maybe, temp[i][j].already)
			}
		}
	}
	return temp
}

func printInput(input [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", input[j][i])
		}
		fmt.Printf("\n")
	}
}
