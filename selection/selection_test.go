package selection

import (
	"math/rand"
	"testing"
)

var sortSlice []int

func TestMaxInSlice(t *testing.T) {
	nums := make([]int, 10)
	for i := range nums {
		nums[i] = rand.Int()
	}

	t.Log(nums)
	max := MaxInSlice(nums)
	t.Log(max)
}

func BenchmarkMaxInSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = MaxInSlice(sortSlice)
	}
}

func TestMaxAndMinInSlice(t *testing.T) {
	nums := make([]int, 10)
	for i := range nums {
		nums[i] = rand.Intn(150)
	}

	t.Log(nums)
	max, min := MaxAndMinInSlice(nums)
	t.Log(max, min)
}

func TestSelect(t *testing.T) {
	nums := make([]int, 10)
	for i := range nums {
		nums[i] = rand.Intn(150)
	}
	t.Log(nums)
	for i := range nums {
		s := Select(nums, i+1)
		t.Log(s)
	}

}

func init() {
	sortSlice = make([]int, 100000)
	for i := range sortSlice {
		// 伪随机，出来的是同样的数据，更可以考察不同排序的性能
		sortSlice[i] = rand.Int()
	}
}
