package helloworld

import "fmt"

// func init() {
// 	fmt.Println("init run")
// }

func main_abandon() {
	fmt.Println("hello world")
	arr := []int{3, 4, 5, 6, 7, 3, 4, 2, 43, 32, 12, 56, 75, 3, 24, 8}
	// quickSort(arr, 0, len(arr)-1)
	// BubbleAsort(arr)
	fmt.Println(arr)

	// node1 := new(node)
	// node1.value = 1
	// node2 := new(node)
	// node2.value = 2
	// node3 := new(node)
	// node3.value = 3
	// node4 := new(node)
	// node4.value = 4
	// node1.nextNode = node2
	// node2.nextNode = node3
	// node3.nextNode = node4
	// printNode(node1)

	// head := reverseNode(node1)
	// printNode(head)
}

//选择排序
func SelectionSort(s []int) {
	l := len(s) //以免每次循环判断都运算
	m := len(s) - 1
	for i := 0; i < m; i++ {
		k := i
		for j := i + 1; j < l; j++ {
			if s[j] < s[k] {
				k = j
			}
		}
		if k != i {
			s[k], s[i] = s[i], s[k]
		}
	}
}

//插入排序
func InsertionSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

//二分查找函数 //假设有序数组的顺序是从小到大（很关键，决定左右方向）
func BinaryFind(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	//判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("没找到")
		return
	}

	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了，下标是%v\n", middle)
	}
}

// 翻转链表
type node struct {
	value    int
	nextNode *node
}

func ReverseNode(head *node) *node {
	//  先声明两个变量
	//  前一个节点
	var preNode *node
	preNode = nil
	//  后一个节点
	nextNode := new(node)
	nextNode = nil
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.nextNode
		//  将头节点指向前一个节点
		head.nextNode = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	return preNode
}

func printNode(head *node) {
	for head != nil {
		//fmt.Print(head.value, "\t")
		fmt.Println(head)
		head = head.nextNode
	}
	fmt.Println()
}

// 冒泡排序（升序）
func BubbleAsort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println(values)
}

// 冒泡排序(降序)
func BubbleZsort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println(values)
}

// 快速排序
func QuickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			QuickSort(arr, start, j)
		}
		if end > i {
			QuickSort(arr, i, end)
		}
	}
}
