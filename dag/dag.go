package dag

import (
	"fmt"
)

type dag struct {
	nodes map[string]*Node
}

func New() *dag {
	d := dag{}
	d.nodes = make(map[string]*Node)
	return &d
}

func (d *dag) AddNode(id string, name string) error {
	if _, exists := d.nodes[id]; exists {
		return fmt.Errorf(CommonErrorFormat, NodeAlreadyExistsError, id)
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
		return nil, fmt.Errorf(CommonErrorFormat, NodeNotFound, id)
	}

	return d.nodes[id].GetParents(), nil
}

func (d *dag) GetChildren(id string) (map[string]*Node, error) {
	if _, exists := d.nodes[id]; !exists {
		return nil, fmt.Errorf(CommonErrorFormat, NodeNotFound, id)
	}

	return d.nodes[id].GetChildren(), nil
}

func (d *dag) GetAncestors(id string, ancestors map[string]*Node) error {
	if _, exists := d.nodes[id]; !exists {
		return fmt.Errorf(CommonErrorFormat, NodeNotFound, id)
	}

	for _, parent := range d.nodes[id].GetParents() {
		if _, exists := ancestors[parent.GetId()]; !exists {
			ancestors[parent.GetId()] = parent
			err := d.GetAncestors(parent.GetId(), ancestors)
			if err != nil {
				return fmt.Errorf(CommonErrorFormat, AncestorsComputationError, id)
			}
		}
	}

	return nil
}

func (d *dag) GetDescendents(id string, descendents map[string]*Node) error {
	if _, exists := d.nodes[id]; !exists {
		return fmt.Errorf(CommonErrorFormat, NodeNotFound, id)
	}

	for _, child := range d.nodes[id].GetChildren() {
		if _, exists := descendents[child.GetId()]; !exists {
			descendents[child.GetId()] = child
			err := d.GetDescendents(child.GetId(), descendents)
			if err != nil {
				return fmt.Errorf(CommonErrorFormat, AncestorsComputationError, id)
			}
		}
	}

	return nil
}

func (d *dag) DeleteRelation(parentId string, childId string) error {
	if _, exists := d.nodes[parentId]; !exists {
		return fmt.Errorf(CommonErrorFormat, NodeNotFound, parentId)
	}

	if _, exists := d.nodes[childId]; !exists {
		return fmt.Errorf(CommonErrorFormat, NodeNotFound, childId)
	}

	if _, relation := d.nodes[parentId].children[childId]; !relation {
		return fmt.Errorf(RelationDagErrorFormat, NodeRelationNotFound, parentId, childId)
	}

	if _, relation := d.nodes[childId].parents[parentId]; !relation {
		return fmt.Errorf(RelationDagErrorFormat, NodeRelationNotFound, parentId, childId)
	}

	delete(d.nodes[parentId].children, childId)
	delete(d.nodes[childId].parents, parentId)
	return nil
}

func (d *dag) delete(id string) error {
	if _, exists := d.nodes[id]; !exists {
		return fmt.Errorf(CommonErrorFormat, NodeNotFound, id)
	}

	for _, n := range d.nodes[id].parents {
		delete(n.children, id)
	}

	for _, n := range d.nodes[id].children {
		delete(n.parents, id)
	}

	d.nodes[id] = nil
	return nil
}

func (d *dag) AddRelation(parentId string, childId string) error {
	if _, exists := d.nodes[parentId]; !exists {
		return fmt.Errorf(CommonErrorFormat, NodeNotFound, parentId)
	}

	if _, exists := d.nodes[childId]; !exists {
		return fmt.Errorf(CommonErrorFormat, NodeNotFound, childId)
	}

	ancestors := make(map[string]*Node)
	err := d.GetAncestors(parentId, ancestors)
	if err != nil {
		return fmt.Errorf(CommonErrorFormat, AncestorsComputationError, parentId)
	}

	if _, exists := ancestors[childId]; exists {
		return fmt.Errorf(RelationDagErrorFormat, CyclicDependencyError, parentId, childId)
	}

	d.nodes[parentId].children[childId] = d.nodes[childId]
	d.nodes[childId].parents[parentId] = d.nodes[parentId]

	return nil
}
