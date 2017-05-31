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

 *  自己实现27进制
 */

public class SheetNumber {
    public static int getInt(char c) throws Exception{
        if(c == '0'){
            return 0;
        }
        if(c > 'Z'){
            throw new Exception("unlegal");
        }
        return c - 'A' + 1;
    }

    public static int getResult(String str) {
        int result = 0;
        try{
            for(int i=0;i<str.length();i++){
                int digit = getInt(str.charAt(i));
                result += digit * Math.pow(27, i);
            }
        }catch(Exception e){
            System.out.println("unlegal");
        }
        return result;
    }

    public static void main(String[] args) {
        String s = "Z";
        System.out.println(getResult(s));
    }
}

