/*
 * ValidAnagram.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 *
 * Given two strings s and t, write a function to determine if t is an anagram of s.
 *
 * For example,
 * s = "anagram", t = "nagaram", return true.
 * s = "rat", t = "car", return false.
 *
 * Note:
 * You may assume the string contains only lowercase alphabets.
 */

public class ValidAnagram
{
	public static boolean testValidAnagram(String s, String t) {
        if(s.length() != t.length()){
            return false;
        }else{
            int len = s.length();
            for(int i=0;i<len;i++){
                if(s.charAt(i) != t.charAt(len-i-1)){
                    return false;
                }
            }
            return true;
        }
	}
	public static boolean check(String s, String t) {
        StringBuilder strb = new StringBuilder(s);
        if(strb.reverse().toString().equals(t)){
            return true;
        }else{
            return false;
        }
    }
    public static void main(String[] args){
        String s = "anagram";
        String t = "margana";

//        System.out.println(testValidAnagram(s, t));
        System.out.println(check(s, t));
		
    }
}

