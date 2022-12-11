import unittest
from main import Solution

class SolutionTest(unittest.TestCase):
    def testClimbStairs(self):
        self.assertEqual(Solution().climbStairs(1), 1)
        self.assertEqual(Solution().climbStairs(3), 3)
        self.assertEqual(Solution().climbStairs(38), 63245986)


if __name__ == '__main__':
    unittest.main()
        