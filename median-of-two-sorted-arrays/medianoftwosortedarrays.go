package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArraysSumLengthComplexity(nums1, nums2)
}

func findMedianSortedArraysSumLengthComplexity(nums1 []int, nums2 []int) float64 {
	nums1Length, nums2Length := len(nums1), len(nums2)
	totalLength := nums1Length + nums2Length
	var counterLimit, nums1Pointer, nums2Pointer int
	counterLimit, isOddTotalLength := (totalLength/2)+1, (totalLength%2) != 0
	mergedNums := []int{}
	for totalCounter := 0; totalCounter != counterLimit; totalCounter += 1 {
		if nums1Pointer != nums1Length && nums2Pointer != nums2Length {
			if nums1[nums1Pointer] < nums2[nums2Pointer] {
				mergedNums = append(mergedNums, nums1[nums1Pointer])
				nums1Pointer += 1
			} else {
				mergedNums = append(mergedNums, nums2[nums2Pointer])
				nums2Pointer += 1
			}
		} else if nums1Pointer != nums1Length {
			mergedNums = append(mergedNums, nums1[nums1Pointer])
			nums1Pointer += 1
		} else if nums2Pointer != nums2Length {
			mergedNums = append(mergedNums, nums2[nums2Pointer])
			nums2Pointer += 1
		}
	}
	if len(mergedNums) > 1 {
		if isOddTotalLength {
			return float64(mergedNums[len(mergedNums)-1])
		}
		return float64(mergedNums[len(mergedNums)-1]+mergedNums[len(mergedNums)-2]) / 2
	}
	if len(mergedNums) == 1 {
		return float64(mergedNums[0])
	}
	return 0
}
