package main

import (
	"fmt"
)

//KMP Return index list if str1 contain str2
func KMP(str1 string, str2 string) []int {
	next := preKMP([]byte(str1))
	i := 0
	j := 0
	m := len(str1)
	n := len(str2)

	x := []byte(str1)
	y := []byte(str2)
	var ret []int

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			fmt.Println("find:", j, i)
			i = next[i]
		}
	}

	return ret
}

func preKMP(x []byte) [30]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [30]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		// fmt.Println(i, j)
		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
		// fmt.Println("2", i, j, length)
	}
	return kmpNext
}

func main() {
	fmt.Println("next=>", preKMP([]byte("cocacola")))
	fmt.Println("next=>", preKMP([]byte("gcagagag")))

	fmt.Println("kmp=>", KMP("co", "cocacola"))
	fmt.Println("kmp=>", KMP("abcabc", "zjifeoabcabcjiefjie"))
}
