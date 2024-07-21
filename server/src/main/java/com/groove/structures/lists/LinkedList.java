package com.groove.structures.lists;

import java.util.ConcurrentModificationException;
import java.util.NoSuchElementException;

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

    public Node(T d, Node<T> p, Node<T> n){ 
      data = d;
      prev = p;
      next = n;
    }
  }

  public LinkedList() {
    doClear();
  }

  public void clear() {
    doClear();
  }

  private void doClear() {
    beginMarker = new Node<T>(null, null, null);
    endMarker = new Node<T>(null, beginMarker, null);
    beginMarker.next = endMarker;

    size = 0;
    modCount++;
  }

  public int size() {
    return size;
  }

  public boolean isTmpty() {
    return size() == 0;
  }

  public boolean add(T x) {
    add(size(), x);
    return true;
  }

  public void add(int idx, T x) {
    addBefore(getNode(idx, 0, size()), x);
  }

  public T get(int idx) {
    return getNode(idx).data;
  }

  public T set(int idx, T newVal) {
    Node<T> p = getNode(idx);
    T oldVal = p.data;
    p.data = newVal;
    return oldVal;
  }

  public T remove(int idx) {
    return remove(getNode(idx));
  }

  /**
   * Adds an item to this collection, at specified position p.
   * Items at or after that position are slid one position higher.
   * @param p Node to add before.
   * @param x any object.
   * @throws IndexOutOfBoundsException if idx is not between 0 and size(),.
   */
  private void addBefore(Node<T> p, T x) {
    Node<T> newNode = new Node<>(x, p.prev, p);
    newNode.prev.next = newNode;
    p.prev = newNode;

    p.prev.next = p.prev = new Node<T>(x, p.prev, p);
    size++;
    modCount++;
  }

  /**
   * Gets the Node at position idx, which must range from 0 to size() - 1.
   * @param idx index to search at.
   * @return internal node corresponding to idx.
   * @throws IndexOutOfBoundsException if idx is not
   * between 0 and size() - 1, inclusive.
  */
  private Node<T> getNode(int idx) {
    return getNode(idx, 0, size() - 1);
  }

  /**
   * Gets the Node at position idx, which must range from lower to upper.
   * @param idx index to search at.
   * @param lower lowest valid index.
   * @param upper highest valid index.
   * @return internal node corresponding to idx.
   * @throws IndexOutOfBoundsException if idx is not
   * between lower and upper, inclusive.
  */
  private Node<T> getNode(int idx, int lower, int upper) {
    Node<T> p;

    if(idx < lower || idx > upper)
      throw new IndexOutOfBoundsException();

    if(idx < size() / 2) {
      p = beginMarker.next;
      for(int i = 0; i < idx; i++)
        p = p.next;
    } else {
      p = endMarker;
      for(int i = size(); i > idx; i--)
        p = p.prev;
    }

    return p;
  }

  /**
   * Removes the object contained in Node p.
   * @param p the Node containing the object.
   * @return the item was removed from the collection.
  */
  private T remove(Node<T> p) {
    p.next.prev = p.prev;
    p.prev.next = p.next;
    size--;
    modCount++;

    return p.data;
  }

  public java.util.Iterator<T> iterator() {
    return new LinkedListIterator();
  }

  private class LinkedListIterator implements java.util.Iterator<T> {
    private Node<T> current = beginMarker.next;
    private int expectedModCount = modCount;
    private boolean okToRemove = false;

    public boolean hasNext() {
      return current != endMarker;
    }

    public T next() {
      if(modCount != expectedModCount)
        throw new ConcurrentModificationException();

      if(!hasNext())
        throw new NoSuchElementException();

      T nextItem = current.data;
      current = current.next;
      okToRemove = true;
      return nextItem;
    }

    public void remove() {
      if(modCount != expectedModCount)
        throw new ConcurrentModificationException();

      if(!okToRemove)
        throw new IllegalStateException();

      /**
       * Ojito: Because we call the remove method of MyLinkedList, 'modCount' is modified.
       * Futhermore, expectedModCount also must be modified.
       */
      LinkedList.this.remove(current.prev);
      expectedModCount++;
      okToRemove = false;
    }
  }
}
