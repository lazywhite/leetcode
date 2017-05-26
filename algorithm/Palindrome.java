import java.util.HashMap;
import java.util.Arrays;
/*
 * Determine whether an integer is a palindrome. Do this without enumtra space.
 * can't use string reverse, that require enumtra space
 */
public class Palindrome {
    public static boolean isPalindrome(int num){
        if (num < 0){
            return false;
        }
        if(num == 0){
            return true;
        }
        int base = 1;
        while(num/base >= 10){
            base *= 10;
        }
        while(num>0){
            int leftDigit = num / base;
            int rightDigit = num % 10;
            if (leftDigit != rightDigit){
                return false;
            }else{
                num -= base * leftDigit;
                base /= 100;
                num /= 10;
            }
        }
        return true;
    }
    public static void main(String[] args){
        System.out.println(isPalindrome(10100));
        System.out.println(isPalindrome(12321));
    }
}
