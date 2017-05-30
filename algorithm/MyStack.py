#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2017 white <white@Whites-Mac-Air.lan>
#
# Distributed under terms of the MIT license.
'''
    self-implemented stack using queue
'''


class MyQueue(object):
    queue = []               
    def __init__(self):
        """
        Initialize your data structure here.
        """
        pass
        

    def push(self, x):
        """
        Push element x to the back of queue.
        :type x: int
        :rtype: void
        """
        self.queue.append(x)
        return 
        

    def pop(self):
        """
        Removes the element from in front of queue and returns that element.
        :rtype: int
        """
        if self.empty():
            pass
        qlen = len(self.queue)
        return self.queue.remove(self.queue[qlen-1])
        

    def peek(self):
        """
        Get the front element.
        :rtype: int
        """
        if self.empty():
            pass
        else:
            qlen = len(self.queue)
            return self.queue[qlen-1]
        

    def empty(self):
        """
        Returns whether the queue is empty.
        :rtype: bool
        """
        if len(self.queue):
            return False
        return True
        


x = 100
obj = MyQueue()
obj.push(x)
obj.push(x)
print obj.pop()
print obj.peek()
print obj.empty()
