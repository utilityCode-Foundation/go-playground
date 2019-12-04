package main

import "fmt"

func main() {
	s := make([]int, 3)

	// Similar to array:
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println(s)

	fmt.Println(s[0])
	fmt.Println(len(s))

	// append function is unique to slices
	s = append(s, 4)
	s = append(s, 5, 6)
	fmt.Println(s)

	// slice syntax
	fmt.Println(s[1:3])

	fmt.Println(s[:3])

	fmt.Println(s[1:])

	// concise slice definition:
	t := []int{100, 200, 300}
	fmt.Println(t)

	//x := s
	//x[0] = 500
	//fmt.Println(x)
	//fmt.Println(s)

	// Use copy to prevent from changing both x and s
	x := make([]int, len(s))
	copy(x, s)

	x[0] = 500
	fmt.Println(x)
	fmt.Println(s)

	// 2-D slices (similar to arrays, although lengths can vary.
	ss := make([][]int, 3)
	for i := 0; i < 3; i++ {

		ss[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			ss[i][j] = j + 1
		}
	}
	fmt.Println(ss)

	// 3-D slices (similar to arrays, although lengths can vary.
	sss := make([][][]int, 3)
	for i := 0; i < 3; i++ {

		sss[i] = make([][]int, 6)
		for j := 0; j < 6; j++ {
			sss[i][j] = make([]int, 10)
			for k := 0; k < 10; k++ {
				sss[i][j][k] = k + 1
			}
		}
	}
	fmt.Println(sss)

}
