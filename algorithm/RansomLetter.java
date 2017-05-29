/*
 * RansomLetter.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.
 *
 * Each letter in the magazine string can only be used once in your ransom note.
 *
 * Note:
 * You may assume that both strings contain only lowercase letters.
 *
 * canConstruct("a", "b") -> false
 * canConstruct("aa", "ab") -> false
 * canConstruct("aa", "aab") -> true
 *
 */
import java.util.HashMap;
import java.util.Iterator;

public class RansomLetter
{
	public static boolean canConstruct(String msg, String magazine) {
        HashMap<Character, Integer> msgMap = new HashMap<Character, Integer>();
        HashMap<Character, Integer> mgzMap = new HashMap<Character, Integer>();

        for(int i=0;i<magazine.length();i++){
            Character c = new Character(magazine.charAt(i));
            if(mgzMap.containsKey(c)){
                mgzMap.put(c, mgzMap.get(c)+1);
            }else{
                mgzMap.put(c, 1);
            }
        }

        for(int i=0;i<msg.length();i++){
            Character c = new Character(msg.charAt(i));
            if(mgzMap.get(c) - 1 < 0){
                return false;
            }else{
                mgzMap.put(c, mgzMap.get(c) - 1);
            }
        }
        
        return true; 
		
	}
    public static void main(String[] args){
        String msg = "aa";
        String magazine = "abcda";
        System.out.println(canConstruct(msg, magazine));
    }
}

