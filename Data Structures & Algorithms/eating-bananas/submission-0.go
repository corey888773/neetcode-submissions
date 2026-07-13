func minEatingSpeed(piles []int, h int) int {
    maxPile := 0
    for _, pile := range piles {
        maxPile = max(maxPile, pile)
    }
    
    lo, hi := 1, maxPile+1

    canEatAll := func(k int) bool {
        actualHours := 0

        for _, pile := range piles {
            if pile % k == 0 {
                actualHours += pile / k
            } else {
                actualHours += pile / k + 1
            }
        }

        return actualHours <= h
    }

    for lo < hi {
        mid := lo + (hi - lo) / 2

        if canEatAll(mid) {
            hi = mid
        } else {
            lo = mid+1
        }
    }

    return lo
}

