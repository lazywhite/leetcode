/*
 * CommonDivisor.java
 * Copyright (C) 2017 white <white@localhost>
 *
 * Distributed under terms of the MIT license.
 */

public class CommonDivisor {
    /* 更相减损术 */    
	public static int getCommonDivisor02(int a, int b) {
        int temp = -1;
        if(b > a){
            temp = b;
            b = a;
            a = temp;
        }
        int left = a  - b;
        if(left == 0){
            return b;
        }else{
            return getCommonDivisor02(b, left);
        }

	}
    /* 辗转相除法 */
	public static int getCommonDivisor(int a, int b) {
        int temp = -1;
        if(b > a){
            temp = b;
            b = a;
            a = temp;
        }
        int divisor = a  % b;
        if(divisor == 0){
            return b;
        }else{
            return getCommonDivisor(b, divisor);
        }

	}
    public static void main(String[] args){
        System.out.println(getCommonDivisor(18, 5));
        System.out.println(getCommonDivisor02(18, 5));
    }
}

