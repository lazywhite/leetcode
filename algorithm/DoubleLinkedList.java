/*
 * DoubleLinkedList.java
 * Copyright (C) 2017 white <white@localhost>
 *
 * Distributed under terms of the MIT license.
 */
class Node{
    public Node pre;
    public Node next;
    public Integer value;
    public Node(Integer value){
        this.value = value;
    }

}

public class DoubleLinkedList {
     
    public static void main(String[] args){
        /*
         * 1 <--> 2 <--> 3 <--> 4
         */
        Node n1 = new Node(1);
        Node n2 = new Node(2);
        Node n3 = new Node(3);
        Node n4 = new Node(4);

        n1.next = n2;
        n2.pre = n1;
        n2.next = n3;
        n3.pre = n2;
        n3.next = n4;
        n4.pre = n3;
        show(n1);

        Node n5 =  new Node(5);
        System.out.println("after add");
        add(n5, n2, n3);
        show(n1);
        del(n5);
        System.out.println("after del");
        show(n1);
    }
    /*遍历链表*/
    public static void show(Node n){
        while(n != null){
            System.out.println(n.value);
            n = n.next;
        }
    }

    /*添加元素*/
    /*
     * @param n: 被插入的节点
     * @param n1: 插入的起始节点
     * @param n2: 插入的结束节点
     */
    public static void add(Node n, Node n1, Node n2){
        if(n1.next != n2 || n2.pre != n1){
            return;
        }
        //1. 顶端插入
        if(n1 == null){
            n2.pre = n;
            n.next = n2;
        }
        //3. 末尾插入
        if(n2 == null){
            n1.next = n;
            n.pre = n1;
        }
        //2. 中间插入
        n1.next = n;
        n.pre = n1;
        n.next = n2;
        n2.pre = n;
    }

    /* 删除节点 */
    public static void del(Node n){
        //判断位置
        if(n.pre == null){
            n.next.pre = null;
        }
        if(n.next == null){
            n.pre.next = null;
        }
//        n1 n2 n3
        n.pre.next = n.next;
        n.next.pre = n.pre;
    }
    /* 闭环判断 
     * 简历一个HashMap保留访问过的节点, 如果存在则为闭环
     * */

}

