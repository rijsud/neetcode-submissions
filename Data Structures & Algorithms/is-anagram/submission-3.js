class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
        if (s.length !== t.length) {return false}
        const array = new Array(26).fill(0)
        for (let i = 0; i < s.length; i++) {
            array[s.charCodeAt(i)-'a'.charCodeAt(0)]++
            array[t.charCodeAt(i)-'a'.charCodeAt(0)]--
        }

        return array.every((val) => val === 0)
    }
}
