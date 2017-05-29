import java.util.HashMap;
import java.util.Arrays;
/*
 * Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
 *
 * The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.
 *
 * You may assume that each input would have exactly one solution and you may not use the same element twice.
 *
 * Input: numbers={1, 2, 7, 11, 15}, target=9
 * Output: index1=1, index2=2
 *
 * }
 */
public class TwoSumII {
    public static int[] twoSum(int[] arr, int target) {
        int[] result = new int[2];
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();
        for(int i=0;i<arr.length;i++){
            int diff = target - arr[i];
            if(map.containsKey(arr[i])){
                result[0] = map.get(arr[i]);
                result[1] = i;
            }
            map.put(diff, i);
        }
        return result;
    }
    public static void main(String[] args){
        int[] arr =  {1, 2, 3, 4, 5, 6};
        int target = 9;
        int[] result = twoSum(arr, target);
        System.out.println(Arrays.toString(result));
    }
}
