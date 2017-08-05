/*
 * TriangleArea.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.local>
 *
 * Distributed under terms of the MIT license.
 */

import java.util.*;
public class TriangleArea {
    public static int a;
    public static int b;
    public static int c;
     
    public static boolean valid(){
        if( ((a+b) > c ) && ((a+c) > b ) && (b+c) > a){
            return true;
        }
        return false;
    }

    public static double getArea(){
        //海伦公式
        double p = ((double) a + b + c ) / 2;
         return Math.sqrt(p * (p-a) * (p-b) * (p-c));
    }

    public static void main(String[] args){
        Random r = new Random();
        int[] list = new int[3];
        for(int i=0;i<3;i++){
            list[i] = r.nextInt(10);
        }
        a = list[0];
        b = list[1];
        c = list[2];
        System.out.println(Arrays.toString(list));
        if(valid()){
            System.out.println(getArea());
        }
    }
}

