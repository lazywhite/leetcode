import java.util.Arrays;
/*
 * Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
 *
 * For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].
 *
 * Note:
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 */
public class  MoveZero {
    /*
     * 1. first find a == 0
     * 2. next find b not 0
     * 3. swith  a , b
     */
    public static int[] moveByIndex(int[] arr) {
        for(int i=0;i<arr.length-1;i++){
            if(arr[i] == 0){
                for(int j=i+1;j<arr.length;j++){
                    if(arr[j] != 0){
                        arr[i] = arr[j];
                        arr[j] = 0;
                        break;
                    }
                }
            }
        }
        return arr;
    }
    public static int[] move(int[] arr) {
        for(int i=0;i<arr.length;i++){
            if(arr[i] == 0){
                for(int j=i;j<arr.length-1;j++){
                    arr[j] = arr[j+1]; 
                }
                arr[arr.length-1] = 0;
            }

        }
        return arr;
    }
    public static void main(String[] args){
        int[] nums = { 1, 0, 1, 0, 3, 12 };
//        int[] nums = { 1, 1, 1, 1, 3, 12 };
        //        move(nums);
        moveByIndex(nums);
        System.out.println(Arrays.toString(nums));
    }
}
