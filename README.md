# Realization algorithms in Golang

## Index<a name=title>
### Search:

   | №  | Algorithm                             | Complexity   | Code                             |
   | :- | :---------                            | :-------     |:------:                          |
   | 1  |  [Linear search](#linear_search)      |  O(n)        | [code](/search/linear_search.go) |
   | 2  |  [Binary search](#binary_search)      |  O(log n)    | [code](/search/binary_search.go) |
   | 3  |  [Jump search](#jump_search)          |  O(sqrt n)   | [code](/search/jump_search.go)   |


### Sorting:

   | №  | Algorithm                             | Average      | Best         | Worst      | Memory |
   | :- | :---------                            | :-------     |:------:      | :-----:    | :-----:|
   | 1  |  [Selection sort](#selection_sort)    |  O(n^2)      | O(n^2)       | O(n^2)     |   O(1) |
   | 2  |  [Bubble sort](#bubble_sort)          |  O(n^2)      | O(n)         | O(n^2)     |   O(1) |
   | 3  |  [Merge sort](#merge_sort)            |  O(n log n)  | O(n log n)   | O(n log n) |   O(n) |
   | 4  |  [Quick sort](#quick_sort)            |  O(n log n)  | O(n log n)   | O(n^2)     |   O(n) |
   | 5  |  [Insertion sort](#insertion_sort)    |  O(n^2)      | O(n)         | O(n^2)     |   O(1) |


### Fibonacci:

   1. [Fibonacci](#fibonacci)



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
