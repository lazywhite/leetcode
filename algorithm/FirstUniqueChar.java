/*
 * FirstUniqueChar.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 * Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
 *
 * Examples:
 *
 * s = "leetcode"
 * return 0.
 *
 * s = "loveleetcode",
 * return 2.
 */
import java.util.Arrays;

public class FirstUniqueChar
{
	public static int getFirstUniqueChar(String s) {
        int[] chars = new int[26];
        Arrays.fill(chars, 0);
        for(int i=0;i<s.length();i++){
            chars[s.charAt(i) - 'a'] += 1;
        }
        for(int j=0;j<s.length();j++){
            if(chars[s.charAt(j) - 'a'] == 1){
                return j;
            }

        }
        return -1;
	}
    public static void main(String[] args){
        System.out.println(getFirstUniqueChar("leetcode"));
        System.out.println(getFirstUniqueChar("loveleetcode"));
		
    }
}

