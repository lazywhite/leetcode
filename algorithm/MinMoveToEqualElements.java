/*
   Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.

Example:

Input:
[1,2,3]

Output:
3

Explanation:
Only three moves are needed (remember each move increments two elements):

[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]

数组[1，2，3], 给除去最大值的其他数字加1，变为[2，3，3]，我们全体减1，并不影响数字间相对差异，变为[1，2，2]，这个结果其实就是原始数组的最大值3自减1
那么问题转化为"将所有数字都减小到最小值"
*/

import java.util.Collections;
import java.util.Arrays;
import java.util.ArrayList;
import java.util.List;
public class  MinMoveToEqualElements {
    public static int getMoves(int[] arr) {
        List<Integer> intList = toList(arr);
        int min = Collections.min(intList);
        int moves = 0;
        for(int i=0;i<arr.length;i++){
            if(arr[i] > min){
                moves += arr[i] - min;
            } 
        }
        return moves;
    }

    public static void main(String[] args){
        int[] arr  = {1, 2, 3, 4, 5};
        System.out.println(getMoves(arr));
    }
    public static List<Integer> toList(int[] arr){
        List<Integer> intList = new ArrayList<Integer>();
        for (int index = 0; index < arr.length; index++)
        {
            intList.add(arr[index]);
        }
        return intList;
    }
    
    
}
