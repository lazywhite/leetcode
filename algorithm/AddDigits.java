/*
 * Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.
 *
 * For example:
 *
 * Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.
 *
 * Follow up:
 * Could you do it without any loop/recursion in O(1) runtime?
 *
 *
 * 输入：1,2,3,4,5,6,7,8,9,10     11, 12, 13, 14, 15, 16， 17， 18， 19，20
 *
 * 结果：1, 2,3,4,5,6,7,8,9,1      2,  3,  4,  5,  6,  7，  8，  9，  1， 2
 *
 */

public class  AddDigits {
    public static int add(int a) {
        return (a-1)%9 + 1;
    }
    public static void main(String[] args){
        System.out.println(add(14010));
    }
}
