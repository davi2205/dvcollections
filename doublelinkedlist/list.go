package doublelinkedlist

import "github.com/cheekybits/genny/generic"

type Generic generic.Type

type GenericListNode struct {
	Value      Generic
	list       *GenericList
	prev, next *GenericListNode
}

type GenericList struct {
	firstNode *GenericListNode
}

func (g *GenericList) Empty() bool {
	return g.firstNode == nil
}

func (g *GenericList) Last() *GenericListNode {
	if g.firstNode == nil {
		return nil
	}

	node := g.firstNode
	for ; node.next != nil; node = node.next {
	}

	return node
}

func (g *GenericList) Append(node *GenericListNode) {
	// TODO: finish this.
}
