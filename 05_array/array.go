package main

import (
	"fmt"
	"log"
	"strconv"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 */

// Array define array object
type Array struct {
	// 定义整型数据data保存数据
	data []int
	// 定义数组长度
	length int
	// 定义数组中实际个数
	count int
}

// New create a new array
func New(capacity int) *Array {
	return &Array{
		data:   make([]int, capacity, capacity),
		length: capacity,
		count:  0,
	}
}

// Find 根据索引，找到数据中的元素并返回
func (array *Array) Find(index int) int {
	if index < 0 || index >= array.count {
		return -1
	}

	return array.data[index]
}

func (array *Array) isFull() bool {
	return array.count == array.length
}

// Insert 插入元素:头部插入，尾部插入
func (array *Array) Insert(index, value int) bool {
	if array.isFull() {
		log.Println("array is full")
		return false
	}
	if index < 0 || index > array.count {
		log.Println("invalid index")
		return false
	}

	if index == array.count && array.count == 0 {
		array.data[index] = value
		array.count++
		return true
	}

	for i := array.count; i > index; i-- {
		array.data[i] = array.data[i-1]
	}
	array.data[index] = value
	array.count++
	return true
}

// Delete 根据索引，删除数组中元素
func (array *Array) Delete(index int) bool {
	if index < 0 || index >= array.count {
		log.Println("invalid index")
		return false
	}
	for i := index + 1; i < array.count; i++ {
		array.data[i-1] = array.data[i]
	}
	array.count--
	return false
}

func (array *Array) String() string {
	for i := 0; i < array.count; i++ {
		fmt.Print(strconv.Itoa(array.data[i]) + " ")
	}

	return ""
}

func main() {
	a := New(5)
	a.Insert(0, 1)
	a.Insert(0, 10)
	a.Insert(1, 1)
	a.Insert(2, 2)
	a.Insert(3, 3)
	a.Delete(3)
	a.Insert(3, 4)
	fmt.Println(a)
}
