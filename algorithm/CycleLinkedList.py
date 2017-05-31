#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2017 white <white@Whites-Mac-Air.local>
#
# Distributed under terms of the MIT license.

"""
given an linked list, check if it has a circle
"""

class ListNode(object):
    def __init__(self, val):
        self.val = val
        self.next = None;

    def setNext(self, node):
        self.next = node

    def getNext(self):
        return self.next


n1 = ListNode(1)
n2 = ListNode(2)
n3 = ListNode(3)
n4 = ListNode(4)

n1.setNext(n2)
n2.setNext(n3)
n3.setNext(n4)
#n4.setNext(n1)


a = []

def checkCircle(head):
    node = head
    while node.next:
        if node not in a:
            a.append(node)
        else:
            return True
            break
        node = node.next
    return False


print checkCircle(n1)
