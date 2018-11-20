package main

import (
	"fmt"
)

//链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表的实现
func reversrList(head *ListNode) *ListNode {
	// cur := head
	// var pre *ListNode = nil
	// for cur != nil {
	// 	log.Error(cur)
	// 	log.Error(cur.Next)
	// 	log.Error(pre)
	// 	pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	// }
	// return pre

	var newHead *ListNode = nil
	next := head
	for next != nil {
		tmp := newHead // nil --- {1, nil}

		newHead = next // {1, 0xc42008a5} --- {2, 0xc42008a5}

		next = next.Next // {2, 0xc42008a5} --- {3, 0xc42008a5}

		newHead.Next = tmp // nil => newHead = {1, nil} --- {1, nil} => newHead = {2, {1,nil}}....
	}
	return newHead
}

func main() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	pre := reversrList(head)
	fmt.Println(pre)
}
