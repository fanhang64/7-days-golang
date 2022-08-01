package gee

import (
	"strings"
)

type node struct {
	pattern  string  // 待匹配的路由  /p/:lang
	part     string  // 路由的部分 :lang
	children []*node // 子节点
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// matchChild 第一次匹配成功，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 匹配所有，用于查找, 只查找n.children，而不会查找child的children
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0, 10)

	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 递归插入
// height 层高
// 插入功能，递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个，
// 有一点需要注意，/p/:lang/doc只有在第三层节点，即doc节点，pattern才会设置为/p/:lang/doc
func (n *node) insert(pattern string, parts []string, height int) { // /hello/:name
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search 递归查找结点
func (n *node) search(parts []string, height int) *node { // assets abc   0
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
