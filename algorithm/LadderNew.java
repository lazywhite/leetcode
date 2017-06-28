/*
 * Ladder.java
 * Copyright (C) 2017 white <white@localhost>
 *
 * Distributed under terms of the MIT license.
 *
 * 有一个10级的楼梯, 每次只能上一级, 两级或三级, 问一共有多少种走法
 */
import java.util.*;

public class LadderNew {
     
    public static void main(String[] args){
        dp();
    }
    
    /*动态规划
     * 1. 最优子结构
     *      f(10) = f(9) + f(8) + f(7)
     *
     * 2. 边界
     *      f(1) = 1
     *      f(2) = 2
     *      f(3) = [1,1,1; 1,2; 2,1; 3] = 4
     *
     * 3. 状态转换方程
     *     f(n) = f(n-1) + f(n-2) + f(n-3)
     * */
    public static void dp(){
        int a = 1;
        int b = 2;
        int c = 4;
        for(int i=4;i<=10;i++){
            int temp = a + b + c;
            a = b;
            b = c;
            c = temp;
        }
        System.out.println(c);
    }
}


