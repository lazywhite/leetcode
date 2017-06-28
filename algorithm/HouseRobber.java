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
 * 房子顺序不可变, 不能抢劫相邻的房子, 求可以抢劫到的最多的金钱
 */

public class HouseRobber {
    public static void dp(int[] houses){
        /*
         * 1. 最优子结构
         *      f(3) = Math.max(f(2), f(1) + house[3])
         * 2. 边界
         *     f(0) = house[0] = 1
         *     f(1) = Math.max(house[0], house[1]) = 3
         *     f(2) = Math.max(3, 1 + house[2]) = 9
         *
         *     初始值  a = 1,  b = 3
         *  3. 状态转换方程
         *      f(n) = Math.max(f(n-1), f(n-2) + house[n])
         *
         *  存储index - 2能抢到最多的钱和index - 1能抢到的最多的钱
         */
        int max_2 = 1; //往前2个能抢到最多的钱
        int max_1 = 3;//往前1个能抢到最多的钱
        for(int i=2;i<houses.length-1;i++){
            int temp = Math.max(max_1, max_2 + houses[i]);//获得当前能抢到最大的值
            max_2 = max_1;//下次迭代, max_2的值应该存放max_1
            max_1 = temp; //下次迭代, max_1的值应该存放当前最大值
        }
        System.out.println(max_1);//目前最大的值存放在max_1
    }
    //递归
    public static int rob(int[] houses){
        if(houses.length == 0){
            return 0;
        }
        int[] dp = new int[houses.length];
        dp[0] = 1;
        dp[1] = 3;

        for(int i=2;i<=houses.length-1;i++){
            dp[i] = Math.max(dp[i-1], dp[i-2] + houses[i]);
        }
        return dp[houses.length -1];
    }
    public static void main(String[] args){
        int[] houses = {1, 3, 8, 1, 4, 2};
//        System.out.println(rob(houses));
        dp(houses);
    }
}

