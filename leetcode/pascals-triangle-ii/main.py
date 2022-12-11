from typing import List 

class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        row : List[int] = [1]

        while (rowIndex > 0):
            new_row : list = []
            for c in range(len(row) + 1):
                if c == 0 or c == len(row):
                    new_row.append(1)
                else:
                    new_row.append(row[c] + row[c-1])
            row = new_row
            rowIndex -= 1

        return row