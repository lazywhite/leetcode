/*
 * The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
 *
 * Given two integers x and y, calculate the Hamming distance.
 *
 * Note:
 * 0 ≤ x, y < 231.
 *
 * Input: x = 1, y = 4
 * Output: 2
 *
 * Explanation:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
 *        ↑  ↑
 */
public class HammingDistance{
    public static int getHamm(int x, int y){
        String s = Integer.toBinaryString(x ^ y);
        int count = 0;
        byte[] a = s.getBytes();
        for(int i=0;i<a.length;i++){
            if(a[i] == 49){
                count += 1;
            }
        }
        return count;

    }

    public static void main(String[] args){
        System.out.println(getHamm(8, 6));
    }
}
