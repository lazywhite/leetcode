/*
 * TreePathMaxSum.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.local>
 *
 * Distributed under terms of the MIT license.
 *
 * 有下列数字树, 从顶向下, 每次只能走下左或下右, 求到达底部时所有路径中的最大值
 *           7
 *         3   8
 *      8    1    0
 *   2    7     4    4
 * 4    5    2    6   5
 *
 *   dp[i][j] = max(dp[i-1][j-1], dp[i-1][j])+ a[i][j]
 *
 */

import java.util.*;
public class TreePathMaxSum {
     
    public static void main(String[] args){
        int[][] arr = {{7}, {3, 8}, {8, 1, 0}, {2, 7, 4, 4}, {4, 5, 2, 6, 5}};
        int[][] dp = new int[5][5];
        dp[0][0] = arr[0][0];
        for(int i=1;i<arr.length;i++){
            for(int j=0;j<arr[i].length;j++){
                if(j==0){
                    dp[i][j] = dp[i-1][j] + arr[i][j]; 
                }else if (j==i){
                    dp[i][j] = dp[i-1][j-1] + arr[i][j];
                }else{
                    dp[i][j] = Math.max(dp[i-1][j-1], dp[i-1][j]) + arr[i][j];
                }
            }
        }
        int max = -1;
        for(int i=0;i<5;i++){
            for(int j=0;j<5;j++){
                max = Math.max(max, dp[i][j]);
            }
        }
        System.out.println(max);
    }
}

