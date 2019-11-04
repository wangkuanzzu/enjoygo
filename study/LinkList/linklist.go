package LinkList

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkList struct {
	head *Node
	tail *Node
	size int
}

func CreateLinkList() LinkList{
	l := LinkList{}
	l.head = nil
	l.head = nil
	l.size = 0
	return l
}

func (l *LinkList) IsEmpty() bool{
	return l.size == 0
}

func (l *LinkList) GetLength() int{
	return l.size
}

func (l *LinkList) Exist(node *Node) bool {
	var p *Node = l.head
	for p != nil {
		if p == node {
			return true
		}else{
			p = p.next
		}
	}
	return false
}

func (l *LinkList) GetNode(e interface{}) *Node{
	var p *Node = l.head
	for p != nil {
		if e == p.data {
			return p
		}else {
			p = p.next
		}
	}
	return nil
}

func (l *LinkList) Append(data interface{}) {
	//根据data创建一个新的Node结构体
	node := Node{}
	node.data = data
	node.next = nil

	//判断链表是否为空
	if l.size == 0 {
		l.head = &node
		l.tail = &node
	}else{
		//获取到当前链表的最后一个元素，该元素的指针原来指向nil，现在指向新的node元素
		l.tail.next = &node
		//链表的最后一个元素更新为新的node元素
		l.tail = &node
	}
	//链表的大小增加1
	l.size++
}

func (l *LinkList) InsertHead(data interface{}) {
	//根据data创建一个新的Node结构体
	node := Node{}
	node.data = data
	node.next = l.head

	l.head = &node
	if l.size == 0 {
		l.tail = &node
	}
	l.size++

}

//在指定的节点pre后面插入data为e的新节点
func (l *LinkList) InsertAfterNode(pre *Node, e interface{}){

	//指定的节点必须在l中存在
	if l.Exist(pre) {
		newnode := Node{}
		newnode.data = e

		if pre.next == nil {
			l.Append(e)
		}else{
			newnode.next = pre.next
			pre.next = &newnode
		}
		l.size++
		return
	}
	fmt.Println("链表中不存在该节点")
	return
}

//在第一次出现指定数据的节点后面增加新节点，仅限第一次
func (l *LinkList) InsertAfterData(preData interface{}, e interface{}) bool {
	var p *Node = l.head
	for p != nil {
		if p.data == preData {
			//找到指定数据的节点后，调用指定节点后插入数据的方法
			l.InsertAfterNode(p, e)
			return true
		}else{
			p = p.next
		}

	}
	fmt.Println("链表中没有查询到指定数据，插入失败！")
	return false
}

//在指定下标中插入数据，参数：position 位置，e 插入数据
func (l *LinkList) Insert(position int, e interface{}) bool {
	if position < 0 {
		fmt.Println("参数不合法，插入失败！")
		return false
	}else if position > l.size {
		fmt.Println("指定下标超出链表长度，插入失败！")
		return false
	}else if position == 0 {
		l.InsertHead(e)
		return true
	}else if position == l.size {
		l.Append(e)
		return true
	}else{
		var index int
		var p *Node = l.head
		//position是插入后新结点的下标，position-1时需要定位到的其前一个结点的下标
		for index = 0; index < position-1; index++ {
			p = p.next
		}
		l.InsertAfterNode(p, e)
		return true
	}
}

//删除指定节点
func (l *LinkList) DeleteNode(node *Node){
	//判断节点是否存在
	if l.Exist(node) {
		if l.head == node {
			//头部节点
			l.head.next = node
		}else if l.tail == node{
			//尾部节点
			var p *Node = l.head
			for p.next != l.tail {
				p = p.next
			}
			p.next = nil
			l.tail = p
		}else {
			//中间节点
			var p *Node = l.head
			//找到指向node的节点p
			for p.next != node {
				p = p.next
			}
			p.next = node.next
		}
		l.size--
	}

}

//删除指定数据的节点
func (l *LinkList) Delete (e interface{}) {
	node := l.GetNode(e)
	if node != nil {
		l.DeleteNode(node)
	}else {
		fmt.Println("链表中没有该数据，删除失败！")
	}
}

//遍历链表
func (l *LinkList) Traverse(){
	if l.IsEmpty() {
		fmt.Println("链表不能为空")
	}else{
		p := l.head
		for p !=nil {
			fmt.Print(p.data, " -> ")
			p = p.next
		}
		fmt.Println()
	}
}

//打印链表信息
func(l *LinkList) PrintInfo(){
	fmt.Println("###############################################")
	fmt.Println("链表长度为：",l.GetLength())
	fmt.Println("链表是否为空:",l.IsEmpty())
	fmt.Print("遍历链表：")
	l.Traverse()
	fmt.Println("###############################################")
}




