package sort

import (
	"math"
	"math/rand"
)

// 虽然C语言建议将数组长度len()取值一次，然后再循环，但是Go中len取的是

// 插入排序
// 理解起来就像抓扑克牌一样，将乱序的几张牌一张张抓到手上，抓起过程中排序
func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		j := i - 1
		// 将nums[i]插入到升序序列nums[0...i-1]
		for j >= 0 && nums[j] > tmp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = tmp
	}

	return nums
}

func OptInsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		left := 0
		right := i - 1
		// 将nums[i]插入到升序序列nums[0...i-1]
		// 既然nums[0...i-1]已经是升序排列的，那我们直接用二分法插入应可以减少不少时间
		for left <= right {
			mid := left + (right-left)/2
			if tmp < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		for j := i - 1; j >= left; j-- {
			nums[j+1] = nums[j]
		}
		nums[left] = tmp
	}

	return nums
}

// 选择排序
// 选择数组中最小的值和第一个元素互换，然后选择剩下数组中最小的数和第二个元素互换，以此类推...
func SelectionSort(nums []int) []int {
	for i := range nums {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	return nums
}

// 冒泡排序
// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 针对所有的元素重复以上的步骤，除了最后一个。
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
func BubbleSort(nums []int) []int {
	for i := range nums {
		flag := true
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}
		if flag {
			return nums
		}
	}

	return nums
}

// 希尔排序，Hibbard序列
// 每次循环，对step,step+1,step+2,...,len-1中每个位置i，把其上的元素放到i-step,i-2step,...中的正确位置上
func getHibbardStepArr(n int) []int {
	arr := []int{}
	tmp := 1
	i := 2
	for tmp <= n {
		arr = append(arr, tmp)
		tmp = (1 << i) - 1
		i++
	}

	return arr
}

func ShellSort(nums []int) []int {
	steps := getHibbardStepArr(len(nums))
	for i := len(steps) - 1; i >= 0; i-- {
		incr := steps[i]
		for j := incr; j < len(nums); j++ {
			k := j
			tmp := nums[j]
			for ; k >= incr; k -= incr {
				if tmp < nums[k-incr] {
					nums[k] = nums[k-incr]
				} else {
					break
				}
			}
			nums[k] = tmp
		}
	}

	return nums
}

// 堆排序
// 1.创建一个堆（最大堆）
// 2.将堆首和堆尾互换
// 堆长度减去1，然后再重置堆
// 再重复2，直到堆大小为1
// 在leecode测试中，效果反而不如shell 排序
func HeapSort(nums []int) []int {
	l := len(nums)

	// build max-heap
	for i := l / 2; i >= 0; i-- {
		heapIFY(nums, i, l)
	}

	// exchange the heapHead and heapTail
	for i := l - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		l--
		heapIFY(nums, 0, l)
	}

	return nums
}

// 保证以i为父节点的子树符合最大堆性质
func heapIFY(nums []int, i, l int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < l && nums[left] > nums[largest] {
		largest = left
	}
	if right < l && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapIFY(nums, largest, l)
	}
}

// 归并排序
func MergeSort(nums []int) []int {
	merge := func(a []int, p, q, r int) {
		lArr := make([]int, 0, q-p+2)
		rArr := make([]int, 0, r-q+1)
		for i := 0; i < q-p+1; i++ {
			lArr = append(lArr, a[p+i])
		}
		lArr = append(lArr, math.MaxInt32)
		for i := 0; i < r-q; i++ {
			rArr = append(rArr, a[q+i+1])
		}
		rArr = append(rArr, math.MaxInt32)
		i := 0
		j := 0
		for k := p; k < r+1; k++ {
			if lArr[i] <= rArr[j] {
				a[k] = lArr[i]
				i++
			} else {
				a[k] = rArr[j]
				j++
			}
		}

	}

	var mergeSort func(a []int, p, r int)
	mergeSort = func(a []int, p, r int) {
		if p < r {
			q := (p + r) / 2
			mergeSort(a, p, q)
			mergeSort(a, q+1, r)
			merge(a, p, q, r)
		}
	}
	mergeSort(nums, 0, len(nums)-1)

	return nums
}

// 快速排序
func QuickSort(nums []int) []int {
	partition := func(nums []int, p, r int) int {
		tmp := nums[r]
		i := p - 1
		for j := p; j < r; j++ {
			if nums[j] <= tmp {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[i+1], nums[r] = nums[r], nums[i+1]

		return i + 1
	}

	randomP := func(nums []int, p, r int) int {
		i := rand.Intn(r-p) + p
		nums[i], nums[r] = nums[r], nums[i]
		return partition(nums, p, r)
	}

	var quickSort func(nums []int, p, r int)
	quickSort = func(nums []int, p, r int) {
		if p < r {
			q := randomP(nums, p, r)
			quickSort(nums, p, q-1)
			quickSort(nums, q+1, r)
		}
	}

	quickSort(nums, 0, len(nums)-1)

	return nums
}
