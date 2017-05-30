/*
 * UglyNumber.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 *
 * Write a program to check whether a given number is an ugly number.
 *
 * Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For example, 6, 8 are ugly while 14 is not ugly since it includes another prime factor 7.
 *
 * Note that 1 is typically treated as an ugly number.
 *
 *
 */

public class UglyNumber {
    public static boolean check(int num){
        while(num % 2 == 0){
            num /= 2;
        }
        while(num % 3 == 0){
            num /= 3;
        }
        while(num % 5 == 0){
            num /= 5;
        }

        return 1 == num; 

    }
    public static void main(String[] args){
        System.out.println(check(14));
    }
}

