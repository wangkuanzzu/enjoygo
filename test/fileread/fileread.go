package fileread

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gobuffalo/packr"
)

// 按照给定字节数 读取文件
func ReadFile4() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 20)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}

// read file a raw one time  按行读取文件内容
func ReadFile3() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// read all contents of file once time
func ReadFile2() {
	box := packr.NewBox("../filehandling")
	data := box.String("test.txt")
	fmt.Println("Contents of file :" + data)
}

// read all contents of file once time
func ReadFile1() {

	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is ", fptr)
	fmt.Println("value of fpath is ", *fptr)
	//data, err := ioutil.ReadFile("test.txt")
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error ", err)
		return
	}
	fmt.Println("content of file :", string(data))
}
