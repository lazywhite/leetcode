/*
 * Matrix.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.local>
 *
 * Distributed under terms of the MIT license.
 *顺时针打印矩阵

 *    1 2 3
 *    4 5 6
 *    7 8 9
 *    >>1 2 3 6 9 8 7 4 5 
 *
 *    1  2  3  4
 *    5  6  7  8 
 *    9  10 11 12
 *    13 14 15 16
 *    >> 1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10
 *    
     x阶矩阵, 则右下角元素为 arr[x-1][x-1]
     一次打印一圈, 则起始位置永远是 arr[i][i], 右下角位置是arr[x-1-i][x-1-i]
      0 <= i <= (x-1) / 2
 */

public class MatrixClockwisePrint {
    public static  int[][] arr = {{1, 2, 3, 4}, {5, 6, 7, 8},{9, 10, 11, 12}, {13, 14, 15, 16} };
    public static void print(int n, int m){
        if ( n * 2 == m ){
            System.out.println(arr[n][n] + " ");
        }else{
            for(int i=n; i<=m-n;i++){
                System.out.print(arr[n][i] + " ");
            }
            for(int i=n+1; i<=m-n;i++){
                System.out.print(arr[i][m-n] + " ");
            }
            for(int i=m-n-1; i>=n;i--){
                System.out.print(arr[m-n][i] + " ");
            }
            for(int i=m-n-1; i>n;i--){
                System.out.print(arr[i][n] + " ");
            }

        }

    }    
    public static void main(String[] args){
        int m = 3;
        for(int i=0;i<=m/2;i++){
            print(i, m);
        }

    }
}

