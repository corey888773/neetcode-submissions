func search(nums []int, target int) int {
    n := len(nums)
    lo, hi := 0, n-1

    for lo < hi {
        mid := lo + (hi - lo) / 2

        if nums[mid] < nums[hi]  {
            hi = mid
        } else {
            lo = mid+1
        }
    }

    searchB := func(lo, hi int) int {
        for lo < hi {
            mid := lo + (hi - lo) / 2

            if nums[mid] >= target {
                hi = mid
            } else {
                lo = mid+1
            }
        }

        if nums[lo] == target {
            return lo
        } else {
            return -1
        }
    }
    
    return max(searchB(0, lo-1), searchB(lo, n-1))
}




