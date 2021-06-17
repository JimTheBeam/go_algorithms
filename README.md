# Realization algorithms in Golang

## Index<a name=title>
### Search:

   1. [Linear search](#linear_search)

   2. [Binary search](#binary_search)

   3. [Jump search](#jump_search)


### Sorting:

   1. [Selection sort](#selection_sort)

   2. [Bubble sort](#bubble_sort)

   3. [Merge sort](#merge_sort)

   4. [Quick sort](#quick_sort)

   5. [Insertion sort](#insertion_sort)


### Fibonacci:

   1. [Fibonacci](#fibonacci)


## Search
---

### Linear search<a name=linear_search>

complexity O(n)

[code](/search/linear_search.go)

[to title](#title)

### Binary search<a name=binary_search>

complexity O(log n)

[code](/search/binary_search.go)

[to title](#title)


### Jump search<a name=jump_search>

complexity O(sqrt n)

[code](/search/jump_search.go)

[to title](#title)


## Sorting
---

### Selection sort<a name=selection_sort>

Complexity: 
   | Average  | Best   | Worst  | Memory |
   | :------- |:------:| :-----:| :-----:|
   |  O(n^2)  | O(n^2) | O(n^2) |   O(1) |

[code](/sorting/selection_sort.go)

![image](/img/selectionsort.gif)

[to title](#title)



### Bubble sort<a name=bubble_sort>

Complexity: 
   | Average  | Best   | Worst  | Memory |
   | :------- |:------:| :-----:| :-----:|
   |  O(n^2)  | O(n)   | O(n^2) |   O(1) |
   
Best case: array already sorted

[code](/sorting/bubble_sort.go)

![image](/img/bubblesort.gif)

[to title](#title)




### Merge sort<a name=merge_sort>

Complexity: 
   | Average      | Best         | Worst      | Memory |
   | :-------     |:------:      | :-----:    | :-----:|
   |  O(n log n)  | O(n log n)   | O(n log n) |   O(n) |

[code](/sorting/merge_sort.go)

![image](/img/mergesort.gif) 

[to title](#title)


### Quick sort<a name=quick_sort>


Complexity: 
   | Average      | Best         | Worst      | Memory |
   | :-------     |:------:      | :-----:    | :-----:|
   |  O(n log n)  | O(n log n)   | O(n^2)     |   O(n) |

[code](/sorting/quick_sort.go)


![image](/img/quicksort.gif)

[to title](#title)


### Insertion sort<a name=insertion_sort>



Complexity: 
   | Average  | Best   | Worst  | Memory |
   | :------- |:------:| :-----:| :-----:|
   |  O(n^2)  | O(n)   | O(n^2) |   O(1) |


[code](/sorting/insertion_sort.go)

![image](/img/insertionsort.gif)

[to title](#title)



## Fibonacci
---

### Fibonacci<a name=fibonacci>

   Reqursive fibbonacci algorithm:
complexity 2^n

   Algorithm with arrays:
complexity O(n)


[code](fibonacci/fibonacci.go)

[to title](#title)
