import unittest
from main import ListNode, Solution
from typing import Union

def makeList(arr) -> Union[ListNode,None]:
    if len(arr) > 0:
        return ListNode(arr[0], makeList(arr[1:]))
    return None

class SolutionTest(unittest.TestCase):
    def testMakeList(self):
        lst = makeList([0,1,2,3,4])
        self.assertEqual(5, Solution().len(lst))
        for i in range(5):
            self.assertEqual(i, lst.val)
            lst = lst.next
        self.assertEqual(None, lst)

    def testEqual(self):
        self.assertEqual(makeList([0,1,2,3]), makeList([0,1,2,3]))
        self.assertNotEqual(makeList([1,2,5]), makeList([1,2,6,7]))

    def testRemoveIth(self):
        lst = makeList([0,1,2,3,4])
        self.assertEqual(makeList([0,1,2,4]), Solution().removeIth(makeList([0,1,2,3,4]), 3))

    def testDeleteMiddle(self):
        self.assertEqual(makeList([0,1,3,4]), Solution().deleteMiddle(makeList([0,1,2,3,4])))
        self.assertEqual(None, Solution().deleteMiddle(makeList([0])))
        self.assertEqual(makeList([0,1,3]), Solution().deleteMiddle(makeList([0,1,2,3])))

if __name__ == '__main__':
    unittest.main()
        