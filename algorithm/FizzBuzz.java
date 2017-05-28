import java.util.ArrayList;
import java.util.Arrays;

public class FizzBuzz {
    public static String getStr(int n) {
        if(n % 3 ==0 && n % 5 == 0){
            return "FizzBuzz";
        }else if(n%3==0){
            return "Fizz";
        }else if(n%5==0){
            return "Buzz";
        }else{
            return Integer.toString(n);
        }
    }
    public static ArrayList<String> getFB(int n) {
        ArrayList<String> strs = new ArrayList<>();
        for(int i=1;i<=n;i++){
            strs.add(getStr(i));
        }
        return strs;
    }
    
    public static void main(String[] args){
        System.out.println(getFB(15));
    }
}
