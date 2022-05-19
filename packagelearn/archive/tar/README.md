## Golang 中 archive/tar 包详解
go 语言官方标准库提供 tar 库，以方便对 tar 的操作

### 一、了解 tar
什么是 tar ？

tar 是一种打包格式，但不对文件进行压缩，所以打包后的文档一般远远大于 zip 和 tar.gz，
因为不需要压缩的原因，所以打包的速度是非常快的，打包时 CPU 占用率也很低。

tar 的目的是什么？

方便文件的管理（帮助理解：就是你存在很多文件的时候，但是你很多要很长时间不去接触的话，
你想要变得更加简洁，可以进行 tar 操作，就可以变的更简洁，比如就像生活中，有很多小箱子分散在不同的房间里，
可以将小箱子放在一个房间里，tar 可以类似这样）。

下面一张图可以帮助很好的理解
![img.png](img.png)

### 二、tar 的操作
打包、解包

### 三、go 如何对 tar 文件操作
打包操作：
- 生成打包后的目标文件
- 获取要打包的文件集
- 往目标文件写入文件

接着给出一个代码案例：
```go
// Tar 打包操作
// 步骤如下：
// 1. 生成打包后的目标文件
// 2. 获取要打包的文件集
// 3. 往目标文件写入文件
func Tar(src []string, dst string) error {
	// (1) 生成打包后的目标文件
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(fw *os.File) {
		err := fw.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(fw)

	tw := tar.NewWriter(fw)
	// 如果关闭失败会造成 tar 包不完整
	defer func() {
		if err := tw.Close(); err != nil {
			log.Println(err)
		}
	}()
	// (2) 获取要打包的文件集
	for _, fileName := range src {
		fi, err := os.Stat(fileName)
		if err != nil {
			log.Println(err)
			continue
		}
		hdr, err := tar.FileInfoHeader(fi, "")
		// 将 tar 的文件信息 hdr 写入到 tw
		err = tw.WriteHeader(hdr)
		if err != nil {
			return err
		}
		// (3) 往目标文件写入文件
		fs, err := os.Open(fileName)
		if err != nil {
			return err
		}
		if _, err := io.Copy(tw, fs); err != nil {
			return err
		}
		err = fs.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
```


解包操作：
- 打开 tar 文件
- 遍历 tar 中文件信息
- 创建文件，写入，保存，关闭文件

接着给出一个代码案例：
```go
// Untar 解包操作
// 步骤如下：
// 1. 打开 tar 文件
// 2. 遍历 tar 中文件信息
// 3. 创建文件，写入，保存，关闭文件
func Untar(srcFile string) error {
	// (1) 打开 tar 文件
	fr, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()

	// (2) 遍历 tar 中文件信息
	tr := tar.NewReader(fr)
	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		if err != nil {
			log.Println(err)
			continue
		}

		// 读取文件信息
		fi := hdr.FileInfo()

		// 创建一个空文件，用来写入解包后的数据
		fw, err := os.Create(fi.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		if _, err := io.Copy(fw, tr); err != nil {
			log.Println(err)
		}
		os.Chmod(fi.Name(), fi.Mode().Perm())
		fw.Close()
	}
	return nil
}
```

### 四、tar 包深入学习
作为一名计算机专业人士，至少要知道原理和实现的，接着我们来分析下原理和实现

打包和解包的原理和实现

#### 打包实现原理
1. 先创建一个文件 x.tar，然后向 x.tar 写入 tar 头部信息。
2. 打开要被 tar 的文件，向 x.tar 写入头部信息，然后向 x.tar 写入文件信息。
3. 重复第二步直到所有文件都被写入到 x.tar 中，关闭 x.tar，整个过程就这样完成了。

#### 解包实现原理
先打开 tar 文件，然后从这个 tar 头部循环读取存储在这个归档文件内的头部信息，
从这个文件头里读取文件名，以这个文件名创建文件，然后向这个文件写入数据。

#### Go 标准解包


### 相关参考链接
