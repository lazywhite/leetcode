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
    public Btree(Node n){
        root = n;
    }
    public void preOrder(Node n){
        if(n != null){
            System.out.print(n.val + " ");
            preOrder(n.left);
            preOrder(n.right);
        }
    }
    public void switchNode(Node n){
        if (n != null){
            Node tmp = n.left;
            n.left = n.right;
            n.right = tmp;
            switchNode(n.left);
            switchNode(n.right);

        }

    }
}

public class BTreeInvert{
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
        System.out.println();
        b1.switchNode(b1.root);
        b1.preOrder(b1.root);
        System.out.println();
    }
}
