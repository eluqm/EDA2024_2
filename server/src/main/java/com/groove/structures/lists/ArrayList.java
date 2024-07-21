package com.groove.structures.lists;

import java.util.Iterator;
import java.util.NoSuchElementException;

public class ArrayList<T> implements Iterable<T> {
  private static final int DEFAULT_CAPACITY = 10;
  
  private T [] items;
  private int size;
  
  /**
   * Construct an empty ArrayList.
   */
  public ArrayList() {
    doClear();
  }
  
  /**
   * @return the number of items in this collection.
   */
  public int size() {
    return size;
  }
  
  /**
   * @return true if  this collection is empty.
   */ 
  public boolean isEmpty() {
    return size() == 0;
  }
  
  /**
   * Returns the item at position idx.
   * @param idx the index to search in.
   * @throws ArrayIndexOutOfBoundsException if  index is out of range.
   */
  public T get(int idx) {
    if (idx < 0 || idx >= size())
      throw new ArrayIndexOutOfBoundsException("Index " + idx + "; size " + size());
    return items[idx];    
  }
      
  /**
   * Changes the item at position idx.
   * @param idx the index to change.
   * @param newVal the new value.
   * @return the old value.
   * @throws ArrayIndexOutOfBoundsException if  index is out of range.
   */
  public T set(int idx, T newVal) {
    if (idx < 0 || idx >= size())
      throw new ArrayIndexOutOfBoundsException("Index " + idx + "; size " + size());

    T old = items[idx];    
    items[idx] = newVal;
    return old;    
  }

  @SuppressWarnings("unchecked")
  public void ensureCapacity(int newCapacity) {
    if (newCapacity < size)
      return;

    T [] old = items;
    items = (T []) new Object[newCapacity];

    for (int i = 0; i < size(); i++)
      items[i] = old[i];
  }
  
  /**
   * Adds an item to this collection, at the end.
   * @param x any object.
   * @return true.
   */
  public boolean add(T x) {
    add(size(), x);
    return true;            
  }
  
  /**
   * Adds an item to this collection, at the specif ied index.
   * @param x any object.
   * @return true.
   */
  public void add(int idx, T x) {
    if (items.length == size())
      ensureCapacity(size() * 2 + 1);

    for (int i = size; i > idx; i--)
      items[i] = items[i - 1];

    items[idx] = x;
    size++;  
  }
    
  /**
   * Removes an item from this collection.
   * @param idx the index of the object.
   * @return the item was removed from the collection.
   */
  public T remove(int idx) {
    T removedItem = items[idx];
    
    for (int i = idx; i < size() - 1; i++)
      items[i] = items[i + 1];
    size--;    
    
    return removedItem;
  }
  
  /**
   * Change the size of this collection to zero.
   */
  public void clear() {
    doClear();
  }
  
  private void doClear() {
    size = 0;
    ensureCapacity(DEFAULT_CAPACITY);
  }
  
  /**
   * Obtains an Iterator object used to traverse the collection.
   * @return an iterator positioned prior to the first element.
   */
  public Iterator<T> iterator() {
    return new ArrayListIterator();
  }

  /**
   * Returns a String representation of this collection.
   */
  public String toString() {
    StringBuilder sb = new StringBuilder("[");

    for (T x : this)
        sb.append(x + " ");
    sb.append("]");

    return new String(sb);
  }

  private class ArrayListIterator implements Iterator<T> {
    private int current = 0;

    public boolean hasNext() {
      return current < size();
    }

    public T next() {
      if(!hasNext())
        throw new NoSuchElementException();
      return items[current++];
    }

    public void remove() {
      ArrayList.this.remove(--current);
    }
  }
}
