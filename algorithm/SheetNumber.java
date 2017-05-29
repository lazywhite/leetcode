/*
 * SheetNumber.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 *
    Given a column title as appear in an Excel sheet, return its corresponding column number.

    For example:

        A -> 1
        B -> 2
        C -> 3
        ...
        Z -> 26
        AA -> 27
        AB -> 28 

 *  自己实现26进制
 */

public class SheetNumber
{
	public static int getChar(char c) {
        return (int)c - 64;
	}
	public static String getStr(String str) {
        StringBuilder strb = new StringBuilder();
        for(int i=0;i<str.length();i++){
            strb.append(getChar(str.charAt(i)));
        }
        return strb.toString();
    }
    public static void main(String[] args){
        String s = "AB";
        System.out.println(Integer.valueOf(getStr(s), 26));
		
    }
}

