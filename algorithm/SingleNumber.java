/*
 * Given an array of integers, every element appears twice except for one. Find that single one.
 * Note:
 * Your algorithm should have a linear runtime complexity. > Could you implement it without using extra memory?
 *
 *
 * A XOR A = 0
 * A XOR B XOR C = A XOR C XOR B
   */

public class SingleNumber{
    public static void main(String[] args){
        int[] arr = {1, 1, 3, 3, 5};
        System.out.println(getSingleNum(arr));
    }
    public static int getSingleNum(int[] arr){
        int result = 0; 
        for(int i=0;i<arr.length;i++){
            result ^= arr[i];
        }
        return result;
    }
}
