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

public class PowerOfFour {
    /* hamming weight 是否为1*/
    /* 4 的次方数， 肯定也是2的次方数*/
    public static boolean check(int num) {
        if(num < 4){
            return false;
        }else{

            return (num & (num -1)) == 0;
        }
    }


    public static void main(String[] args){
        System.out.println(check(256));
        System.out.println(check(4));
    }

}

