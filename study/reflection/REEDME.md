# reflection

## what is reflection

> 反射就是程序能够在运行时检查变量和值，求出它们的类型。

## why use reflection

> 首先我们会产生一个疑惑：在写程序的时候每个我们要用到的变量都是我们自己定义的，所以我们很清楚是什么类型的变量，为什么我们还需要在运行时检查变量，求出它的类型呢？假设我们需要完成一个函数createQuery，函数的入参时任何类型的参数：interface{}，我们需要根据传入的参数返回参数类型和值。

```go
func createQuery(q interface{}) string {
}

//当q为如下结构体o或e
o := order {
    ordId: 1234,
    customerId: 567
}

e := employee {
        name: "Naveen",
        id: 565,
        address: "Science Park Road, Singapore",
        salary: 90000,
        country: "Singapore",
}

//此时将o传入函数createQuery，我们期待返回结果
(1234, 567)
//此时将e传入函数createQuery，我们期待返回结果
("Naveen", 565, "Science Park Road, Singapore", 90000, "Singapore")

//其中包含string和int两种类型的字段，此时就需要使用反射来解析传入的参数及其地城具体类型和具体值。
```
