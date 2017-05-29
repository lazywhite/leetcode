/*
 * MissingNumber.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

 * For example,
 * Given nums = [0, 1, 3] return 2.
 */

public class MissingNumber
{
	public static int blackMagic(int[] arr) {
        int res = 0;
        for(int i=0;i<arr.length-1;i++){
            res ^= (i+1) ^ arr[i];
        }
        return res;
	}
    public static void main(String[] args){
        int[] arr = {1, 3, 4, 0, 5};

        System.out.println(blackMagic(arr));
		
    }
}

