package com.groove.structures.bplustree;

public class TestBPlusTree {
  public static void main(String[] args) {
    BPlusTree<Integer, String> bPlusTree = new BPlusTree<>(3);

    for (int i = 1; i < 1000; i++)
      bPlusTree.insert(i, "Se pudo!");

    System.out.println("Árbol B+ después de las inserciones:");
    bPlusTree.print();
  }
}
