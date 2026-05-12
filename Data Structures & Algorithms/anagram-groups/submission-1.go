func groupAnagrams(strs []string) [][]string {
    res := make(map[[26]int][]string)
    for _, word := range strs {
        char := [26]int{}
        for _, letter := range word {
            char[letter - 'a']++
        }
        res[char] = append(res[char], word)
    }

    output := [][]string{}
    for _, list := range res {
        output = append(output, list)
    }
    return output
}
