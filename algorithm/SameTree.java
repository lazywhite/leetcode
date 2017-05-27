/*
   Given two binary trees, write a function to check if they are equal or not.

   Two binary trees are considered equal if they are structurally identical and the nodes have the same value.
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
    public StringBuilder treeStr;
    public Btree(Node n){
        root = n;
        treeStr = new StringBuilder();
    }
    public void preOrder(Node n){
        if(n != null){
            treeStr.append(n.val);
            preOrder(n.left);
            preOrder(n.right);
        }
    }
}

public class SameTree{
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


        Btree b1 = new Btree(n1);
        b1.preOrder(b1.root);

        Btree b2 = new Btree(n2);
        b2.preOrder(b2.root);
        if(b1.treeStr.toString() == b2.treeStr.toString()){

            System.out.println("they are same tree");

        }else{
            System.out.println("they are not the same tree");

        }
    }
}
