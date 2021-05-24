# -*- coding: utf8 -*-
'''
前序遍历二叉树
反转二叉树
'''
class BTree:
    def __init__(self, root_node):
        self.root_node = root_node

    def forward_traversal_tree(self):
        self.root_node.forward_traversal()

    def backward_traversal_tree(self):
        self.root_node.backward_traversal()
        
    def middle_traversal_tree(self):
        self.root_node.middle_traversal()
        
    def reverse(self):
        self.root_node.reverse()

class BTreeNode:
    def __init__(self, left, right, value):
        self.value = value
        self.left = left
        self.right = right
    def reverse(self):
        if self.left is not None:
            self.left.reverse()
        if self.right is not None:
            self.right.reverse()
        self.left, self.right = self.right, self.left
    def forward_traversal(self):
        print(self.value, end=",")
        if self.left is not None:
            self.left.forward_traversal()
        if self.right is not None:
            self.right.forward_traversal()
    def backward_traversal(self):
        if self.left is not None:
            self.left.backward_traversal()
        if self.right is not None:
            self.right.backward_traversal()
        print(self.value, end=",")

    def middle_traversal(self):
        if self.left is not None:
            self.left.middle_traversal()
        print(self.value, end=",")
        if self.right is not None:
            self.right.middle_traversal()


node4 = BTreeNode(None, None, 4)
node5 = BTreeNode(None, None, 5)
node6 = BTreeNode(None, None, 6)
node7 = BTreeNode(None, None, 7)


node2 = BTreeNode(node4, node5, 2)
node3 = BTreeNode(node6, node7, 3)

node1 = BTreeNode(node2, node3, 1)


btree = BTree(node1)
print("前序遍历")
btree.forward_traversal_tree()
print()
print("中序遍历")
btree.middle_traversal_tree()
print()
print("后序遍历")
btree.backward_traversal_tree()
print()

btree.reverse()
print("反转二叉树")
btree.forward_traversal_tree()
print()
