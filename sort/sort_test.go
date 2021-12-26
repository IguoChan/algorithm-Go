package sort

import (
	"math/rand"
	"sort"
	"testing"
)

const MaxCount = 10000

var sortSlice []int

func TestCountingSort(t *testing.T) {
	a := []int{2, 5, 3, 0, 2, 3, 0, 3}
	b := CountingSort(a, 5)
	t.Log(b)

	aa := make([]int, 100)
	for i := range aa {
		aa[i] = rand.Intn(50)
	}

	t.Log(aa)
	bb := CountingSort(aa, 50)
	t.Log(bb)
}

func BenchmarkSelectionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelectionSort(sortSlice)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InsertSort(sortSlice)
	}
}

func BenchmarkOptInsertSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OptInsertSort(sortSlice)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BubbleSort(sortSlice)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		HeapSort(sortSlice)
	}
}

// 实测中shell排序的效果好的不行
func BenchmarkShellSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ShellSort(sortSlice)
	}
}

func TestQuickSort(t *testing.T) {
	t.Log("[reverse pair]: ", getReversePair(sortSlice))
	bb := QuickSort(sortSlice)
	if !checkSorted(bb) {
		t.Errorf("wrong sort")
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		QuickSort(sortSlice)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(sortSlice)
	}
}

func BenchmarkCountingSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountingSort(sortSlice, MaxCount)
	}
}

// 测测go包标准的排序算法
func BenchmarkStandardSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sort.Ints(sortSlice)
	}
}

func init() {
	sortSlice = make([]int, 100000)
	for i := range sortSlice {
		// 伪随机，出来的是同样的数据，更可以考察不同排序的性能
		sortSlice[i] = rand.Intn(MaxCount)
	}
}

func checkSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func getReversePair(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				cnt++
			}
		}
	}

	return cnt
}
