package lists;

import java.util.Iterator;

public class ArrayList<T> implements Iterable<T> {
  private static final int DEFAULT_CAPACITY = 10;
  
  private T [] items;
  private int size;
  
  public ArrayList() {
    ;
  }

  public Iterator<T> iterator() {
    return new ArrayListIterator();
  }

  private class ArrayListIterator implements Iterator<T> {
    public boolean hasNext() {
      return false;
    }

    public T next() {
      return null;
    }

    public void remove() {
      ;
    }
  }
}
