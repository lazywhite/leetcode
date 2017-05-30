/*
 * Find the contiguous subarray within an array (containing at least one number) which has the largest sum.
 *
 * For example, given the array [-2,1,-3,4,-1,2,1,-5,4], 
 * the contiguous subarray [4,-1,2,1] has the largest sum = 6.]])
 */

public class MaxSubArray{
    public static int maxSubArray(int[] arr) {
        int len = arr.length;
        int[] dp = new int[len];
        dp[0] = arr[0];
        for(int i=1;i<len;i++){
            dp[i] = Math.max(dp[i-1]+ arr[i], arr[i]);
        }
        int rt = dp[0];
        for(int i=0;i<len;i++){
            rt = Math.max(rt, dp[i]);                            
        }
        return rt;
    }
    public static void main(String[] args){
        int[] arr = {-2,1,-3,4,-1,2,1,-5,4};
        System.out.println(maxSubArray(arr));
    }
}
