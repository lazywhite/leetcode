#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2017 white <white@Whites-Mac-Air.lan>
#
# Distributed under terms of the MIT license.

"""
You are given a string representing an attendance record for a student. The record only contains the following three characters:

    'A' : Absent.
    'L' : Late.
    'P' : Present.
    A student could be rewarded if his attendance record doesn't contain more than one 'A' (absent) or more than two continuous 'L' (late).

    You need to return whether the student could be rewarded according to his attendance record.

Example 1:
    Input: "PPALLP"
    Output: True
Example 2:
    Input: "PPALLL"
    Output: False
"""

record = "PPALLP"

def check(record=None):
    return record.count("A") <= 1 and "LLL" not in record


print check(record)




