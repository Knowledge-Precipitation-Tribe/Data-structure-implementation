package graph

/**
* @Author: super
* @Date: 2020-10-14 14:50
* @Description:
**/

// 邻接表中的具体链方式采用链表实现
type AdjNode struct {
	Data interface{}
	Next *AdjNode
}

// 邻接表的定义
type AdjList struct {
	List []*AdjNode
}

