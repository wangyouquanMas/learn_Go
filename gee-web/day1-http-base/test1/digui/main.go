package digui

import "strings"

type node struct {
	path     string
	children []*node
	part     string
	isWild   bool
}

func (n *node) checkPath(path string, parts []string, height int) (res bool) {
	//结束判断
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.path != "" {
			return true
		}
	}
	//遍历递归查询
	//查询当前节点的子节点中是否有和待查询节点匹配的，如果有，都需要进行继续校验
	//找到一个匹配的，直接退出查询
	children := n.isExistInChildren(parts[height])
	for _, child := range children {
		result := child.checkPath(path, parts, height+1)
		if result {
			return true
		}
	}
	return
}

func (n *node) isExistInChildren(part string) (res []*node) {
	res = make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			res = append(res, child)
		}
	}
	return res
}

func (n *node) registerPath(path string, parts []string, height int) {
	//结束判断
	if len(parts) == height {
		n.path = path
		return
	}
	child := &node{part: parts[height], isWild: parts[height] == ":" || parts[height] == "*"}
	//待添加节点值是否已经存在
	if isExist := n.isExist(parts[height]); isExist == false {
		//如果，不存在则将该子节点加入当前节点的子节点数组中
		n.children = append(n.children, child)
	}
	//从该节点递归注册
	child.registerPath(path, parts, height+1)
}
func (n *node) isExist(part string) (res bool) {
	for _, child := range n.children {
		if child.part == part || child.part == "*" || child.part == ":" {
			return true
		}
	}
	return false
}
