class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        k = 0
        for i in range(len(nums)):
            el = nums[i]
            if el != val:
                nums[k] = el
                k+=1
        
        return k