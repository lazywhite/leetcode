/*
 * CommonMultiple.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.local>
 *
 * Distributed under terms of the MIT license.
 * 两个数的最小公倍数就是 两数的乘积除以两数的最小公倍数
 */

public class CommonMultiple {
     
    public static void main(String[] args){
        int a = 18, b = 5;
        CommonDivisor cd = new CommonDivisor();
        int divisor = cd.getCommonDivisor(a, b);
        System.out.println((a * b) / divisor);
    }
}

