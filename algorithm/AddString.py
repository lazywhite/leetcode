#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2017 white <white@Whites-Mac-Air.lan>
#
# Distributed under terms of the MIT license.

"""
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:
    The length of both num1 and num2 is < 5100.
    Both num1 and num2 contains only digits 0-9.
    Both num1 and num2 does not contain any leading zero.
    You must not use any built-in BigInteger library or convert the inputs to integer directly.

"""

str1 = "147580"
str2 = "471235580"

len1 = len(str1)
len2 = len(str2)

length = len1 if len1 > len2 else len2

result = []
carry = 0

for i in range(-1, -length-1, -1):
    try:
        a = str1[i]
    except:
        a = 0
    try:
        b = str2[i]
    except:
        b = 0
    num = int(a)+int(b)+carry
    if num >=10:
        num -= 10
        carry = 1
    else:
        carry = 0
    result.insert(0, num)

print result
