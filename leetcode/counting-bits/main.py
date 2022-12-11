from typing import List
import math

class Solution:
    def countBits(self, n: int) -> List[int]:
        arr : List[int] = [0]
        for i in range(1, n+1): 
            arr.append(1 + arr[i - (2 ** math.floor(math.log2(i)))])
        return arr

