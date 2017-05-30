import java.util.Arrays;

/*
 * Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.
 * You may assume the integer do not contain any leading zero, except the number 0 itself.
 * The digits are stored such that the most significant digit is at the head of the list.
 *
 * ie:  [1, 6, 9, 8] --> 1699;
 */
public class PlusOne{
    public static int[] plusOne(int[] digits) {
        int carry = 1;
        for(int i=digits.length-1;i>=0;i--){
            int num = digits[i];
            int rt = num + carry;
            int left = rt % 10;
            carry = rt / 10;
            digits[i] = left;
            if(carry == 1){
                continue;
            }else{
                break;
            }
        }
        //循环结束， 要么中途已经没有进位了, 要么循环耗尽
        if(carry == 1){//加到最后还有进位， 说明结果一定是1和n个0
            int len = digits.length;
            int[] newArray = new int[len + 1];
            Arrays.fill(newArray, 0);
            newArray[0] = 1;
            return newArray;
        }else{
            return digits;
        }
    }

    public static void main(String[] args){
        int[] digits = {1, 3, 4, 7, 9, 9,9,9,9,9,9, 9,9,9,9};
        System.out.println(Arrays.toString(plusOne(digits)));
        int[] newdigits = {9,9,9,9,9,9,9, 9,9,9,9};
        System.out.println(Arrays.toString(plusOne(newdigits)));
        int[] arr = {9,9,9,9,1,9,9, 9,9,9,9};
        System.out.println(Arrays.toString(plusOne(arr)));
    }
}
