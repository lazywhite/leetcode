/*
 * Ladder.java
 * Copyright (C) 2017 white <white@localhost>
 *
 * Distributed under terms of the MIT license.
 *
 * 有一个10级的楼梯, 每次只能上一级或两级, 问一共有多少种走法
 */
import java.util.*;

public class Ladder {
     
    public static void main(String[] args){
        dp();
//        System.out.println(f(10));
//       System.out.println(fUpdated(10));
    }
    
    /*递归
     * 1. 时间复杂度 0(2^n)
     * */
    public static int f(int n){
        if(n == 1){
            return 1;
        }
        if(n == 2){
            return 2;
        }
        return f(n-1) + f(n-2);
    }

    /*避免重复计算的递归
     * 1. 节省了时间复杂度
     * 2. 增大了空间复杂度
     * */
    public static int fUpdated(int n){
        Map<Integer, Integer> map = new HashMap<>(); 
        if(n == 1){
            return 1;
        }
        if(n == 2){
            return 2;
        }
        if(map.containsKey(n)){
            return map.get(n);
        }
        int value = f(n-1) + f(n-2);
        map.put(n, value);
        return value;
    }
    /*动态规划
     * 从前向后推导
     * 当前状态只与前两个状态有关, 因此只需同时保存3个值即可
     * */
    public static void dp(){
        int a = 1;
        int b = 2;
        int temp = 0;
        for(int i=3;i<=10;i++){
            temp = a + b;
            a = b;
            b = temp;
        }
        System.out.println(temp);
    }
}


