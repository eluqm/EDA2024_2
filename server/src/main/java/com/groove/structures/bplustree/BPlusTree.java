package com.groove.structures.bplustree;

import java.util.Collections;

public class BPlusTree<K extends Comparable<K>, V>  {
  private BPlusNode<K, V> root;
  private final int branchingFactor;

  public BPlusTree(int branchingFactor) {
    if (branchingFactor < 3)
      throw new IllegalArgumentException("The branching factor must be at least 3.");

    this.branchingFactor = branchingFactor;
    this.root = new BPlusNode<>(true);
  }

  public void insert(K key, V value) {
    BPlusNode<K, V> newNode = insertInternal(root, key, value);
    if (newNode != null) {
      BPlusNode<K, V> newRoot = new BPlusNode<>(false);
      newRoot.keys.add(newNode.keys.get(0));
      newRoot.children.add(root);
      newRoot.children.add(newNode);
      root = newRoot;
    }
  }

  private BPlusNode<K, V> insertInternal(BPlusNode<K, V> node, K key, V value) {
    int idx = Collections.binarySearch(node.keys, key);

    if (node.isLeaf) {
      int insertIndex = idx >= 0 ? idx : -idx - 1;  // Index to insert the new key-value

      if (idx >= 0) {
        node.values.set(insertIndex, value);
      } else {
        node.keys.add(insertIndex, key);
        node.values.add(insertIndex, value);
        if (node.keys.size() >= branchingFactor)
          return splitLeafNode(node);
      }
    } else {
      int childIndex = idx >= 0 ? idx + 1 : -idx - 1;   

      BPlusNode<K, V> newNode = insertInternal(node.children.get(childIndex), key, value);

      if (newNode != null) {
        node.keys.add(childIndex, newNode.keys.get(0));
        node.children.add(childIndex + 1, newNode);
        if (node.keys.size() >= branchingFactor)
          return splitInternalNode(node);
      }
    }
    return null;
  }

  private BPlusNode<K, V> splitLeafNode(BPlusNode<K, V> node) {
    int mid = branchingFactor / 2;
    BPlusNode<K, V> newNode = new BPlusNode<>(true);

    // Transfer keys and values/pointers
    newNode.keys.addAll(node.keys.subList(mid, node.keys.size()));
    newNode.values.addAll(node.values.subList(mid, node.values.size()));
    node.keys.subList(mid, node.keys.size()).clear();
    node.values.subList(mid, node.values.size()).clear();

    newNode.next = node.next;
    node.next = newNode;
    return newNode;
  }

  private BPlusNode<K, V> splitInternalNode(BPlusNode<K, V> node) {
    int mid = branchingFactor / 2;
    BPlusNode<K, V> newNode = new BPlusNode<>(false);

    // Transfer keys and values/pointers
    newNode.keys.addAll(node.keys.subList(mid + 1, node.keys.size()));
    newNode.children.addAll(node.children.subList(mid + 1, node.children.size()));
    node.keys.subList(mid, node.keys.size()).clear();
    node.children.subList(mid + 1, node.children.size()).clear();

    return newNode;
  }

  // Method to print the tree
  public void print() {
    print(root, "", true);
  }

  private void print(BPlusNode<K, V> node, String prefix, boolean isTail) {
    System.out.println(prefix + (isTail ? "└── " : "├── ") + node.keys + " " + (node.next != null ? node.next.keys : ""));
    if (!node.isLeaf) {
      for (int i = 0; i < node.children.size() - 1; i++)
        print(node.children.get(i), prefix + (isTail ? "    " : "│   "), false);

      if (node.children.size() > 0)
        print(node.children.get(node.children.size() - 1), prefix + (isTail ? "    " : "│   "), true);
    }
  }
}
