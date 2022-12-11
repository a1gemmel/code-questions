from typing import List 

class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        rows : list = [[1]]

        for r in range(1, numRows):
            row : list = []
            for c in range(r + 1):
                if c == 0 or c == r:
                    row.append(1)
                else:
                    row.append(rows[r-1][c] + rows[r-1][c-1])
            rows.append(row)

        return rows