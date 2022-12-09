# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
    
    def __eq__(self, other):
        if other == None:
            return False
        return self.val == other.val and self.next == other.next

    def __repr__(self):
        arr = []
        while self != None:
            arr.append(self.val)
            self = self.next
        return repr(arr)

class Solution(object):
    def len(self, head):
        length = 0
        while head != None:
            head = head.next
            length +=1
        return length

    def removeIth(self, head, i : int):
        node = head
        for n in range(i-1):
            node = node.next
        node.next = node.next.next

        return head


    def deleteMiddle(self, head):
        """
        :type head: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        length = self.len(head)
        if length == 1:
            return None

        middle = length // 2
        pos = 0
        node = head
        return self.removeIth(head, middle)
