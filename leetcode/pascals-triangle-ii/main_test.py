import unittest
from main import Solution

class SolutionTest(unittest.TestCase):
    def testGenerate(self):
        self.assertEqual(Solution().getRow(0), [1])
        self.assertEqual(Solution().getRow(4), [1,4,6,4,1])


if __name__ == '__main__':
    unittest.main()
        