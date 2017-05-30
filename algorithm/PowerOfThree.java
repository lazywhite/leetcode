/*
 * PowerOfThree.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 *
 * Given an integer, write a function to determine if it is a power of three.
 *
 * Follow up:
 * Could you do it without using any loop / recursion?
 *
 *
 * 3^x=n
 *
 * log(3^x) = log(n)
 *
 * x log(3) = log(n)
 *
 * x = log(n) / log(3)
 */


public class PowerOfThree {
	public static boolean check(int num) {
        return 1162261467 % num == 0; //Integer范围内3的最大次方数1162261467
        // 若能被num整除， 则num一定是3的次方数
    }
	public static boolean checkPowerOfThree(int num) {
        double res = Math.log(num) / Math.log(3);
        return Math.abs(res - Math.rint(res)) < 0.0000000001;
	}
    public static void main(String[] args){
        System.out.println(checkPowerOfThree(10));
        System.out.println(check(9));
    }
}

