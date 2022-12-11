import unittest
from main import Solution

class SolutionTest(unittest.TestCase):
    def testCoinChange(self):
        self.assertEqual(Solution().coinChange([1,2,5], 11), 3)
        self.assertEqual(Solution().coinChange([2,4], 11), -1)
        self.assertEqual(Solution().coinChange([146,66,355,93,255,152,225,244,168,205], 8069), 25)
if __name__ == '__main__':
    unittest.main()
        