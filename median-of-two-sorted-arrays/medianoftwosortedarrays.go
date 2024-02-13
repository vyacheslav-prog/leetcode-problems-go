package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Length, nums2Length := len(nums1), len(nums2)
	if nums1Length != 0 {
		if nums2Length != 0 {
			nums1First, nums2First := nums1[0], nums2[0]
			return float64(nums1First+nums2First) / 2
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
