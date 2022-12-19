import unittest
from main import ordered

class OrderedTest(unittest.TestCase):
    def test(self):
        self.assertTrue(ordered([[1]], [[[[4,7,0]],[],[[],8],[[7,1]]]]))
        self.assertTrue(ordered([], [1]))
        self.assertTrue(ordered([1], [2]))
        self.assertTrue(ordered([1], [[[[[[[[[[3]]]]]]]]]]))
        self.assertTrue(ordered([[0]], [[0,0,0]]))
        self.assertTrue(ordered([[0, 0]], [[0,1]]))

        self.assertFalse(ordered([1], []))
        self.assertFalse(ordered([[8,9,[8],9]], [[],[3,6],[[],[[],3,5],[],6,[[],[5,4,5],10,6]]]))
        self.assertFalse(ordered([[7,0,7,[4,[7,7],[2,1,9,6],8],10],[],[]],[]))


if __name__ == '__main__':
    unittest.main()
        