/*
#! /usr/bin/env python
# -*- coding: utf-8 -*-
# vim:fenc=utf-8
#
# Copyright © 2017 white <white@Whites-Mac-Air.lan>
#
# Distributed under terms of the MIT license.

"""

"""
#l1 = [1, 2, 5, 6]
#target = 3
l1 = [1, 2, 2, 5]
target = 5

def getIndex(lst, tgt):
    if tgt < lst[0]:
        return 0
    for index, i in enumerate(l1):
        if i >= target:
            return index + 1
        else:
            continue
    return len(l1)

print getIndex(l1, target) 
*/
public class InsertPosition{
    public static int getIndex(int[]arr, int tgt){
        if(tgt < arr[0]){
            return 0;//边界条件
        }
        for(int i=0;i<arr.length;i++){
            if(arr[i] >= tgt){
                return i+1;
            }else{
                continue;
            }
        }
        return arr.length;//if not found, should be insert to end of array
    }
    public static void main(String[] args){
        int[] arr = { 1, 2, 2, 5 };
        System.out.println(getIndex(arr, 0));
        System.out.println(getIndex(arr, 2));
        System.out.println(getIndex(arr, 6));
    }
}
