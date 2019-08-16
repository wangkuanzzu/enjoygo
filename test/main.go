package main

import (
	"test/reflection"
)

func main() {
	/*
		fmt.Println("begin test your algorithm...")

		arr := []int{3, 4, 5, 6, 2, 3, 4, 2, 2, 32, 12, 56, 2, 3, 24, 8}
		fmt.Println("sort ahead", arr)
		// 快速排序---以数组为操作对象实现
		helloworld.QuickSort(arr, 0, len(arr)-1)
		// 插入排序---以数组为操作对象实现
		helloworld.InsertionSort(arr)
		// 选择排序---以数组为操作对象实现
		helloworld.SelectionSort(arr)
		fmt.Println("sort behind", arr)

		// 求数组中出现次数最多的元素和出现的次数
		status, maxCount, maxValue := algorithm.ArrayCountValues(arr)
		fmt.Println("results:", status, maxCount, maxValue)

		fmt.Println(" end  test your algorithm!!!")
	*/

	// 将字符串写入文件中
	//filewrite.WriteFileWithString()
	// 将字节写入文件中
	//filewrite.WriteFileWithByte()

	// 字符串转bytes
	// b := "i love you"
	// data := []byte(b)
	// fmt.Println("===", data)

	/*
		filepath := "C:/love/"
		var filename string
		for index := 0; index < 10; index++ {
			filename = "love" + strconv.Itoa(index) + ".txt"
			filewrite.WriteFileWithByteAndPath(filepath, filename)
		}
	*/

	// filewrite.WriteFileWithRawString()

	// filewrite.WriteFileAppendString()

	/*
		// 创建写入和读取数据的 channel
		   	data := make(chan int)
		   	// 创建done这个chan，此chan用于消费者goroutinue 完成任务之后通知主函数
		   	done := make(chan bool)
		   	// 创建 Waitgroup 的实例 wg，用于等待所有生产随机数的 goroutine 完成任务
		   	wg := sync.WaitGroup{}
		   	// 使用 for 循环创建 100 个 goroutines
		   	for i := 0; i < 100; i++ {
		   		wg.Add(1)
		   		go filewrite.Produce(data, &wg)
		   	}
		   	// 创建消费者goroutine 将data chan中的随机数写入文件，并将写入文件的成功失败的标识（true false）写入done chan中
		   	go filewrite.Consume(data, done)
		   	// 调用 waitgroup 的 wait() 方法等待所有的 goroutines 完成随机数的生成。然后关闭 channel。
		   	go func() {
		   		wg.Wait()
		   		close(data)
		   	}()

		   	// 主函数 读取done chan中表示进行打印信息
		   	d := <-done
		   	if d {
		   		fmt.Println("file written successfully")
		   	} else {
		   		fmt.Println("file written failed")
			   }
	*/

	//简单爬虫
	//reptile.SimpleReptile()
	/*
		//字符串转bytes，测试单笔交易所占字节数
		b := "{'hash':'83c630f87540aab3dc4249cde32188fa20660d0666ae84ee0cad94616d469473','height':'620347','tx':{'type':'qbase/txs/stdtx',"
		b += "'value':{'itx':[{'type':'transfer/txs/TxTransfer','value':{'senders':[{'addr':'address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa','qos':'100','qscs':null}],"
		b += "'receivers':[{'addr':'address1l0wn66gh45nfta2r4vq8z54wu9hgarss298e9g','qos':'100','qscs':null}]}}],"
		b += "'sigature':[{'pubkey':{'type':'tendermint/PubKeyEd25519','value':'PK4Yw0gH14QU0jtcJnGZN5paeA151HIbkMrooXGdxT4='},"
		b += "'signature':'TpDLM6IpMydf+p5iYJq6PQFAptaO7tN5BuvOCsEpAvUEaHnC7NKun4OfOo1saCFxUFmHa5UUfp2VYq7xRytJDw==','nonce':'16'}],"
		b += "'chainid':'capricorn-3000','maxgas':'100000'}},'result':{'gas_wanted':'100000','gas_used':'18190',"
		b += "'tags':[{'key':'c2VuZGVy','value':'YWRkcmVzczFubnZkcWVmdmE4OXh3cHB6czQ2dnVza2NrcjdrbHZ6azhyNXVhYQ=='},"
		b += "{'key':'cmVjZWl2ZXI=','value':'YWRkcmVzczFsMHduNjZnaDQ1bmZ0YTJyNHZxOHo1NHd1OWhnYXJzczI5OGU5Zw=='}]}}"
		data := []byte(b)
		fmt.Println(len(data))
		fmt.Println("===", data)
	*/

	//学习反射
	o := reflection.NewOrder(123, 234)
	reflection.CreateQuery(o)

}
