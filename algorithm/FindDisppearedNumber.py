#!/usr/bin/python
'''
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
'''

def getDN(arr):
    aSet = set(arr)
    rt = []
    for i in range(1, len(arr)+1):
        if i not in aSet:
            rt.append(i)
    return rt

arr = [1, 3, 3, 5, 6, 2]

print getDN(arr)
