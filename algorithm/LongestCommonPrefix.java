public class LongestCommonPrefix {
    public static String getLongestCommonPrefix(String[] strs) {
        if (strs == null) {
            return null;
        }

        if (strs.length == 0) {
            return "";
        }

        String s = strs[0];
        int len = s.length();

        for (int i = 0; i < len; i++) {
            char c = s.charAt(i);

            for (int j = 1; j < strs.length; j++) {
                if (strs[j].length() <= i || c != strs[j].charAt(i)) {
                    return s.substring(0, i);
                }
            }
        }

        return s;
    }
    public static void main(String[] args){
        String[] strs = {"abcd", "abc", "ab"};
        System.out.println(getLongestCommonPrefix(strs));
    }
}
