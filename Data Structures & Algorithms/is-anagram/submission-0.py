class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        char_s = {}
        char_t = {}
        for i in range(len(s)):
            char_s[s[i]] = 1 + char_s.get(s[i], 0)
            char_t[t[i]] = 1 + char_t.get(t[i], 0)
        for i in char_s:
            if char_s[i] != char_t.get(i, 0):
                return False
        return True
