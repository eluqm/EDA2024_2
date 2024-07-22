package com.groove.structures.bplustree;

import java.util.ArrayList;
import java.util.List;

class BPlusNode<K extends Comparable<K>, V> {
  List<K> keys;
  List<V> values;   // Just uses in leafs
  List<BPlusNode<K, V>> children;
  BPlusNode<K, V> next;   // Maintains the link between leaf nodes
  boolean isLeaf;

  BPlusNode(boolean isLeaf) {
    this.keys = new ArrayList<>();
    this.isLeaf = isLeaf;
    if (isLeaf)
      this.values = new ArrayList<>();
    else
      this.children = new ArrayList<>();
  }
}
