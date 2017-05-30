/*
 * HammingWeight.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 *
 *
 *
 * Write a function that takes an unsigned integer and returns the number of ’1' bits it has (also known as the Hamming weight).
 *
 * For example, the 32-bit integer ’11' has binary representation 00000000000000000000000000001011, so the function should return 3.
 */


/*
 * def getHammingWeight(num):
 *     return bin(num)[2:].count('1')
 * 
 * print getHammingWeight(8) # --> 1
 */
public class HammingWeight {
	public static int getHammingWeight(int num) {
        int re = 0;
        while(num != 0){
            num = num & ( num - 1);
            re++;
        }
	    return re;	
	}
    public static void main(String[] args){
        System.out.println( getHammingWeight(8)); 
    }
}

