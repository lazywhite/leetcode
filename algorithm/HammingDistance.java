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
	public static int getHammingWeight(int num) {
        int re = 0;
        while(num != 0){
            num = num & ( num - 1);
            re++;
        }
	    return re;	
	}

    public static int getHamm(int x, int y){
        int num = x ^ y;
        return getHammingWeight(num);

    }

    public static void main(String[] args){
        System.out.println(getHamm(8, 6)); //--> 3 
    }
}
