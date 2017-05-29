/*
Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:
Input: 
         1
       /   \
      2     3
Output: 1
Explanation: 
Tilt of node 2 : 0
Tilt of node 3 : 0
Tilt of node 1 : |2-3| = 1
Tilt of binary tree : 0 + 0 + 1 = 1
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
    public static int getRightSum(Node n){
        if(n == null || n.right == null){
            return 0;
        }else{
            return n.right.val + getRightSum(n.left) + getRightSum(n.right);
        }
    }
}

public class BTreeTilt{
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
        System.out.println(Math.abs(b1.getLeftSum(b1.root)-b1.getRightSum(b1.root)));
    }
}
