/*
 * SwitchNumInPlace.java
 * Copyright (C) 2017 white <white@localhost>
 *
 * Distributed under terms of the MIT license.
 * 原地交换两个int的值
 *
 *
 * 异或法则
 *      1. 交换律    a ^ b = b ^ a 
 *      2. 结合律    a ^ (b ^ c) == (a ^ b) ^ c
 * 推论
 *      a ^ a = 0 , a ^ 0 = a
 *      所以有  a ^ b ^ b == a ^ 0 == a
 *
 */

public class SwitchNumInPlace {
     
	public static void Switch01() {
        int a = 10;
        int b = 20;
        a = a+b;  //此处容易发生溢出
        b = a - b;
        a = a - b;
        System.out.println(a + " " + b);
	}
	public static void Switch02() {
        int a = 10;
        int b = 20;

        a = a ^ b;
        b = a ^ b;
        a = a ^ b;
        System.out.println(a + " " + b);
	}
    public static void main(String[] args){
        Switch02();
    }
}

