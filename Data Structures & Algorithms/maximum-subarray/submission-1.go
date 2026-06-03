func maxSubArray(nums []int) int {
    maxSum, currentSum := -math.MaxInt, 0

    for _, n := range nums {
        currentSum = max(currentSum, 0)
        currentSum += n

        maxSum = max(maxSum, currentSum)
    }

    return maxSum
}
