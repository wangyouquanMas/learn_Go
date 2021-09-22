package gee

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

//功能: 匹配到子节点就返回,其中匹配：就看part是否匹配，或者子节点是否是模糊匹配
//匹配条件：part精确匹配 或者是模糊匹配
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//参数
//pattern：完整路径
//parts：按 /划分的字符串数组
//height: trie树高度
func (n *node) insert(pattern string, parts []string, height int) {
	//标识已经插入结束
	if len(parts) == height {
		//一个完整的节点的pattern 不为null
		// 如 :/hello/:name" 中 hello节点 pattern为"" , :name 节点pattern为 "/hello/:name"
		n.pattern = pattern
		return
	}

	//n : 当前节点
	part := parts[height]
	//child :当前节点n的下一个节点
	child := n.matchChild(part)
	//下一个节点不存在，则创建该节点，并加入当前节点的子节点切片中
	if child == nil {
		child = &node{part: part, isWild: parts[0] == ":" || parts[0] == "*"}
		n.children = append(n.children, child)
	}
	//当前节点变成子节点，进入下一层，进行递归插入
	child.insert(pattern, parts, height+1)
}

//返回所有匹配的子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//递归查找
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// pattern 为空字符串标识它不是一个完整的url,匹配失败
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		//是否会递归遍历其它树
		result := child.search(parts, height+1)
		//返回第一个找到的result
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		//一层一层递归查找pattern是非空的节点
		child.travel(list)
	}
}
