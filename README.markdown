# vptree

vptree is a port of Steve Hanov's C++
[implementation](http://stevehanov.ca/blog/index.php/?id=130) of [Vantage-point
trees](https://en.wikipedia.org/wiki/Vantage-point_tree) to the Go programming
language. Vantage-point trees are useful for nearest-neighbour searches in
high-dimensional metric spaces.

This is a modified version https://github.com/dgryski/go-simstore/vptree

It has been customized to search specifically for Items consisting of a 64-bit
signature and an string ID using Hamming distance as the distance measure.

Also it has no external dependencies.
