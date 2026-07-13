const (
	WHITE = 0
	GREY = 1
	BLACK = 2
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
	}

	visited := make([]int, numCourses)
	order := []int{}

	var dfs func (idx int) bool
	dfs = func (idx int) bool {
		visited[idx] = GREY

		for _, prereq := range adj[idx] {
			if visited[prereq] == GREY {
				return true // cycle detected
			}

			if visited[prereq] == WHITE && dfs(prereq) {
				return true
			}
		}

		visited[idx] = BLACK
		order = append(order, idx)
		return false
	}

	for i := range numCourses {
		if visited[i] == WHITE && dfs(i) {
			return []int{}
		}
	}

	return order
}
