/*
 *  给定一个int类型数字， 返回反转次序后的数字
 *  in(123) --> out(321)
 *  in(-123) --> out(-321)
 *  超出int范围的， 返回0
 */
public class ReverseInteger {
    public static int reverse(int x) {
        int temp = Math.abs(x);
        int result;
        String str = Integer.toString(temp);
        StringBuffer sb = new StringBuffer(str);
        str = sb.reverse().toString();
        if(Long.parseLong(str) > Integer.MAX_VALUE){
            result = 0;
        }
        result = x > 0? Integer.parseInt(str):-Integer.parseInt(str);
        return result;
    }
    public static void main(String[] args){
        System.out.println(reverse(134));
    }
}
