func foreignDictionary(words []string) string {
    inWords := map[byte]bool{}
    for _, word := range words {
        for i := 0; i < len(word); i++ {
            inWords[word[i]] = true
        }
    }

    adj := map[byte]map[byte]bool{}
    indegree := map[byte]int{}
    for c := range inWords {
        adj[c] = map[byte]bool{}
        indegree[c] = 0
    }

    for i := 0; i < len(words)-1; i++ {
        a, b := words[i], words[i+1]
        if len(a) > len(b) && strings.HasPrefix(a, b) {
            return ""
        }
        for k := 0; k < min(len(a), len(b)); k++ {
            if a[k] != b[k] {
                u, v := a[k], b[k]
                if !adj[u][v] {
                    adj[u][v] = true
                    indegree[v]++
                }
                break
            }
        }
    }

    // toposort
    queue := []byte{}
    for c := range inWords {
        if indegree[c] == 0 {
            queue = append(queue, c)
        }
    }

    var sb strings.Builder
    for len(queue) > 0 {
        c := queue[0]; queue = queue[1:]
        sb.WriteByte(c)
        for neigh := range adj[c] {
            indegree[neigh]--
            if indegree[neigh] == 0 {
                queue = append(queue, neigh)
            }
        }
    }

    if sb.Len() != len(inWords) {
        return "" // cykl
    }

    return sb.String()
}