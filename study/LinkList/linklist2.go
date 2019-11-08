package LinkList

import (
	"fmt"
	"sync"
)

//定义节点存储的数据类型
type DoubleObject interface {}

//定义双向链表的节点
type DoubleNode struct {
	Data DoubleObject
	Prev *DoubleNode
	Next *DoubleNode
}

//定义双向链表
type DoubleList struct {
	mutex *sync.RWMutex
	Size uint
	Head *DoubleNode
	Tail *DoubleNode
}
//链表初始化
func (l *DoubleList) Init() {
	l.mutex = new(sync.RWMutex)
	l.Size = 0
	l.Head = nil
	l.Tail = nil
}

//定义节点的基本方法
func (node *DoubleNode) GetData() interface{} {
	return node.Data
}

//根据指定索引获取链表节点,提及索引就是从0开始
func (l *DoubleList) GetNode(index uint) *DoubleNode {
	if l.Size == 0 || index > l.Size - 1{
		return nil
	}
	if index == 0 {
		return l.Head
	}
	node := l.Head
	var i uint
	for i = 1; i <= index; i++ {
		node = node.Next
	}
	return node
}

//链表节点的新增分为两种，一种是在链表后面追加节点，该方式，我们称为append；另外一种方式是在指定位置插入节点，我们叫做insert。
func (l *DoubleList) Append(node *DoubleNode) bool {
	if node == nil {
		return false
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.Size == 0 {
		l.Head = node
		l.Tail = node
		node.Next = nil
		node.Prev = nil
	}else {

		node.Prev = l.Tail
		node.Next = nil
		l.Tail.Next = node
		l.Tail = node
	}
	l.Size++
	return true
}

func (l *DoubleList) Insert(index uint, node *DoubleNode) bool {
	if index > l.Size || node == nil {
		return false
	}

	if l.Size == index {
		return l.Append(node)
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()
	if index == 0 {
		node.Prev = nil
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}else {
		indexNode := l.GetNode(index)
		node.Next = indexNode
		node.Prev = indexNode.Prev
		indexNode.Prev.Next = node
		indexNode.Prev = node
	}
	l.Size++
	return true
}

func (l *DoubleList) Delete(index uint) bool {

	if index > l.Size-1 {
		return false
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	//删除的元素为第一个，index == 0
	//  1.链表的长度为1
	//  2.链表的长度大于1
	if index == 0 {
		if l.Size == 1 {
			l.Head = nil
			l.Tail = nil
		}else {
			l.Head.Next.Prev = nil
			l.Head = l.Head.Next
		}
		l.Size--
		return true
	}
	// 删除的元素为最后一个，index == l.Size-1
	if index == l.Size-1 {
		l.Tail.Prev.Next = nil
		l.Tail = l.Tail.Prev
		l.Size--
		return true
	}

	//删除的元素为中间
	indexNode := l.GetNode(index)
	indexNode.Prev.Next = indexNode.Next
	indexNode.Next.Prev = indexNode.Prev
	l.Size--
	return true

}

// 正向打印链表结构
func (l *DoubleList) Display(){
	if l == nil || l.Size == 0 {
		fmt.Println("the doublelist is empty !")
		return
	}
	l.mutex.Lock()
	defer l.mutex.Unlock()

	fmt.Println("the size of the double list is : ", l.Size)
	p := l.Head
	for p != nil {
		fmt.Printf("data is %v\n", p.Data)
		p = p.Next
	}

}

