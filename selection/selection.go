package selection

import "math/rand"

func MaxInSlice(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func MinInSlice(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func MaxAndMinInSlice(nums []int) (int, int) {
	var max, min, start int
	if len(nums)%2 == 0 {
		if nums[0] > nums[1] {
			max, min = nums[0], nums[1]
		} else {
			max, min = nums[1], nums[2]
		}
		start = 2
	} else {
		max, min = nums[0], nums[0]
		start = 1
	}
	println("hello")

	for i := start; i < len(nums); i += 2 {
		if nums[i] > nums[i+1] {
			if nums[i] > max {
				max = nums[i]
			}
			if nums[i+1] < min {
				min = nums[i+1]
			}
		} else {
			if nums[i] < min {
				min = nums[i]
			}
			if nums[i+1] > max {
				max = nums[i+1]
			}
		}
	}

	return max, min
}

func Select(nums []int, i int) int {
	partition := func(nums []int, p, r int) int {
		tmp := nums[r]
		i := p
		for j := p; j < r; j++ {
			if nums[j] <= tmp {
				if i != j {
					nums[i], nums[j] = nums[j], nums[i]
				}
				i++
			}
		}
		nums[i], nums[r] = nums[r], nums[i]

		return i
	}

	randomP := func(nums []int, p, r int) int {
		i := rand.Intn(r-p) + p
		nums[i], nums[r] = nums[r], nums[i]
		return partition(nums, p, r)
	}

	var select1 func(nums []int, p, r, i int) int
	select1 = func(nums []int, p, r, i int) int {
		if p == r {
			return nums[p]
		}

		q := randomP(nums, p, r)
		k := q - p + 1
		if i == k {
			return nums[q]
		} else if i < k {
			return select1(nums, p, q-1, i)
		} else {
			return select1(nums, q+1, r, i-k)
		}
	}

	return select1(nums, 0, len(nums)-1, i)
}
