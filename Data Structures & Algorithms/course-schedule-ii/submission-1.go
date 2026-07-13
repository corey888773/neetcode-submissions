func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
		indegree[p[0]]++
	}

	queue := []int{}
	for idx, degree := range indegree {
		if degree == 0 { queue = append(queue, idx) }
	}

	order := []int{}
	for len(queue) > 0 {
		class := queue[0]; queue = queue[1:]
		order = append(order, class)

		for _, dep := range adj[class] {
			indegree[dep]--
			if indegree[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	if len(order) != numCourses {
		return []int{}
	}

	return order
}
