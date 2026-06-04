func findDuplicate(nums []int) int {
    fast, slow := nums[0], nums[0]

    for {
        slow = nums[slow]
        fast = nums[nums[fast]]

        if slow == fast {
            break
        }
    }

    slow = nums[0]
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }

    return slow
}