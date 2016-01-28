This small exercise illustrates issues with go vendoring experiment and nested vendor folders.

Comparison of two equal values (from human point of view)
```
xxx.A == yyy.A
```
will not even compile because of mismatched types.


Check this out:
```
GO15VENDOREXPERIMENT=1 go get github.com/coxx/try-go-vendoring
```


This issue is discussed in detail in [Dave Cheney's talk "Reproducible Builds"](https://www.youtube.com/watch?v=c3dW80eO88I).


