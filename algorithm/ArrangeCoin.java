/*
 * ArrangeCoin.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.local>
 *
 * Distributed under terms of the MIT license.
 *
 *
 * You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.
 *
 * Given n, find the total number of full staircase rows that can be formed.
 *
 * Example 1:
 *
 * n = 5
 *
 * The coins can form the following rows:
 * ¤
 * ¤ ¤
 * ¤ ¤
 *
 * Because the 3rd row is incomplete, we return 2.
 * Example 2:
 *
 * n = 8
 *
 * The coins can form the following rows:
 * ¤
 * ¤ ¤
 * ¤ ¤ ¤
 * ¤ ¤
 *
 * Because the 4th row is incomplete, we return 3.
 *
 */

public class ArrangeCoin {
    public static int getFloor(int total){
        int i = 1;
        while(total >= 0){
            total -= i;
            i++;
        }
        if(total <0){
            return i-2;
        }else{
            return i-1;
        }
    }
    public static void main(String[] args){
        System.out.println(getFloor(1000));
        System.out.println(getFloor(5));
        System.out.println(getFloor(7));
    }
}

