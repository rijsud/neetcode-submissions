class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        indices = {}
        for i in range(len(nums)):
            difference = target - nums[i]
            if difference in indices:
                return [indices[difference], i]
            indices[nums[i]] = i