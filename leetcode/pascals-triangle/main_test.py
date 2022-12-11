import unittest
from main import Solution

class SolutionTest(unittest.TestCase):
    def testGenerate(self):
        self.assertEqual(Solution().generate(1), [[1]])
        self.assertEqual(Solution().generate(5), [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]])


if __name__ == '__main__':
    unittest.main()
        