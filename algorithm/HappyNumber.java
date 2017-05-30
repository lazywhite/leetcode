import java.util.HashSet;

/*
 * HappyNumber.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 */

public class HappyNumber {
    public static boolean isHappy(int n) {
        if (n <= 0) {
            return false;
        }
        HashSet<Integer> set = new HashSet<Integer>();
        set.add(n);
        while (n != 1) {
            n = helper(n);
            if (set.contains(n)) {
                return false;
            }
            set.add(n);
        }
        return true;
    }

    private static int helper(int n) {
        int base = 10;
        int ret = 0;
        while (n > 0) {
            ret += (n % 10) * (n % 10);
            n = n / 10;
        }
        return ret;
    }
    public static void main(String[] args){
        System.out.println(isHappy(12));
    }
}




/*
def getNext(i):
    rt = 0
    while i > 0:
        rt += pow(i%10, 2)
        i /= 10
    return rt

def check(i):
    cSet = set()
    if i <=0:
        return False
    cSet.add(i)
    while(i!=1):
        i = getNext(i)
        if i not in cSet:
            cSet.add(i)
        else:
            return False
    return True


if __name__ == '__main__':
    print check(19)
*/
