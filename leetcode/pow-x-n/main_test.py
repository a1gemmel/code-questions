import unittest
from main import Solution

class SolutionTest(unittest.TestCase):
    def testCoinChange(self):
        self.assertEqual(Solution().myPow(0.00001, 2147483647), 0)
        
if __name__ == '__main__':
    unittest.main()
        