/*
 * PowerOfTwo.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 *
 * Given an integer, write a function to determine if it is a power of two.
 *
 * hamming weight 是否为1
 *
 *
 */

public class PowerOfTwo {
	public static boolean checkPowerOfTwo(int num) {
        if (num % 2 != 0 || num <= 0){
            return false;
        }
        while(num % 2 == 0){
            num /= 2;
        }
	    return num == 1;	
	}
    public static void main(String[] args){
        System.out.println(checkPowerOfTwo(12));
        System.out.println(checkPowerOfTwo(256));
        System.out.println(checkPowerOfTwo(0));
    }
}

