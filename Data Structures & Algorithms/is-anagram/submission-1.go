func isAnagram(s string, t string) bool {
    if len(s) != len(t) {return false}
    char_s, char_t := make(map[rune]int), make(map[rune]int)
    for _, letter := range s {
        if _, ok := char_s[letter]; !ok {
            char_s[letter] = 0
        }
        char_s[letter] += 1
    }
    for _, letter := range t {
        if _, ok := char_t[letter]; !ok {
            char_t[letter] = 0
        }
        char_t[letter] += 1
    }

    for key, _ := range char_s {
        if _, ok := char_t[key]; !ok{return false}
        if char_s[key] != char_t[key] {return false}
    }
    return true
}
