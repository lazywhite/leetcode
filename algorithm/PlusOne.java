/*
 * Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.
 * You may assume the integer do not contain any leading zero, except the number 0 itself.
 * The digits are stored such that the most significant digit is at the head of the list.
 *
 * ie:  [1, 6, 9, 8] --> 1699;
 */
public class PlusOne{
    public static int plus(int[] digits) {
        StringBuilder strb = new StringBuilder();
        for(int i=0;i<digits.length;i++){
            strb.append(digits[i]);
        }
        String s =  strb.toString();
        Long l = Long.parseLong(s);
        if(l+1 > Integer.MAX_VALUE){
            return 0;
        }
        return Integer.parseInt(s) + 1;
    }

    public static void main(String[] args){
        int[] digits = {1, 3, 4, 7, 9, 9,9,9,9,9,9, 9,9,9,9};
        System.out.println(plus(digits));
        int [] adigits = {1, 3, 7, 9};
        System.out.println(plus(adigits));
    }
}
