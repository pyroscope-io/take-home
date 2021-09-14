## Pyroscope take-home assignment

### Context

[Pyroscope](https://github.com/pyroscope-io/pyroscope/) is an open source continuous profiling platform. The system consists of a server that acts as database and client implementations for various programming languages such as Go, Ruby and Python.

This take home assignment is based on a real problem we faced when building pyroscope server database. The code sample provided here is a slimmed down version of the code we use in pyroscope.

If you're not familiar with pyroscope or you want to learn more about the storage design you might find these links useful:

* [pyroscope demo](https://demo-dev.pyroscope.io/) — this can be helpful for understanding what the data looks like on the user side
* [storage design doc](https://github.com/pyroscope-io/pyroscope/blob/main/docs/storage-design.md) — this is an in-depth animated guide to the storage design

### Summary

You're given a program that consists of two parts:
* `tree.go` — implementation of a tree that's used to store profiling data. The format is very similar to one we use in pyroscope. Here's a diagram that illustrates the data structure:

![tree representation](https://user-images.githubusercontent.com/662636/133024145-146b9ce1-88cd-40ee-b924-6ad4dc93dcab.png)

* `main.go` — code that creates hundreds of synthesized (generated) trees and measures how many objects are created on heap + how long garbage collection (GC) takes. It usually outputs something like this:

```shell
$ go run ./
1. allocates one simple tree
  * objects allocated: 9
2. measures GC time as a function of number of objects on heap
allocating 100 trees with width 200 and depth 500
  * objects allocated: 39166322
```

### Problem

As users ingest more and more data, the number of trees as well as their size increases, meaning there's more and more individual objects that Go's garbage collector has to keep track of.

Even though it is pretty smart we found that at about 100,000,000 objects garbage collection starts to take a significant amount of time and ends up taking up to 50% of CPU time, [here's an example from our demo pyroscope server profiling itself](https://demo.pyroscope.io/?name=pyroscope.server.cpu%7B%7D&from=now-30d).

### Objective

The goal for this exercise is to find a way to reduce the number of objects that are allocated on the heap.

You can modify the main code (`main.go`) and the data structures as much as you'd like. The only constraint is that you have to implement `Insert` and `Iterate` methods — they are implemented in the reference implementation (`tree.go`) and provide simple ways to write to and read from profiling trees.

### How to run the code

Assuming you have go installed on your machine it should be as easy as running this command:
```shell
go run ./
```
