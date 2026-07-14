func dailyTemperatures(temperatures []int) []int {
    stack := []int{}
    days := make([]int, len(temperatures))
    
    for curr := range len(temperatures) {
        for len(stack) > 0 {
            top := stack[len(stack)-1]
            if temperatures[top] >= temperatures[curr] {
                break
            }

            days[top] = curr - top
            stack = stack[:len(stack)-1]
        }
        
        stack = append(stack, curr)
    }

    return days
}
