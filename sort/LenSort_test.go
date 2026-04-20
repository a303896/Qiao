package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var s1 byLenth[rune] = [][]rune{[]rune("apple"), []rune("banana"), []rune("pear")}
	var s2 byLenth[int] = [][]int{{1, 2, 3, 5, 6, 7}, {1, 2, 3}, {1, 2, 3, 4}}
	sort.Sort(s1)
	var s3 string = string([]rune{'a', 'b', 'c', 'd'})
	fmt.Printf("%v \n", s3)
	fmt.Printf("%v \n", s1)
	fmt.Printf("%s \n", s1)
	sort.Sort(s2)
	fmt.Println(s2)
}
