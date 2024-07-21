package com.groove.structures.lists;

public class LinkedList<T> implements Iterable<T> {
  private int size;
  private int modCount = 0;     // The number of changes to the linked list since construction.
  private Node<T> beginMarker;
  private Node<T> endMarker;

  private static class Node<T> { 
    // Item stored
    public T data;

    // Links to the previous and next Node
    public Node<T> prev;
    public Node<T> next;

    public Node( T d, Node<T> p, Node<T> n ){ 
      data = d;
      prev = p;
      next = n;
    }
  }
}
