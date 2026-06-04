func findDuplicate(nums []int) int {
    n := len(nums) - 1
    res := 0

    for bit := range 32 {
        mask := 1 << bit
        expected, actual := 0, 0

        for i := 1; i <= n; i++ {
            if i & mask != 0 {
                expected++
            }
        }

        for _, num := range nums {
            if num & mask != 0 {
                actual++
            }
        }

        if actual > expected {
            res = res | mask
        }
    }

    return res
}