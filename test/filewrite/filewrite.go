package filewrite

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

// 产生随机数函数
func Produce(data chan int, wg *sync.WaitGroup) {
	// 产生随机数并且将 data 写入到 channel 中
	n := rand.Intn(999)
	data <- n
	// 调用 waitGroup 的 Done 方法来通知任务已经完成
	wg.Done()
}

// 将产生的随机数写入文件
func Consume(data chan int, done chan bool) {

	f, err := os.Create("../test/filewrite/testgoroutines.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

// 在现有文件中追加字符串
func WriteFileAppendString() {

	f, err := os.OpenFile("../test/filewrite/testlines.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	newline := "I love you three thousand times."

	_, err = fmt.Fprintln(f, newline)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file append successfully!")

}

// 将字符串一行一行写入文件
func WriteFileWithRawString() {
	f, err := os.Create("../test/filewrite/testlines.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	raws := []string{"I love three things in this world.", "Sun ,moon and you.", "Sun for morning, moon for night, and you forever."}

	for _, value := range raws {
		n, err := fmt.Fprintln(f, value)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(n, "bytes written successfully!")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("all raws have written successfully!")
	return
}

// 将字符串写入文件中
func WriteFileWithString() {
	f, err := os.Create("../test/filewrite/test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("i am iron man")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, " bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}

// 将字节写入文件中
func WriteFileWithByte() {
	f, err := os.Create("../test/filewrite/test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	b := []byte{105, 32, 108, 111, 118, 101, 32, 121, 111, 117}
	n, err := f.Write(b)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n, " bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 给定路径和文件名称进行字节写入
func WriteFileWithByteAndPath(path string, filename string) {
	exist, err := PathExists(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !exist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println("mkdir failed!", err)
		} else {
			fmt.Println("mkdir success!")
		}
	}

	f, err := os.Create(path + filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	b := []byte{105, 32, 108, 111, 118, 101, 32, 121, 111, 117}
	n, err := f.Write(b)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n, " bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	return

}
