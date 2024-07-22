package com.groove.structures.bplustree;

public class BPlusTree<K extends Comparable<K>, V>  {
  private BPlusNode<K, V> root;
  private final int branchingFactor;

  public BPlusTree(int branchingFactor) {
    if (branchingFactor < 3)
      throw new IllegalArgumentException("The branching factor must be at least 3.");

    this.branchingFactor = branchingFactor;
    this.root = new BPlusNode<>(true);
  }
}
