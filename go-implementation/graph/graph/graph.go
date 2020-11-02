package graph

/**
* @Author: super
* @Date: 2020-10-14 14:47
* @Description:
**/

type Graph struct {
	AdjList // 邻接表
	N int //节点个数
}

func NewGraph() Graph{
	return Graph{
		AdjList{
			List:make([]*AdjNode, 0),
		},
		0,
	}
}

func (g *Graph)AddNode(e interface{}){
	
}