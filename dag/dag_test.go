package dag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRelation(t *testing.T) {

	family := New()

	basenodeId := "basenodeId"
	childAId := "id-A"

	family.AddNode(basenodeId, "simple-name")
	family.AddNode(childAId, "child-A")

	baseNode := family.nodes[basenodeId]
	childA := family.nodes[childAId]

	tests := []struct {
		name     string
		nodeId   string
		parentId string
	}{
		{
			"Add child to node",
			childAId,
			basenodeId,
		},
	}

	for _, tc := range tests {
		family.AddRelation(tc.parentId, tc.nodeId)
		assert.Equal(t, baseNode, childA.parents[basenodeId], tc.name)
		// assert.Equal(t, &tc.expectedNode, tc.node.parents[sampleNode.GetId()], tc.name)
	}
}

/*
func TestDeleteRelation(t *testing.T) {

	sampleNode := NewNode("basenodeId", "simple-name")
	node1 := NewNode("node-1", "name-1")
	node2 := NewNode("node-2", "name-2")

	AddRelation(&node1, &sampleNode)
	AddRelation(&sampleNode, &node2)

	var empty *Node

	tests := []struct {
		name         string
		node         Node
		expectedNode Node
	}{
		{
			"Remove relation",
			node1,
			sampleNode,
		},
	}

	for _, tc := range tests {
		DeleteRelation(&sampleNode, &tc.node)
		assert.Equal(t, tc.expectedNode.children[tc.node.GetId()], empty, tc.name)
	}
}

func TestGetAncestors(t *testing.T) {
	node1 := NewNode("node-1", "name-1")
	node2 := NewNode("node-2", "name-2")
	node3 := NewNode("node-3", "name-3")
	node4 := NewNode("node-4", "name-4")
	node5 := NewNode("node-5", "name-5")
	node6 := NewNode("node-6", "name-6")

	AddRelation(&node2, &node1)
	AddRelation(&node3, &node1)
	AddRelation(&node4, &node2)
	AddRelation(&node5, &node2)
	AddRelation(&node6, &node3)

	var empty []Node

		// 4   5     6
		//  \ /     /
		//   2     3
		//    \   /
		//      1

	tests := []struct {
		name          string
		node          Node
		expAnscestors []Node
	}{
		{
			"Multilevel ancestors",
			node1,
			[]Node{
				node2,
				node3,
				node4,
				node5,
				node6,
			},
		},
		{
			"Singlelevel ancestors",
			node3,
			[]Node{
				node6,
			},
		},
		{
			"No ancestors",
			node5,
			empty,
		},
	}

	for _, tc := range tests {
		actualAncestors := tc.node.getAncestors()
		sort.Slice(actualAncestors, func(i, j int) bool { return actualAncestors[i].GetId() > actualAncestors[j].GetId() })
		sort.Slice(tc.expAnscestors, func(i, j int) bool { return tc.expAnscestors[i].GetId() > tc.expAnscestors[j].GetId() })

		assert.Equal(t, actualAncestors, tc.expAnscestors, tc.name)
	}
}

func TestGetDescendents(t *testing.T) {
	node1 := NewNode("node-1", "name-1")
	node2 := NewNode("node-2", "name-2")
	node3 := NewNode("node-3", "name-3")
	node4 := NewNode("node-4", "name-4")
	node5 := NewNode("node-5", "name-5")
	node6 := NewNode("node-6", "name-6")

	AddRelation(&node2, &node1)
	AddRelation(&node3, &node1)
	AddRelation(&node4, &node2)
	AddRelation(&node5, &node2)
	AddRelation(&node6, &node3)

	var empty []Node

		// 4   5     6
		//  \ /     /
		//   2     3
		//    \   /
		//      1

	tests := []struct {
		name           string
		node           Node
		expDescendents []Node
	}{
		{
			"No descendents",
			node1,
			empty,
		},
		{
			"Singlelevel descendents",
			node2,
			[]Node{
				node1,
			},
		},
		{
			"Multilevel descendents",
			node5,
			[]Node{
				node2,
				node1,
			},
		},
	}

	for _, tc := range tests {
		actualDescendents := tc.node.getDescendents()
		sort.Slice(actualDescendents, func(i, j int) bool { return actualDescendents[i].GetId() > actualDescendents[j].GetId() })
		sort.Slice(tc.expDescendents, func(i, j int) bool { return tc.expDescendents[i].GetId() > tc.expDescendents[j].GetId() })

		assert.Equal(t, actualDescendents, tc.expDescendents, tc.name)
	}
}

*/
