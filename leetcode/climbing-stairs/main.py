class Solution:

    memoize : dict = {}

    def climbStairs(self, n: int) -> int:
        memoized = self.memoize.get(n)
        if memoized != None:
            return memoized
        if n <= 1:
            return 1
        sub1 =  self.climbStairs(n-1)
        self.memoize[n-1] = sub1 
        sub2 = self.climbStairs(n - 2)
        self.memoize[n-2] = sub2
        return sub1 + sub2
