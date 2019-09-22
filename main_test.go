package main

import(
	"testing"
	"sort"
)

func Test1(t *testing.T) {
	s := []int{12,1,61,5,9,2}
	k := 24
	want := []int{12,9,2,1}
	got := GetSubset(s, k)

	if !ContainsSameElements(got,want) {
		t.Errorf("test1 : got=%v  want=%v\n", got, want)
	}

	s = []int{2,8,30,5,9,2,6,1,3,6,12,15,98}
	k = 59
	want = []int{1,6,9,5,30,8}
	got = GetSubset(s, k)

	if !ContainsSameElements(got,want) {
		t.Errorf("test 2 : got=%v  want=%v\n", got, want)
	}

}

func TestContainsSameElements(t *testing.T) {
	ia1 := []int{1,2,3,4,5}
	ia2 := []int{5,2,1,3,4}

	if !ContainsSameElements(ia1, ia2) {
		t.Errorf("arrays aren't equavilent")
	}
}

func ContainsSameElements (ia1 []int, ia2 []int) bool {
	if len(ia1) != len(ia2) {
		return false
	}

	sort.Ints(ia1)
	sort.Ints(ia2)
	for i, v := range ia1 { if v != ia2[i] { return false } }
	return true
}