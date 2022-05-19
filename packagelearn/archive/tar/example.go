package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// write content to test file
	test1File := filepath.Join(currentDir, "test-1.txt")
	test2File := filepath.Join(currentDir, "test-2.txt")
	test3File := filepath.Join(currentDir, "test-3.txt")
	src := []string{test1File, test2File, test3File}
	dst := filepath.Join(currentDir, "test.tar")
	for index, fileName := range src {
		fw, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fw.Write([]byte(fmt.Sprintf("test-%d", index+1)))
		if err != nil {
			log.Fatal(err)
		}
		fw.Close()
	}
	err = Tar(src, dst)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("call Tar function successfully")

	// delete the file
	for _, fileName := range src {
		err = os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
		}
	}
	// 执行解包操作
	err = Untar(dst)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("call Untar function successfully")
}

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
