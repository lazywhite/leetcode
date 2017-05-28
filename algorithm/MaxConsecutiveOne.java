/*
 * Given a binary array, find the maximum number of consecutive 1s in this array.
 *
 * Example 1:
 * Input: [1,1,0,1,1,1]
 * Output: 3
 * Explanation: The first two digits or the last three digits are consecutive 1s.
 *     The maximum number of consecutive 1s is 3.
 */
public class MaxConsecutiveOne {
    public static int getMaxConOne(int[] arr) {
        int local = 0;
        int global = 0;
        for(int i=0;i<arr.length;i++){
            if(arr[i] == 1){
               local++; 
               global = global>local?global:local;
            }else{
                local = 0;
            }
        }
        return global;
    }
    public static void main(String[] args){
        int[] arr = {1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1};
        System.out.println(getMaxConOne(arr));

    }
}
