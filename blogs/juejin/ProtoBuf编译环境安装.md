## ProtoBuf 编译环境
安装 protobuf 的编译器 protoc。protoc 需要 protoc-gen-go 来完成 Go 语言的代码转换，
因此我们需要安装 protoc 和 protoc-gen-go 这两个工具，它们的安装方法比较简单，
直接看下面给出的代码和操作注释就可以了
```shell
# 第一步：安装protobuf
$ cd /tmp/
$ git clone --depth=1 https://github.com/protocolbuffers/protobuf
$ cd protobuf
$ ./autogen.sh
$ ./configure
$ make
$ sudo make install
$ protoc --version   # 查看 protoc 版本，成功输出版本号，说明安装成功
libprotoc 3.15.6

# 第二步：安装 protoc-gen-go
$ go get -u github.com/golang/protobuf/protoc-gen-go
```