## 考察知识点
- `defer`关键字
- `panic`关键字

## 程序结果
```
after doing
doing
before doing
panic: trigger a panic

goroutine 1 [running]:
main.deferCall()
        /Users/devin/work/golang/workspace/github.com/devincd/hardworking/golang-interview/chap1/001/main.go:14 +0x85
main.main()
        /Users/devin/work/golang/workspace/github.com/devincd/hardworking/golang-interview/chap1/001/main.go:6 +0x25

Process finished with exit code 2

```

## 原因分析
### `defer`关键字
Go语言的`defer`会在当前函数或者方法返回之前执行传入函数，它常常被用于关闭文件描述符、关闭数据库连接以及解锁资源。作为编程语言中的关键字，
`defer`的实现一定是有编译器和运行时共同完成的，不过在深入源码分析它之前我们还是需要了解`defer`关键字的常见使用场景以及使用时注意事项。

#### 现象
我们在Go语言中使用`defer`时会遇到两个比较常见的问题，这里会介绍具体的场景并分析这两个现象背后的设计原理
- `defer`关键字的调用时机以及多次调用`defer`时执行顺序是如何确定的
- `defer`关键字使用传值的方式传递参数时会进行预计算，导致不符合预期的结果

##### 作用域
`defer`传入的函数不是在退出代码块的作用域执行的，它只会在当前函数和方法返回之前被调用，并且`defer`的执行顺序按照栈的数据结构，先进后出的
方式来调用多个`defer`

##### 预计算参数
Go语言中的所有的函数调用都是传值的，`defer`虽然是关键字，但是也继承了这个特性。假设我们想要计算`main`函数运行的时间，可能会写以下的代码:
```
package main

import (
	"fmt"
	"time"
)

func main() {
	startedAt := time.Now()

	defer fmt.Printf("main cost %v\n", time.Since(startedAt))

	time.Sleep(time.Second)
}

$ go run main.go
main cost 1.153µs
```
然而上述代码的运行结果并不符合我们的预期，这个现象背后的原因是什么呢？经过分析，我们会发现调用`defer`关键字会立刻对函数中引用的外部参数进行
拷贝(将本次`time.Since(startedAt)`执行的结果存储起来)，所以`time.Since(startedAt)`的结果不是在`main`函数退出之前计算的，而是在`defer`
关键字调用时计算的，最终导致结果不符合预期

想要解决这个办法非常简单，我们只需要向`defer`关键字传入匿名函数
```
package main

import (
	"fmt"
	"time"
)

func main() {
	startedAt := time.Now()

	defer func() {fmt.Printf("main cost %v\n", time.Since(startedAt))}()

	time.Sleep(time.Second)
}

$ go run main.go
main cost 1.000662083s
```
虽然调用`defer`关键字时也使用值传递，但是因为拷贝的是函数指针，所以`time.Since(startedAt)`会在`main`函数返回前被调用并打印出符合预期的
结果

##### 现象小结
- defer语句并不会马上执行，而是会进入一个栈，函数return前，会按先进后出的顺序执行。也说是说最先被定义的defer语句最后执行。先进后出的原因
是后面定义的函数可能会依赖前面的资源，自然要先执行；否则，如果前面先执行，那后面函数的依赖就没有了。
- 在defer函数定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。作为函数参数，则在defer定义时就把值传递给defer，
并被cache起来；作为闭包引用的话，则会在defer函数真正调用时根据整个上下文确定当前的值。
  



### `panic`关键字
- `panic`关键字能够改变程序的控制流，函数调用`panic`时会立刻停止执行函数的其他代码,并在执行结束后在当前Goroutine中递归执行调用方延迟函数调
用`defer`
- `recover`可以中止`panic`造成的程序崩溃，它是一个只能在`defer`中发挥作用的函数，在其他作用域中调用不会发挥任何作用

结论
- `panic`只会触发当前Goroutine的延迟函数调用
- `recover`只能在`defer`函数中调用才会生效


## 参考文献
- https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/
- https://zhuanlan.zhihu.com/p/56557423