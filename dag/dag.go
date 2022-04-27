package dag

import (
	"fmt"
)

type dag struct {
	nodes map[string]*Node
}

const (
	NodeAlreadyExistsError = "error : node already exists"
	NodeNotFound           = "error : node not found"
	NodeRelationNotFound   = "error : node relation not present"
	CyclicDependencyError  = "error : can not create cyclic dependency"
)

func New() *dag {
	d := dag{}
	d.nodes = make(map[string]*Node)
	return &d
}

func (d *dag) AddNode(id string, name string) error {
	if _, exists := d.nodes[id]; exists {
		return fmt.Errorf(NodeAlreadyExistsError)
	}

	node := NewNode(id, name)
	if err := node.validate(); err != nil {
		return err
	}

	d.nodes[id] = &node
	return nil
}

func (d *dag) GetParents(id string) (map[string]*Node, error) {
	if _, exists := d.nodes[id]; !exists {
		return nil, fmt.Errorf(NodeNotFound)
	}

	return d.nodes[id].GetParents(), nil
}

func (d *dag) GetChildren(id string) (map[string]*Node, error) {
	if _, exists := d.nodes[id]; !exists {
		return nil, fmt.Errorf(NodeNotFound)
	}

	return d.nodes[id].GetChildren(), nil
}

func (d *dag) GetAncestors(id string) (map[string]*Node, error) {
	if _, exists := d.nodes[id]; !exists {
		return nil, fmt.Errorf(NodeNotFound)
	}

	var ancestors map[string]*Node = make(map[string]*Node)
	for _, parentNode := range d.nodes[id].parents {
		ancestors[parentNode.GetId()] = parentNode

		grandParents, err := d.GetAncestors(parentNode.GetId())
		if len(grandParents) == 0 {
			continue
		}

		if err != nil {
			return ancestors, err
		}

		for _, grandParent := range grandParents {
			ancestors[grandParent.GetId()] = grandParent
		}
	}

	return ancestors, nil
}

func (d *dag) GetDescendents(id string) (map[string]*Node, error) {
	if _, exists := d.nodes[id]; !exists {
		return nil, fmt.Errorf(NodeNotFound)
	}

	var descendents map[string]*Node = make(map[string]*Node)
	for _, childNode := range d.nodes[id].children {
		descendents[childNode.GetId()] = childNode

		grandchildren, err := d.GetDescendents(childNode.GetId())
		if len(grandchildren) == 0 {
			continue
		}

		if err != nil {
			return descendents, err
		}

		for _, grandchild := range grandchildren {
			descendents[grandchild.GetId()] = grandchild
		}
	}

	return descendents, nil
}

func (d *dag) DeleteRelation(parentId string, childId string) error {
	if _, exists := d.nodes[parentId]; !exists {
		return fmt.Errorf(NodeNotFound)
	}

	if _, exists := d.nodes[childId]; !exists {
		return fmt.Errorf(NodeNotFound)
	}

	if _, relation := d.nodes[parentId].children[childId]; !relation {
		return fmt.Errorf(NodeRelationNotFound)
	}

	if _, relation := d.nodes[childId].parents[parentId]; !relation {
		return fmt.Errorf(NodeRelationNotFound)
	}

	delete(d.nodes[parentId].children, childId)
	delete(d.nodes[childId].parents, parentId)
	return nil
}

func (d *dag) delete(nodeId string) error {
	if _, exists := d.nodes[nodeId]; !exists {
		return fmt.Errorf(NodeNotFound)
	}

	for _, n := range d.nodes[nodeId].parents {
		delete(n.children, nodeId)
	}

	for _, n := range d.nodes[nodeId].children {
		delete(n.parents, nodeId)
	}

	d.nodes[nodeId] = nil
	return nil
}

func (d *dag) AddRelation(parentId string, childId string) error {
	if _, exists := d.nodes[parentId]; !exists {
		return fmt.Errorf(NodeNotFound)
	}

	if _, exists := d.nodes[childId]; !exists {
		return fmt.Errorf(NodeNotFound)
	}

	ancestors, _ := d.GetAncestors(parentId)
	fmt.Println("anc ", ancestors)
	if _, exists := ancestors[childId]; exists {
		return fmt.Errorf(CyclicDependencyError+" between node %v and %v", parentId, childId)
	}

	d.nodes[parentId].children[childId] = d.nodes[childId]
	d.nodes[childId].parents[parentId] = d.nodes[parentId]

	return nil
}

func (d *dag) AddMember(newNodeId string, parentId string, childId string) error {
	if _, exists := d.nodes[parentId]; !exists {
		return fmt.Errorf(NodeNotFound)
	}

	if _, exists := d.nodes[childId]; !exists {
		return fmt.Errorf(NodeNotFound)
	}

	if _, exists := d.nodes[newNodeId]; !exists {
		return fmt.Errorf(NodeNotFound)
	}

	err := d.AddRelation(parentId, newNodeId)
	if err != nil {
		return err
	}

	err = d.AddRelation(newNodeId, childId)
	if err != nil {
		return err
	}

	return nil
}
