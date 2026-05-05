class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        res = {}
        for number in nums:
            if res.get(number):
                return True
            res[number] = 1
        return False