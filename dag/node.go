package dag

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Node struct {
	id       string
	name     string
	children map[string]*Node
	parents  map[string]*Node
}

func NewNode(id string, name string) Node {
	node := Node{
		id:       id,
		name:     name,
		children: make(map[string]*Node),
		parents:  make(map[string]*Node),
	}

	return node
}

func (node Node) GetId() string {
	return node.id
}

func (node Node) GetName() string {
	return node.name
}

func (node Node) GetChildren() map[string]*Node {
	return node.children
}

func (node Node) GetParents() map[string]*Node {
	return node.parents
}

func (node Node) validate() error {
	err := validation.ValidateStruct(&node,
		validation.Field(&node.id, validation.Required),
		validation.Field(&node.name, validation.Required),
	)

	if err != nil {
		return fmt.Errorf("error : expected non empty value %v", err)
	}

	return nil
}
