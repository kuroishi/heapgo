package main

import (
	"fmt"
	"os"
	"math"
)

func min_heapify(s []int, idx int) {
	length := len(s)
	if idx < 0 || (length / 2) <= idx {
		return
	}
	fmt.Print("----\n")
	display_pyramid(s)
	idx_left_chi  := 2 * idx + 1
	idx_right_chi := 2 * idx + 2
	if length <= idx_right_chi || s[idx_left_chi] <= s[idx_right_chi] {
		if  s[idx_left_chi] < s[idx] {
			s[idx], s[idx_left_chi] = s[idx_left_chi], s[idx]
			min_heapify(s, idx_left_chi)
		}
	} else {
		if  s[idx_right_chi] < s[idx] {
			s[idx], s[idx_right_chi] = s[idx_right_chi], s[idx]
			min_heapify(s, idx_right_chi)
		}
	}
}

func get_height(s []int) int {
	len := len(s)
	min := 0
	cap := 1
	hgt := 1

	for {
		//fmt.Print("min: ", min,
		//" cap: ", cap,
		//" min + cap: ", min + cap, "\n")
		if min <= len && len <= (min + cap) {
			break;
		}
		min = min + cap
		cap = cap * 2
		hgt++
	}

	return hgt
}

func display_padding(n int) {
	//x := (n / 2) + (n / 2) - 1
	x := n - 1
	for x > 0 {
		x--
		fmt.Print("  ")
	}
}

func display_pyramid(s []int) {
	var len int = len(s)
	var idx int = 0
	var min int = 0
	var cap int = 1
	var hgt int = get_height(s)
	var cur int = 1

	for idx < len {
		display_padding(int(math.Pow(2, float64(hgt) - float64(cur))))
		fmt.Print(s[idx])
		display_padding(int(math.Pow(2, float64(hgt) - float64(cur))))
		idx++
		if idx == (min + cap) {
			fmt.Print("\n")
			min = min + cap
			cap = cap * 2
			cur++
		} else {
			fmt.Print("  ")
		}
	}
	fmt.Print("\n")
}

func build_heap(s []int) {
	len := len(s)
	for cnt := len / 2; 0 <= cnt; cnt-- {
		min_heapify(s, cnt - 1)
		display_pyramid(s)
	}
}

func main() {
	var num int
	s := make([]int, 0)
	for {
		_, err := fmt.Scanln(&num)
		if err != nil {
			break
		}
		s = append(s, num)
	}

	//fmt.Printf("%v\n", s)
	fmt.Print(s, "\n")
	length := len(s)
	fmt.Printf("length: %d\n", length)
	display_pyramid(s)

	build_heap(s)

	cnt := 0
	a := make([]int, 0)
	for length > 0 {
		cnt++
		fmt.Print("\n==== ", cnt, " ====\n")
		last := length - 1
		s[0], s[last] = s[last], s[0]
		a = append(a, s[last])
		s = s[:last]
		length = len(s)
		min_heapify(s, 0)
	}
	fmt.Print(a, "\n")
	//display_pyramid(a)
	os.Exit(0)

}
