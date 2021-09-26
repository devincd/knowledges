## Go 语言中的竞争条件
竞争条件：当两个或者以上的 goroutine 访问相同资源的时候，例如一个变量或者结构体，并且在不考虑其他 goroutine 的情况下对资源进行读写操作，
这类代码能够造成你所能见到的最令人头疼，最随机的 bugs。

## 参考文献
- https://studygolang.com/articles/15349
- https://medium.com/trendyol-tech/race-conditions-in-golang-511314c0b85
- https://golang.org/doc/articles/race_detector