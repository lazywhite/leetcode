/*
 * PowerOfTwo.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 *
 * Given an integer, write a function to determine if it is a power of two.
 *
 *
 *
 */

public class PowerOfTwo {
    /* hamming weight 是否为1*/
    public static boolean check(int num) {
        return (num & (num -1)) == 0;
    }


    public static boolean checkPowerOfTwo(int num) {
        /* too slow */
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
        System.out.println(check(256));
    }

}

