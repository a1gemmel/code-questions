from typing import Dict, List

class Solution:

    # memoize[i, n] contains the least number of coins needed to make n with the first i coin types 
    memoize : Dict[int, Dict[int, int]] = {}

    def minIfPossible(self, a, b):
        if a == -1:
            return b
        if b == -1:
            return a
        return min(a,b)

    def coinChangeInner(self, coins, firstN, amount, ) -> int:
        #print(coins, firstN, amount)
        if amount < 0:
            return -1
        if amount == 0:
            return 0
        if firstN == 0:
            return -1
        if self.memoize[firstN].get(amount) != None:
            return self.memoize[firstN][amount]

        withNewCoin = self.coinChangeInner(coins, firstN, amount - coins[firstN - 1])
        if withNewCoin != -1:
            withNewCoin +=1 
        withoutNewCoin = self.coinChangeInner(coins, firstN - 1, amount)
        val = self.minIfPossible(withNewCoin, withoutNewCoin)
        self.memoize[firstN][amount] = val
        return val

    def coinChange(self, coins: List[int], amount: int) -> int:
        for i in range(1, len(coins) + 1):
            self.memoize[i] = {}

        return self.coinChangeInner(coins, len(coins), amount)