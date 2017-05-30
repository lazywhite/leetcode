#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2017 white <white@Whites-Mac-Air.lan>
#
# Distributed under terms of the MIT license.

"""
Reverse a singly linked list.
"""

class ListNode:  
    def __init__(self,x):  
        self.val=x
        self.next=None
  
      
class LinkedList:      
    def __init__(self, head):
        self.head = head

    def reverse(self, head): 
        if head.next is None:  #如果是最后一个节点
            return head
        if head.next.next is None:# 如果是倒数第二个节点
            head.next.next = head
            return head.next
        else:  
            newHead=self.reverse(head.next) # 递归反转下一个子链表
            head.next.next=head
            head.next=None
        self.head = newHead
        return newHead

    def printLinkedList(self):
        node = self.head
        while node:
            print node.val
            node = node.next
      
head=ListNode(1)
p1=ListNode(2)
p2=ListNode(3)
p3=ListNode(4)
head.next=p1
p1.next=p2
p2.next=p3

l1 = LinkedList(head)
l1.printLinkedList()
l1.reverse(l1.head)
l1.printLinkedList()
