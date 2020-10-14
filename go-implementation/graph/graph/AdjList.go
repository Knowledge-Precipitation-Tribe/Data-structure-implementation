package graph

/**
* @Author: super
* @Date: 2020-10-14 14:50
* @Description:
**/

type AdjNode struct {
	Data interface{}
	Next *AdjNode
}

type AdjList struct {
	List []*AdjNode
}
