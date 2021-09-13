## Pyroscope take-home assignment

### Summary

You're given a program that consists of two parts:
* `tree.go` — implementation of a tree that's used to store profiling data. The format is very similar to one we use in pyroscope. Here's a diagram that illustrates the data structure:

![tree representation](https://user-images.githubusercontent.com/662636/133024145-146b9ce1-88cd-40ee-b924-6ad4dc93dcab.png)

* `main.go` — code that creates hundreds of synthesized (generated) trees and measures how many objects are created on heap + how long GC takes. It usually outputs something like this:

```shell
$ go run ./
1. allocates one simple tree
  * garbage collection took: 0s
  * objects allocated: 9
2. measures GC time as a function of number of objects on heap
allocating 100 trees with width 200 and depth 500
  * garbage collection took: 45.29787ms
  * objects allocated: 39166322
```

### Problem

As users ingest more and more data, the number of trees as well as their size increases, meaning there's more and more individual objects that Go's garbage collector has to keep track of.

Even though it is pretty smart we found that at about 100,000,000 objects garbage collection starts to take a significant amount of time and ends up taking up to 50% of CPU time.

### Objective

For this exercise you need to find a way to reduce the number of objects that are allocated on the heap.

You're can modify the main code (`main.go`) and the data structures as much as you'd like. The only constraint is that you have to implement `Insert` and `Iterate` methods — they are implemented in the reference implementation (`tree.go`) and provide simple ways to write to and read from profiling trees.


