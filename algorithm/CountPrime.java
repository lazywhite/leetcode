/*
 * CountPrime.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.local>
 *
 * Distributed under terms of the MIT license.
 *
 * 埃拉托斯特尼筛法Sieve of Eratosthenes
 */

import java.util.Arrays;
public class CountPrime {
    public static int count(int n){
        boolean[] prime = new boolean[n + 1];
        Arrays.fill(prime, true);
        prime[0] = false;
        prime[1] = false;
         
        int limit = (int)Math.sqrt(n);

        for(int i=2;i<=limit;i++){
            if(prime[i]){
                for(int j=i*i;j<n;j+=i){
                    prime[j] = false;
                }
            }
        }
        int count = 0;
        for(int i=0;i<n;i++){
            if(prime[i]) count++;
        }
        return count;
    }
    public static void main(String[] args){
        System.out.println(count(1000));
    }
}

