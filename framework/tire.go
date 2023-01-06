package framework

import (
	"errors"
	"strings"
)

type node struct {
	isLast  bool              // 代表这个节点是否可以成为最终的路由规则，该节点是否能成为一个独立的uri，是否自身就是一个终极节点
	segment string            // uri中的字符串，代表这个节点表示的路由中某个段的字符串
	handler ControllerHandler // 代表这个节点对应的控制器
	childs  []*node           // 代表这个节点下的子节点

	handlers []ControllerHandler
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

// Tree 代表树结构
type Tree struct {
	root *node
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root: root}
}

//判断一个segment是否是通用segment，即以：开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

//过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	// 如果segment是通配符，则所有下一层子节点都满足要求
	if isWildSegment(segment) {
		// /home/log/:id/name
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	// 过滤所有的下一层子节点
	for _, cnode := range n.childs {
		// 某一个节点满足条件的子节点就2种（后续好好验证）
		if isWildSegment(cnode.segment) {
			// 如果下一层子节点有通配符，则满足要求
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			// 如果下一层子节点，没有通配符，但是文本完全满足，则满足要求
			nodes = append(nodes, cnode)
		}
	}

	return nodes
}

// 判断路由是否已经在节点的所有子节点树中存在了
func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)

	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	cnodes := n.filterChildNodes(segment)

	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}

	// 如果只有一个segment，则是最后一个标记
	if len(segments) == 1 {
		// 如果segment已经是最后一个节点，判断这些cnode是否有isLast标志
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}

		// 一个都没有找到，说明这个是新的，需要加入树中
		return nil
	}

	// 如果有2个segment，递归每一个子节点继续进行查找
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

// 增加路由节点, 路由节点有先后顺序
/*
/book/list
/book/:id (冲突)
/book/:id/name
/book/:student/age
/:user/name
/:user/name/:age (冲突)
*/
func (tree *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	n := tree.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist:" + uri)
	}

	segments := strings.Split(uri, "/")
	// 对每个segment
	for index, segment := range segments {

		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点

		childNodes := n.filterChildNodes(segment)
		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				objNode = cnode
				break
			}
		}

		if objNode == nil {
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handlers = handlers
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}

		n = objNode
	}

	return nil
}

// FindHandler 匹配uri
func (tree *Tree) FindHandler(uri string) []ControllerHandler {
	// 直接复用matchNode函数
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handlers
}
