/*
 * Given a positive integer, output its complement number. The complement strategy is to flip the bits of its binary representation.
 *
 * Note:
 * The given integer is guaranteed to fit within the range of a 32-bit signed integer.
 * You could assume no leading zero bit in the integer’s binary representation.
 *
 * Example:  
 *     Input: 5
 *     Output: 2
 * Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
 */
public class NumberComplement {
    public static int getComplement(int num){
        String str = Integer.toBinaryString(num);
        byte[] arr = str.getBytes();
        StringBuilder strb = new StringBuilder();
        for(int i=0;i<arr.length;i++){
            if(arr[i] == 49){
                strb.append(0);
            }else if(arr[i] == 48){
                strb.append(1);
            }
        }
        String strr = strb.toString();
        return Integer.valueOf(strr, 2);
    }
    public static void main(String[] args){
        System.out.println(getComplement(5));//--> 2
        System.out.println(getComplement(1));//--> 0
        System.out.println(getComplement(6));//--> 1
    }
}
