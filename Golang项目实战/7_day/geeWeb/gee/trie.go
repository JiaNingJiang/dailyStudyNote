package gee

import "strings"

type node struct {
	pattern  string  // 当前节点作为某一注册路由的根节点，保存该路由的全路径。如：/user/userInfo/register
	part     string  // 当前节点保存的路由的对应部分。如user、userInfo或者register (根节点的part为空)
	children []*node // 当前节点的子节点
	isWild   bool    // 是否允许模糊匹配，若设置为true，则当前节点的part需要以`:`或者`*`开头，用于匹配所有的part
}

// 返回当前节点第一个满足匹配的子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild { //子节点匹配成功，或者子节点允许被模糊匹配
			return child
		}
	}
	return nil
}

// 返回当前节点所有满足匹配的子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild { //子节点匹配成功，或者子节点允许被模糊匹配
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 插入若干新节点(pattern指定的路径涉及到的节点都会进行插入) --- 只有路径的叶子结点的 n.pattern = pattern
// pattern: 路由的全路径
// parts: 路由全路径包含的所有part(/home/ShanDong/QingDao的parts = [home ShanDong QingDao]);但如果路径中包含以*开始的part,则仅会到该part前一个位置(/home/*/QingDao的parts = [home ])
// height: 当前的插入深度，总是从0开始(由根节点开始)，随着递归逐步递增
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height { // 当前节点就是此路由的叶子结点，保存pattern
		n.pattern = pattern
		return
	}

	part := parts[height]       // 获取目标子节点的part
	child := n.matchChild(part) // 返回第一个匹配的子节点
	if child == nil {           // 如果不存在任何子节点，则为其创建一个子节点(如果该子节点的part支持模糊匹配,则isWild = true)
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1) // 递归插入
}

// 返回parts组成的完整路由路径的叶子结点 或者 中途遇到包含通配符*的叶子结点(对于以*开头的part,必定在insert时作为叶子结点存储n.pattern)
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	// if children == nil {
	// 	self := &node{
	// 		pattern:  n.pattern,
	// 		part:     n.part,
	// 		children: n.children,
	// 		isWild:   n.isWild,
	// 	}
	// 	return self
	// }

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
