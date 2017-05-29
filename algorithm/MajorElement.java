/*
 * MajorElement.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 *
 *
 * Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
 * You may assume that the array is non-empty and the majority element always exist in the array.
 *
 *
 * 数目超过数组长度一般的元素有且仅有一个
 *
 * 摩尔投票算法， 每次删除一对不同的元素， 最后剩下的一个元素即为出现最多的元素
 */

public class MajorElement
{
	public static int getMajorElement(int[] arr) {
        int major  = arr[0];
        int count = 1;

        for(int i=1;i<arr.length;i++){
            if(count == 0){
                major = arr[i];
                count++;
            }else if (major == arr[i]){
                count++;
            }else{
                count--;
            }
        }
//        System.out.println(count);
        return major;
	}
    public static void main(String[] args){
	    int[] arr = {1, 2, 4, 5, 2, 2, 3, 2, 2};
        System.out.println(getMajorElement(arr));
    }
}

