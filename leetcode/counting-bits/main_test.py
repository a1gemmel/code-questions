import unittest
from main import Solution

class SolutionTest(unittest.TestCase):
    def testGenerate(self):
        self.assertEqual(Solution().countBits(1), [0, 1])
        self.assertEqual(Solution().countBits(2), [0, 1, 1])
        self.assertEqual(Solution().countBits(5), [0,1,1,2,1,2])


if __name__ == '__main__':
    unittest.main()
        