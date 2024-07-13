package bplustree;

import lists.*;

public class BPlusTree<K extends Comparable<K>, V>  {
  public static class BPlusNode {
    ArrayList<Integer> items;
    ArrayList<BPlusNode> links;
  }
}
