package main

func main() {



	//链表数据结构的构建和基本方法测试
	//linkList:=LinkList.CreateLinkList()
	//fmt.Println("链表创建成功：")
	//linkList.PrintInfo()
	//fmt.Println()
	//fmt.Println("---------------------------现在开始测试插入数据-------------------------")
	//fmt.Println("在空链表尾部加一个int 1")
	//linkList.Append(1)
	//linkList.PrintInfo()
	//
	//fmt.Println()
	//fmt.Println("在头部加一个int 0")
	//linkList.InsertHead(0)
	//linkList.PrintInfo()
	//
	//fmt.Println()
	//fmt.Println("在0后面插入一个int 2")
	//linkList.InsertAfterData(0,2)
	//linkList.PrintInfo()
	//
	//fmt.Println()
	//fmt.Println("在下标0插入一个int 100")
	//flag:=linkList.Insert(0,100)
	//fmt.Println("插入成功？",flag)
	//linkList.PrintInfo()
	//
	//fmt.Println()
	//fmt.Println("在下标2插入一个int 200")
	//flag=linkList.Insert(2,200)
	//fmt.Println("插入成功？",flag)
	//linkList.PrintInfo()
	//
	//fmt.Println()
	//fmt.Println("在下标5插入一个string sunlintong")
	//flag=linkList.Insert(5,"sunlintong")
	//fmt.Println("插入成功？",flag)
	//linkList.PrintInfo()
	//
	//fmt.Println()
	//fmt.Println("在下标7插入一个int 400")
	//flag=linkList.Insert(7,400)
	//fmt.Println("插入成功？",flag)
	//linkList.PrintInfo()
	//
	//fmt.Println()
	//fmt.Println("---------------------------现在开始测试删除数据-------------------------")
	//fmt.Println("删除 0")
	//linkList.Delete(0)
	//linkList.PrintInfo()
	//
	//fmt.Println()
	//fmt.Println("删除 string sunlintong")
	//linkList.Delete("sunlintong")
	//linkList.PrintInfo()
	//
	//fmt.Println()
	//fmt.Println("删除 500")
	//linkList.Delete(500)
	//linkList.PrintInfo()

	//数据库连接mysql
	//x, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8")
	//if err != nil {
	//	fmt.Println(err)
	//}


	////创建日志文件，路径可执行，同样可以使用封装好的框架："github.com/arthurkiller/rollingwriter"
	//f, err := os.Create("sql.log")
	//if err != nil {
	//	log.Fatalf("Fail to create log file: %v\n", err)
	//	return
	//}
	////日志文件进行logger实例创建
	//var logger *xorm.SimpleLogger = xorm.NewSimpleLogger(f)
	////设置logger
	//x.SetLogger(logger)
	////使用日志的方式记录sql语句
	//x.ShowSQL(true)
	////启动logger
	//logger.Info("test code begin")

	//分页查询，通过传递的页码，计算出偏移量，然后配合每页最多展示的数据量完成
	//err = x.Limit(3,1).Iterate(new(Account), func(idx int, bean interface{}) error {
	//	acc := bean.(*Account)
	//	fmt.Println(idx, "===>" , acc)
	//	return nil
	//})

	//只查询出表中所有数据的除name外的字段，那么字段默认为零值
	//err = x.Omit("name").Iterate(new(Account), func(idx int, bean interface{}) error {
	//	acc := bean.(*Account)
	//	fmt.Println(idx, "===>" , acc)
	//	return nil
	//})

	//只查询出表中所有数据的name字段，其他字段均为零值
	//err = x.Cols("name").Iterate(new(Account), func(idx int, bean interface{}) error {
	//	acc := bean.(*Account)
	//	fmt.Println(idx, "===>" , acc)
	//	return nil
	//})

	//获取到id大于5的数据，循环遍历
	//err = x.Where("id > ?",5).Iterate(new(Account), func(idx int, bean interface{}) error {
	//	acc := bean.(*Account)
	//	fmt.Println(idx, "===>" , acc)
	//	return nil
	//})

	////创建事务对象
	//seesion := x.NewSession()
	////事务结束后释放资源
	//defer seesion.Close()
	//
	////开启事务
	//if err = seesion.Begin(); err != nil {
	//	fmt.Println(err)
	//}
	////执行sql语句：完成业务逻辑
	//done, err := seesion.Insert(Account{Name:"madongmei2",Balance:8000})
	//if done == 0 || err != nil {
	//	//插入错误，调用事务回滚
	//	seesion.Rollback()
	//	fmt.Println(err)
	//}
	////插入成功，调用事务提交
	//seesion.Commit()



	//查询多条
	//all := make(map[int]Account,10)
	//err = x.Find(all,Account{Balance:4000})
	//if err != nil {
	//	fmt.Println("find data fail ,the error is :", err)
	//}
	////map 是无序的
	//for a, v := range all {
	//	fmt.Println(a,"===>",v)
	//}

	//创建表
	//if err = x.Sync2(new(Account)); err != nil {
	//	log.Fatalf("Fail to sync database: %v\n", err)
	//}

	//插入一条数据
	//_, err = x.Insert(&Account{Name:"jane5",Balance:1000})
	//if err != nil {
	//	log.Fatalf("Fail to insert test！")
	//}
	//根据传入的结构查询
	//has, err := x.Get(&Account{Id:1})
	//if has {
	//	fmt.Println("查询到ID为1的数据")
	//}
	//条件查询where
	//has, err = x.Where("name=? and balance=?","julia",1000).Get(&Account{})
	//if has {
	//	fmt.Println("查询到name balance所匹配到的数据")
	//}

	//查询表中数据量
	//a := new(Account)
	//int,err := x.Count(a)
	//fmt.Println("count = ",int)
	//if err != nil {
	//	fmt.Println("获取表中数据量：",err )
	//}

	//更新语句，传入的是指针
	//a := &Account{}
	//a.Balance = 6000
	//a.Version = 4
	//affected, err := x.Id(1).Update(a)
	//if affected == 0 && err != nil {
	//	fmt.Println("update fail ,the error is :", err)
	//	return
	//}








	//创建qos账户
	//filewrite.GenAccount()

	//defer是栈的形式，先进后出
	//routine.Defer_call()

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
	// o := reflection.NewOrder(123, 234)
	// reflection.CreateQuery(o)

	//同步调用
	//sync_test()

	// var Codec = amino.NewCodec()
	// privKey := ed25519.GenPrivKey()
	// priKeyBytes, err := Codec.MarshalJSON(privKey)

	// pubKey := privKey.PubKey()
	// pubKeyBytes, err := Codec.MarshalJSON(pubKey)

	// address := btypes.Address(pubKey.Address())

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("privKey = ", string(priKeyBytes))
	// fmt.Println("pubKey  = ", string(pubKeyBytes))
	// fmt.Println(address)

}
