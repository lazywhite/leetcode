
public class  DetectCapital {
    public static boolean checkCapital(String str) {
        int count = 0;
        char[] cArr = str.toCharArray();
        for(int i=0;i<cArr.length;i++){
            if(Character.isUpperCase(cArr[i])){
                count++;
            }
        }
        if(cArr.length == count || count == 0 || count == 1 && Character.isUpperCase(cArr[0])){
            return true;
        }else{
            return false;
        }
    }
    public static void main(String[] args){
        String str = "HeLlo";
        System.out.println(checkCapital(str));
    }
}
