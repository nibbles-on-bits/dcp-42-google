package main

import(
	"fmt"
)

// GetSubset will attempt to find a subset of int's in s that add up to k
func GetSubset(s []int, k int) []int{
	var ret []int

	fmt.Printf("GetSubset() s=%v k=%d\n", s, k)

	// first thing we want to do is remove anything > k
	//var tmp []int {}
	tmp := []int{}

	for i,num := range(s) {
		fmt.Printf("i=%v  num=%v  \n", i, num)
		if num <= k {
			tmp = append(tmp,num)
		}
	}

	// Now tmp only contains ints < k.  Now perform a combinations algorithm
	// 

	n := len(tmp)
	for num:=0;num < (1 << uint(n));num++ {
		combination := []int{}
		sum := 0
		for ndx:=0;ndx<n;ndx++ {
			// (is the bit "on" in this number?)
			if num & (1 << uint(ndx)) != 0 {
				// (then add it to the combination)
				sum += tmp[ndx]
				if sum > k {
					continue
				}
				combination = append(combination, tmp[ndx])
				if sum==k {
					fmt.Printf("----------------------->%v\n",combination)
					return combination
				}
			}
		}
	}

	fmt.Printf("tmp=%v\n", tmp)

	return ret
}