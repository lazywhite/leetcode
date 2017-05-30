import java.util.Arrays;

/*
 * HouseRobber.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 *
 *
 * You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.
 *
 *
 * 房子顺序不可变, 不能抢劫相邻的房子
 */

public class HouseRobber {

    public static int rob(int[] houses){
        if(houses.length == 0){
            return 0;
        }
        int[] dp = new int[houses.length + 1];
        dp[0] = 0;
        dp[1] = houses[0];

        for(int i=2;i<=houses.length;i++){
            dp[i] = Math.max(dp[i-1], dp[i-2] + houses[i-1]);
        }
        return dp[houses.length];
    }
    public static void main(String[] args){
        int[] houses = {1, 3, 8, 1, 4, 2};
        System.out.println(rob(houses));
    }
}

