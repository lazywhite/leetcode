/*
 * MergeSortedArray.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.local>
 *
 * Distributed under terms of the MIT license.
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
 *
 * Note: You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2. The number of elements initialized in nums1 and nums2 are m and n respectively.
 */

import java.util.*;
public class MergeSortedArray {
    public static int[] merge(int[] arr1, int[] arr2){
        int len1 = arr1.length;
        int len2 = arr2.length;
        int len = len1 + len2;
        int[] merged = new int[len1 + len2];
        int i=len1-1, j=len2-1, k=len-1;
        while(i>=0 || j>=0){
            if(i<0){
                merged[k--] = arr2[j--];
            }if(j<0){
                merged[k--] = arr2[i--];
            }else{
                merged[k--] = (arr1[i] > arr2[j])?arr1[i--]:arr2[j--];
            }
        }
        return merged;
    }
    public static void main(String[] args){
        int[] arr1 = {1, 3, 4, 5, 6, 7} ;
        int[] arr2 = {1, 3, 8, 9, 16, 17}; 
        System.out.println(Arrays.toString(merge(arr1, arr2)));
    }
}

