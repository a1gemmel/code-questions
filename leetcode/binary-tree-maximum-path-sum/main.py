from typing import Optional

import math

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def maxPath(self, root: Optional[TreeNode]) -> int:
        if root == None:
            return 0

        return max(0, root.val + max(0, self.maxPath(root.left), self.maxPath(root.right)))

    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        if root == None:
            return 0
        return max(
            self.maxPathSum(root.left) if root.left != None else -1001, 
            self.maxPathSum(root.right) if root.right != None else -1001,
            root.val + self.maxPath(root.left) + self.maxPath(root.right),
            root.val)