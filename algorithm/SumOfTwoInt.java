/*
* 不使用 + - 运算符实现求两个数的和
*/
public class SumOfTwoInt {
    public static int getSum(int a, int b) {
        while (b != 0) {
            int c = a ^ b;
            b = (a & b) << 1;
            a = c;
        }
        return a;
    }


    public static void main(String[] args){
        System.out.println(getSum(1, 10));
    }
}
