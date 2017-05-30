/*
 * ReverseLinkedList.java
 * Copyright (C) 2017 white <white@Whites-Mac-Air.lan>
 *
 * Distributed under terms of the MIT license.
 */
class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

public class RemoveDupLinkedListNode {
    public static String getVal(ListNode node) {
        if(node == null){
            return " ";
        }
        return node.val + " " + getVal(node.next);
    }

    public static void deleteNode(ListNode n) {
        if (n.next == null){
            return;//末尾节点不可删除
        }else{
            n.val = n.next.val;
            n.next = n.next.next;
        }
    }

    public static void removeDupNode(ListNode node) {
        //循环找到后面所有跟当前节点值相同的节点， 删除之
        //递归下一个节点
        ListNode next = node.next;
        int val = node.val;

        while(next != null){
            if(next.val == val){
                deleteNode(next);
            }else{
                next = next.next;
            }
        }
        if(node.next != null){
            removeDupNode(node.next); 
        }else{
            return;
        }
    }

    public static void main(String[] args){
        ListNode n1 = new ListNode(1);
        ListNode n2 = new ListNode(2);
        ListNode n3 = new ListNode(2);
        ListNode n4 = new ListNode(3);
        ListNode n5 = new ListNode(3);
        ListNode n6 = new ListNode(4);
        ListNode n7 = new ListNode(5);

        n1.next = n2;
        n2.next = n4;
        n4.next = n3;
        n3.next = n5;
        n5.next = n6;
        n6.next = n7;

        System.out.println(getVal(n1));
        removeDupNode(n1);
        System.out.println(getVal(n1));


    }
}

