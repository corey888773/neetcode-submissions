from collections import deque, defaultdict

class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        indegree = [0] * numCourses
        next_courses = defaultdict(list)

        for preq in prerequisites:
            course, req = preq

            indegree[course] += 1
            next_courses[req].append(course)

        queue = deque([])
        for idx, deg in enumerate(indegree):
            if deg == 0:
                queue.append(idx)

        
        path = []
        while len(queue) > 0:
            course = queue.popleft()
            path.append(course)

            for next_course in next_courses[course]:
                indegree[next_course] -= 1
                if indegree[next_course] == 0:
                    queue.append(next_course)
            
                
        if len(path) != numCourses:
            return []

        return path
