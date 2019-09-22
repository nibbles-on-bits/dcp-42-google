package main

import(
	"fmt"
)

func main() {
	fmt.Println("Iterative Combinations Playground")
	fmt.Println("please run unit tests")
}

func iterate_combinations(elems []int) {
	n := len(elems)
	for num:=0;num < (1 << uint(n));num++ {
		combination := []int{}
		for ndx:=0;ndx<n;ndx++ {
			// (is the bit "on" in this number?)
			if num & (1 << uint(ndx)) != 0 {
				// (then add it to the combination)
				combination = append(combination, elems[ndx])
			}
		}

		fmt.Println(combination)
	}
}