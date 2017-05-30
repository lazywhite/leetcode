import java.util.Arrays;

/*
 *
Given numRows, generate the first numRows of Pascal's triangle.

For example, given numRows = 5,
Return
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/
public class PascalTriangle{
    public static void main(String[] args){ 
        gen(7);
    }
    public static void gen(int num){ 
        for(int i=1;i<=num;i++){
            int[] arr = new int[i];
            for(int j=1;j<=i;j++){
                arr[j-1] = getByIndex(i, j);
            }
            
            for(int z=0;z<=num-i;z++){
                System.out.print(" ");
            }
            System.out.print(Arrays.toString(arr));
            System.out.println();
        }
    }
    public static int getByIndex(int x, int y){
        if (y == 1 || x == y){
            return 1;
        }else{
            return getByIndex(x -1, y-1) + getByIndex(x-1, y);
        }
    }
}
