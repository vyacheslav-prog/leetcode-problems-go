package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArraysSumLengthComplexity(nums1, nums2)
}

func findMedianSortedArraysSumLengthComplexity(nums1 []int, nums2 []int) float64 {
	nums1Length, nums2Length := len(nums1), len(nums2)
	if nums1Length == 0 {
		if nums2Length != 0 {
			return findMedianSortedArraysSumLengthComplexity(nums2, nums1)
		}
		return 0
	} else if nums2Length != 0 {
		if nums2[0] < nums1[0] {
			return findMedianSortedArraysSumLengthComplexity(nums2, nums1)
		}
	}
	var medianNums []int
	var nums1Pointer, nums2Pointer int
	halfTotalLength := (nums1Length + nums2Length) / 2
	if (halfTotalLength*2) != (nums1Length + nums2Length) {
		halfTotalLength += 1
	}
	for totalCounter := 0; totalCounter != halfTotalLength; totalCounter += 1 {
		if nums1Pointer != nums1Length {
			if nums2Pointer == nums2Length {
				nums1Pointer += 1
			} else if nums1[nums1Pointer] <= nums2[nums2Pointer] {
				nums1Pointer += 1
			}
		}
		if totalCounter == halfTotalLength {
			break
		}
		if nums2Pointer != nums2Length {
			if nums1Pointer == nums1Length {
				nums2Pointer += 1
			} else if nums2[nums2Pointer] <= nums1[nums1Pointer] {
				nums2Pointer += 1
			}
		}
	}
	medianNumsLength := len(medianNums)
	if medianNumsLength == 2 {
		return float64(medianNums[0] + medianNums[1]) / 2
	} else if medianNumsLength == 1 {
		return float64(medianNums[0])
	}
	return 0
}


func findMedianSortedArraysConstantComplexity(nums1 []int, nums2 []int) float64 {
	nums1Length, nums2Length := len(nums1), len(nums2)
	if nums1Length != 0 {
		if nums2Length != 0 {
			nums1First, nums2First := nums1[0], nums2[0]
			nums1Last, nums2Last := nums1[nums1Length-1], nums2[nums2Length-1]
			var maxNumber, minNumber int
			if nums1First < nums2First {
				minNumber = nums1First
			} else {
				minNumber = nums2First
			}
			if nums1Last < nums2Last {
				maxNumber = nums2Last
			} else {
				maxNumber = nums1Last
			}
			return float64(minNumber+maxNumber) / 2
		} else {
			nums1HalfLength := nums1Length / 2
			if nums1HalfLength == 0 {
				return float64(nums1[0])
			}
			if nums1HalfLength*2 != nums1Length {
				return float64(nums1[nums1HalfLength])
			}
			return float64(nums1[nums1HalfLength-1]+nums1[nums1HalfLength]) / 2
		}
	} else if nums2Length != 0 {
		return findMedianSortedArrays(nums2, nums1)
	}
	return 0
}
