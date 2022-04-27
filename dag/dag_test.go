package dag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNode(t *testing.T) {

	family := New()

	tests := []struct {
		name     string
		nodeId   string
		nodeName string
		expErr   error
	}{
		{
			"Add Relation Valid Node",
			"node-1-id",
			"node-1-name",
			nil,
		},
		{
			"Add Relation Valid Node 2",
			"node-2-id",
			"node-2-name",
			nil,
		},
		{
			"Add Relation Invalid Node 2 (node already exists)",
			"node-2-id",
			"node-2-name",
			fmt.Errorf(NodeAlreadyExistsError),
		},
	}

	for _, tc := range tests {
		actualErr := family.AddNode(tc.nodeId, tc.nodeName)

		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
			assert.NotNil(t, family.nodes[tc.nodeId], tc.name)
			assert.Equal(t, family.nodes[tc.nodeId].GetId(), tc.nodeId, tc.name)
			assert.Equal(t, family.nodes[tc.nodeId].GetName(), tc.nodeName, tc.name)
		} else {
			assert.ErrorContainsf(t, actualErr, tc.expErr.Error(), tc.name)
		}
	}
}

func TestGetParents(t *testing.T) {

	family := New()

	node1id, node1name := "N1-id", "N1-name"
	node2id, node2name := "N2-id", "N2-name"
	node3id, node3name := "N3-id", "N3-name"
	node4id, node4name := "N4-id", "N4-name"
	node5id, node5name := "N5-id", "N5-name"

	family.AddNode(node1id, node1name)
	family.AddNode(node2id, node2name)
	family.AddNode(node3id, node3name)
	family.AddNode(node4id, node4name)
	family.AddNode(node5id, node5name)

	// 4    5
	//   \ /
	//    2    3
	//      \ /
	//       1

	family.AddRelation(node2id, node1id)
	family.AddRelation(node3id, node1id)
	family.AddRelation(node4id, node2id)
	family.AddRelation(node5id, node2id)

	tests := []struct {
		name    string
		nodeId  string
		parents map[string]*Node
		expErr  error
	}{
		{
			"Get Parents Valid Node",
			node1id,
			map[string]*Node{
				node2id: family.nodes[node2id],
				node3id: family.nodes[node3id],
			},
			nil,
		},
		{
			"Get Parents valid Node - (root node)",
			node5id,
			map[string]*Node{},
			nil,
		},
		{
			"Get Parents Invalid Node",
			"node9id",
			map[string]*Node{},
			fmt.Errorf(NodeNotFound),
		},
	}

	for _, tc := range tests {
		actualParents, actualErr := family.GetParents(tc.nodeId)

		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
			assert.Equal(t, actualParents, tc.parents, tc.name)
		} else {
			assert.ErrorContainsf(t, actualErr, tc.expErr.Error(), tc.name)
		}
	}
}

func TestGetChildren(t *testing.T) {

	family := New()

	node1id, node1name := "N1-id", "N1-name"
	node2id, node2name := "N2-id", "N2-name"
	node3id, node3name := "N3-id", "N3-name"
	node4id, node4name := "N4-id", "N4-name"
	node5id, node5name := "N5-id", "N5-name"

	family.AddNode(node1id, node1name)
	family.AddNode(node2id, node2name)
	family.AddNode(node3id, node3name)
	family.AddNode(node4id, node4name)
	family.AddNode(node5id, node5name)

	// 4    5
	//   \ /
	//    2    3
	//      \ /
	//       1

	family.AddRelation(node2id, node1id)
	family.AddRelation(node3id, node1id)
	family.AddRelation(node4id, node2id)
	family.AddRelation(node5id, node2id)

	tests := []struct {
		name     string
		nodeId   string
		children map[string]*Node
		expErr   error
	}{
		{
			"Get Children Valid Node",
			node4id,
			map[string]*Node{
				node2id: family.nodes[node2id],
			},
			nil,
		},
		{
			"Get Children valid Node - (leaf node)",
			node1id,
			map[string]*Node{},
			nil,
		},
		{
			"Get Children Invalid Node",
			"node9id",
			map[string]*Node{},
			fmt.Errorf(NodeNotFound),
		},
	}

	for _, tc := range tests {
		actualParents, actualErr := family.GetChildren(tc.nodeId)

		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
			assert.Equal(t, actualParents, tc.children, tc.name)
		} else {
			assert.ErrorContainsf(t, actualErr, tc.expErr.Error(), tc.name)
		}
	}
}

func TestGetAncestors(t *testing.T) {

	family := New()

	node1id, node1name := "N1-id", "N1-name"
	node2id, node2name := "N2-id", "N2-name"
	node3id, node3name := "N3-id", "N3-name"
	node4id, node4name := "N4-id", "N4-name"
	node5id, node5name := "N5-id", "N5-name"

	family.AddNode(node1id, node1name)
	family.AddNode(node2id, node2name)
	family.AddNode(node3id, node3name)
	family.AddNode(node4id, node4name)
	family.AddNode(node5id, node5name)

	// 4    5
	//   \ /
	//    2    3
	//      \ /
	//       1

	family.AddRelation(node2id, node1id)
	family.AddRelation(node3id, node1id)
	family.AddRelation(node4id, node2id)
	family.AddRelation(node5id, node2id)

	tests := []struct {
		name      string
		nodeId    string
		ancestors map[string]*Node
		expErr    error
	}{
		{
			"Get Ancestors Valid Node - Test Multilevel",
			node1id,
			map[string]*Node{
				node2id: family.nodes[node2id],
				node3id: family.nodes[node3id],
				node4id: family.nodes[node4id],
				node5id: family.nodes[node5id],
			},
			nil,
		},
		{
			"Get Ancestors valid Node - One Level Parents",
			node2id,
			map[string]*Node{
				node4id: family.nodes[node4id],
				node5id: family.nodes[node5id],
			},
			nil,
		},
		{
			"Get Ancestors valid Node - root node",
			node3id,
			map[string]*Node{},
			nil,
		},
		{
			"Get Ancestors Invalid Node",
			"node9id",
			map[string]*Node{},
			fmt.Errorf(NodeNotFound),
		},
	}

	for _, tc := range tests {
		actualAncestors, actualErr := family.GetAncestors(tc.nodeId)

		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
			assert.Equal(t, actualAncestors, tc.ancestors, tc.name)
		} else {
			assert.ErrorContainsf(t, actualErr, tc.expErr.Error(), tc.name)
		}
	}
}

func TestGetDescendents(t *testing.T) {

	family := New()

	node1id, node1name := "N1-id", "N1-name"
	node2id, node2name := "N2-id", "N2-name"
	node3id, node3name := "N3-id", "N3-name"
	node4id, node4name := "N4-id", "N4-name"
	node5id, node5name := "N5-id", "N5-name"

	family.AddNode(node1id, node1name)
	family.AddNode(node2id, node2name)
	family.AddNode(node3id, node3name)
	family.AddNode(node4id, node4name)
	family.AddNode(node5id, node5name)

	// 4    5
	//   \ /
	//    2    3
	//      \ /
	//       1

	family.AddRelation(node2id, node1id)
	family.AddRelation(node3id, node1id)
	family.AddRelation(node4id, node2id)
	family.AddRelation(node5id, node2id)

	tests := []struct {
		name        string
		nodeId      string
		descendents map[string]*Node
		expErr      error
	}{
		{
			"Get Descendents Valid Node - Test Multilevel",
			node5id,
			map[string]*Node{
				node1id: family.nodes[node1id],
				node2id: family.nodes[node2id],
			},
			nil,
		},
		{
			"Get Descendents valid Node - One Level Parents",
			node2id,
			map[string]*Node{
				node1id: family.nodes[node1id],
			},
			nil,
		},
		{
			"Get Descendents valid Node - leaf node",
			node1id,
			map[string]*Node{},
			nil,
		},
		{
			"Get Descendents Invalid Node",
			"node9id",
			map[string]*Node{},
			fmt.Errorf(NodeNotFound),
		},
	}

	for _, tc := range tests {
		actualDescendents, actualErr := family.GetDescendents(tc.nodeId)

		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
			assert.Equal(t, actualDescendents, tc.descendents, tc.name)
		} else {
			assert.ErrorContainsf(t, actualErr, tc.expErr.Error(), tc.name)
		}
	}
}

func TestAddRelation(t *testing.T) {

	family := New()

	node1name := "P1"
	node2name := "C2"

	family.AddNode(node1name, "PN1")
	family.AddNode(node2name, "CN2")

	node1 := family.nodes[node1name]
	node2 := family.nodes[node2name]

	tests := []struct {
		name     string
		nodeId   string
		parentId string
		expErr   error
	}{
		{
			"Add Relation Valid Case",
			node2name,
			node1name,
			nil,
		},
		{
			"Add Relation - Invalid Case - Cyclic Dependency (Dependent on previous testcase",
			node1name,
			node2name,
			fmt.Errorf(CyclicDependencyError),
		},
	}

	for _, tc := range tests {
		actualErr := family.AddRelation(tc.parentId, tc.nodeId)

		fmt.Println(node1, "--", node2)

		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
			assert.Equal(t, node1, node2.parents[node1name], tc.name)
			assert.Equal(t, node1.children[node2name], node2, tc.name)
		} else {
			fmt.Println(actualErr, "act")
			assert.ErrorContainsf(t, actualErr, tc.expErr.Error(), tc.name)
		}
	}
}

func TestDeleteRelation(t *testing.T) {

	family := New()

	node1id, node1name := "N1-id", "N1-name"
	node2id, node2name := "N2-id", "N2-name"
	node3id, node3name := "N3-id", "N3-name"
	node4id, node4name := "N4-id", "N4-name"
	node5id, node5name := "N5-id", "N5-name"

	family.AddNode(node1id, node1name)
	family.AddNode(node2id, node2name)
	family.AddNode(node3id, node3name)
	family.AddNode(node4id, node4name)
	family.AddNode(node5id, node5name)

	// 4    5
	//   \ /
	//    2    3
	//      \ /
	//       1

	family.AddRelation(node2id, node1id)
	family.AddRelation(node3id, node1id)
	family.AddRelation(node4id, node2id)
	family.AddRelation(node5id, node2id)

	tests := []struct {
		name    string
		node1Id string
		node2Id string
		expErr  error
	}{
		{
			"Delete Valid Relation",
			node2id,
			node1id,
			nil,
		},
		{
			"Delete Invalid Relation - (un connected nodes)",
			node5id,
			node1id,
			fmt.Errorf(NodeRelationNotFound),
		},
	}

	for _, tc := range tests {
		actualErr := family.DeleteRelation(tc.node1Id, tc.node2Id)

		if tc.expErr == nil {
			assert.Equal(t, tc.expErr, actualErr, tc.name)
		} else {
			assert.ErrorContainsf(t, actualErr, tc.expErr.Error(), tc.name)
		}
	}
}

func TestDelete(t *testing.T) {

	family := New()

	node1id, node1name := "N1-id", "N1-name"
	node2id, node2name := "N2-id", "N2-name"
	node3id, node3name := "N3-id", "N3-name"
	node4id, node4name := "N4-id", "N4-name"
	node5id, node5name := "N5-id", "N5-name"

	family.AddNode(node1id, node1name)
	family.AddNode(node2id, node2name)
	family.AddNode(node3id, node3name)
	family.AddNode(node4id, node4name)
	family.AddNode(node5id, node5name)

	// 4    5
	//   \ /
	//    2    3
	//      \ /
	//       1

	family.AddRelation(node2id, node1id)
	family.AddRelation(node3id, node1id)
	family.AddRelation(node4id, node2id)
	family.AddRelation(node5id, node2id)

	tests := []struct {
		name           string
		nodeIdToDelete string
		expErr         error
	}{
		{
			"Delete Valid Relation",
			node2id,
			nil,
		},
	}

	for _, tc := range tests {
		actualErr := family.delete(tc.nodeIdToDelete)

		if tc.expErr == nil {
			assert.Equal(t, tc.expErr, actualErr, tc.name)
		}
	}
}
