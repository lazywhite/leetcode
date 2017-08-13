/*
 * Copyright (C) 2017 white <white@Whites-Mac-Air.local>
 *
 * Distributed under terms of the MIT license.
 *
 * 沙滩上有一堆桃子, 猴子A吃掉一个后, 发现剩下的桃子刚好能平均分为5分, 于是拿走了一份; 猴子B吃掉剩下中的一个后, 发现剩下的桃子也刚好能分为5份, 于是也拿走一份, 一直到猴子E都进行了同样的操作, 问沙滩上原来最少有多少只桃子?
 * 
 */
import java.util.Scanner;

public class ADF {
     
	public MonkeyAndPeach() {
	}
    public static void main(String[] args){

        Scanner sc = new Scanner(System.in);
        System.out.println("输入猴子数目");
        int monkeyNum = sc.nextInt();
        int minAmount = 0;

        boolean foundMinAmount = false;
        while(!foundMinAmount){

            int i = -1;
            int amount = minAmount;

            for(i=0;i<monkeyNum;i++){
                if(amount % monkeyNum == 1){
                    int k = (amount - 1) / monkeyNum;
                    amount = amount - 1 - k;
                }else{
                    break;
                }
            }

            if(i == monkeyNum){
                foundMinAmount = true;
                System.out.println(minAmount); 
            }else{
                minAmount++;
            }
        }

    }
}

