# iterator
benchmark of iterator

Measure the overhead of using an iterator.

See discussion: https://github.com/gonum/graph/issues/46

I wrote a minimal graph implementation and compare speed of using a slice vs an iterator.

To make the comparison somewhat realistic I do an operation on a node value:

```
node.Value = node.Value + 1
```

## Results

```
Number of Nodes = 100,000

BenchmarkIterator-4   	    3000	    525292 ns/op
BenchmarkSlice-4      	   10000	    107560 ns/op
```

## Conclusion

The difference in speed using a slice vs an iterator is not going to dominate overall performace. Even with a simple operation the overall speed difference is not very large. Any realisitc application will have multiple operations on the nodes making the difference insignificant in relative terms.
