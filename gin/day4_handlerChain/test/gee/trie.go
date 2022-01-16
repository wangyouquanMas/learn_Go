package gee

import (
	"strings"
)

type Node struct {
	part        string
	parts       []string
	path        string
	children    []*Node
	isWildMatch bool
}

func NewNode(p string, isWild bool) *Node {
	return &Node{
		part:        p,
		isWildMatch: isWild,
	}
}

//查询树是否存在某个路由
//注意通配符情况
func (node *Node) checkTree(parts []string, height int) (res *Node) {
	//相等是删选条件1
	if len(parts) == height || strings.HasPrefix(node.part, "*") {
		//删选条件2
		if node.path == "" {
			return
		}
		return node
	}
	children := node.allMathchChildren(parts[height])
	if children == nil {
		return
	}

	for _, child := range children {
		result := child.checkTree(parts, height+1)
		if result != nil {
			return result
		}
	}
	return
}

//建树
func (node *Node) buildTree(parts []string, path string, height int) {
	if len(parts) == height {
		node.path = path
		return
	}
	child := node.existInChildren(parts[height])
	if child == nil {
		child = NewNode(parts[height], parts[height][0] == '*' || parts[height][0] == ':')
		node.children = append(node.children, child)
	}
	child.buildTree(parts, path, height+1)
}

//待插入节点是否出现在当前节点的子节点中
func (node *Node) existInChildren(part string) (child *Node) {
	for _, v := range node.children {
		if v.part == part || v.isWildMatch {
			return v
		}
	}
	return
}

//如果子节点中有和待查询一致的节点，那么遍历所有子节点
func (node *Node) allMathchChildren(part string) (children []*Node) {
	children = make([]*Node, 0)
	for _, v := range node.children {
		if v.part == part || v.isWildMatch {
			children = append(children, v)
		}
	}
	return children
}
