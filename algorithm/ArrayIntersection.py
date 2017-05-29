#!/usr/bin/python
'''
Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

Note:
Each element in the result must be unique.
The result can be in any order.

'''
arr1 = [1, 2, 2, 1]
arr2 = [2, 2]


arr3 = [ i for i in arr1 if i in arr2]
print set(arr3)
