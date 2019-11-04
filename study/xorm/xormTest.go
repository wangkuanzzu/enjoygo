package xorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` // 乐观锁
}

var X *xorm.Engine

func (a *Account) BeforeInsert() {
	fmt.Printf("before insert: %s", a.Name)
}

func (a *Account) AfterInsert() {
	fmt.Printf("after insert: %s", a.Name)
}
func UseXorm(){
	//数据库连接mysql
	x, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}


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
	err = x.Limit(3,1).Iterate(new(Account), func(idx int, bean interface{}) error {
		acc := bean.(*Account)
		fmt.Println(idx, "===>" , acc)
		return nil
	})

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
}
