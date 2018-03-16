package main

import (
	"./Matrix"
	"./Stack"
	"fmt"
)

type Output struct {
	I     int
	J     int
	Value string
}

func main() {
	str1, str2 := "cnblogs", "belong"
	//str1, str2 := "ABA","BBA"
	//str1,str2 := "abcd","egxg"
	//str1, str2 := "acgbfhk", "cegefkh"
	//str1, str2 := "abcbdab", "bdcaba"
	//str1, str2 := "fkjfdaljfdsa71894fdsaf", "ghhjhgdhhhgiuhghfhdgfsdsfdafhhl"
	s := int64(100)
	q := Stack.MakeStack(s)
	m := Matrix.NewMatrix(len(str1)+1, len(str2)+1)
	MostSameStringLength(str1, str2, m)
	m.Print()
	n := Matrix.NewMatrix(len(str1)+1, len(str2)+1)
	LD(str1, str2, n)
	n.Print()
	fmt.Printf("字符串 %s 和 %s 的编辑距离为:%d \n", str1, str2, n.Get(len(str1)+1, len(str2)+1).Length)
	SubSequence(len(str1), len(str2), str1, str2, m, q)
	fmt.Printf("字符串 %s 和 %s最大公共子序列的长度为:%d,序列为:%s \n", str1, str2, m.Get(len(str1)+1, len(str2)+1).Length, Report(q))
}

func MostSameStringLength(str1, str2 string, m *Matrix.Matrix) {
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				m.Set(i+1, j+1, Matrix.Data{m.Get(i, j).Length + 1, "left_up"})
			} else {
				if m.Get(i, j+1).Length >= m.Get(i+1, j).Length {
					m.Set(i+1, j+1, Matrix.Data{m.Get(i, j+1).Length, "left"})
				} else {
					m.Set(i+1, j+1, Matrix.Data{m.Get(i+1, j).Length, "up"})
				}

			}
		}
	}
}

func SubSequence(i, j int, str1, str2 string, m *Matrix.Matrix, q *Stack.Stack) {
	if i == 0 || j == 0 {
		return
	}
	if m.Get(i+1, j+1).Point == "left_up" {
		//fmt.Println(string(str2[j-1]), ": 当前坐标:", i, ",", j)
		q.Push(&Output{i, j, string(str2[j-1])})
		SubSequence(i-1, j-1, str1, str2, m, q)
	} else if m.Get(i+1, j+1).Point == "left" {
		SubSequence(i-1, j, str1, str2, m, q)
	} else {
		SubSequence(i, j-1, str1, str2, m, q)
	}
}

func Report(q *Stack.Stack) []string {
	var out []string
	for !q.IsEmpty() {
		if v, err := q.Pop(); err == nil {
			out = append(out, v.(*Output).Value)
		}
	}
	return out
}

func LD(str1, str2 string, m *Matrix.Matrix) {
	for i := 1; i <= len(str1)+1; i++ {
		m.Set(i, 1, Matrix.Data{i-1, ""})
	}
	for j := 1; j <= len(str2)+1; j++ {
		m.Set(1, j, Matrix.Data{j-1, ""})
	}
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				m.Set(i+1, j+1, Matrix.Data{m.Get(i, j).Length, ""})
			} else {
				min := m.Get(i, j).Length
				if min > m.Get(i+1, j).Length {
					min = m.Get(i+1, j).Length
				}
				if min > m.Get(i, j+1).Length {
					min = m.Get(i, j+1).Length
				}
				m.Set(i+1, j+1, Matrix.Data{min + 1, ""})
			}
		}
	}

}
