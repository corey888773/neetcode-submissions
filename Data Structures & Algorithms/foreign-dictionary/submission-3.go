func foreignDictionary(words []string) string {
    // unique letters
    uniqueLetters := map[byte]bool{}
    for _, word := range words {
        for i := range len(word) {
            uniqueLetters[word[i]-'a'] = true
        }
    }

    // create the edges
    edges := map[byte]map[byte]bool{}
    var indegrees [26]int

    for i:=1; i < len(words); i++ {
        wordA, wordB := words[i-1], words[i]
        diffrence := false

        for j := range min(len(wordA), len(wordB)) {
            if wordA[j] != wordB[j] {
                bigger, smaller := wordA[j]-'a', wordB[j]-'a'

                if _, ok := edges[bigger]; !ok {
                    edges[bigger] = map[byte]bool{}
                }

                if !edges[bigger][smaller] {
                    indegrees[smaller]++
                }
                edges[bigger][smaller] = true
                diffrence = true
                break
            }
        }

        if !diffrence && len(wordA) > len(wordB) {
            return "" // longer word before its prefix
        }
    }

    queue := []byte{}
    for letter, degree := range indegrees {
        if uniqueLetters[byte(letter)] && degree == 0 {
            queue = append(queue, byte(letter))
        }
    }

    order := []byte{}
    for len(queue) > 0 {
        letter := queue[0]; queue = queue[1:]
        order = append(order, letter + 'a')

        for biggerLetter, _ := range edges[letter] {
            indegrees[biggerLetter]--

            if indegrees[biggerLetter] == 0 {
                queue = append(queue, biggerLetter)
            }
        }
    }

    if len(order) != len(uniqueLetters) {
        return ""
    }

    return string(order)
}