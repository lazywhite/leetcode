/*
 * Find the contiguous subarray within an array (containing at least one number) which has the largest sum.
 *
 * For example, given the array [-2,1,-3,4,-1,2,1,-5,4], 
 * the contiguous subarray [4,-1,2,1] has the largest sum = 6.]])
 */

public class MaxSubArray{
    public static int maxSubArray(int[] arr) {
        //递归求解
        int len = arr.length;
        int[] dp = new int[len];
        dp[0] = arr[0];
        for(int i=1;i<len;i++){
            dp[i] = Math.max(dp[i-1]+ arr[i], arr[i]); //如果加了当前数, 结果却比当前数小, 说明之前的连续数组的和为负数, 则全部抛弃之前的数组
        }
        int rt = dp[0];
        for(int i=0;i<len;i++){
            rt = Math.max(rt, dp[i]);                            
        }
        return rt;
    }
    public static void dp(int[] arr){
        /*
         * 1. 最优子结构
         *      f(2) = Math.max(f(1) + arr[1], arr[1])
         * 2. 边界
         *      f(0) =  -2
         *      f(1) = Math.max(-2 + 1, 1)
         * 3. 状态转换方程
         *      f(n) = Math.max(f(n-1) + arr[n], arr[n])
         */
        int global = arr[0];
        int local = arr[0];
        for(int i=1;i<arr.length;i++){
            local = Math.max(arr[i] + local, arr[i]);
            global = local>global?local:global;
        }
        System.out.println(global);
    }
    public static void main(String[] args){
        int[] arr = {-2,1,-3,4,-1,2,1,-5,4};
        System.out.println(maxSubArray(arr));
        dp(arr);
    }
}
