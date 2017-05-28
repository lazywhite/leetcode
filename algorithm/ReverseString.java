/*
 * Write a function that takes a string as input and returns the string reversed.
 *
 * Example:
 * Given s = "hello", return "olleh".
 */
public class ReverseString {
    public static String reverse(String str) {
        StringBuilder strb = new StringBuilder();
        strb.append(str);
        strb.reverse();
        return strb.toString();
    }
    public static void main(String[] args){
        System.out.println(reverse("hello world"));
    }
}
