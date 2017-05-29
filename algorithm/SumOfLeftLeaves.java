/*
Find the sum of all left leaves in a given binary tree.

Example:

    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
 */
class Node{
    public int val = -1;
    public Node left;
    public Node right;
    public Node(int val){
        this.val = val;
    }
}

class Btree{
    public Node root;
    public int depth;
    public Btree(Node n){
        root = n;
    }
    public static int getLeftSum(Node n){
        if(n == null || n.left == null){
            return 0;
        }else{
            return n.left.val + getLeftSum(n.left) + getLeftSum(n.right);
        }
    }
}

public class SumOfLeftLeaves{
    public static void main(String[] args){

        Node n1 = new Node(1);
        Node n2 = new Node(2);
        Node n3 = new Node(3);
        Node n4 = new Node(4);
        Node n5 = new Node(5);
        Node n6 = new Node(6);
        Node n7 = new Node(7);

        n1.left = n2;
        n1.right = n3;
        n2.left = n4;
        n2.right = n5;
        n4.right = n6;
        n5.left = n7;

        /*                1
         *              2   3
         *           4    5
         *            6  7
         */
 
        Btree b1 = new Btree(n1);
        System.out.println(b1.getLeftSum(b1.root));
    }
}
