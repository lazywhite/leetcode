/*
 * Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space characters only.
 *
 * For example, 
 * Given s = "Hello World",
 * return 5.
 */
public class LengthOfLastWord{
    public static int getLength(String str) {
        if(str.length() == 0){
            return 0;
        }
        String[] s = str.split(" ");
        return s[s.length-1].length();
    }

    public static void main(String[] args){
        String str = "HellosWorld";
        System.out.println(getLength(str));
        str = "";
        System.out.println(getLength(str));
        str = "Hello World";
        System.out.println(getLength(str));
    }
}
