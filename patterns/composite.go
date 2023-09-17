package patterns

import (
	"fmt"
	"log"
)

type Node interface {
	Display() string
}

type NodeTree interface {
	Node
	Components() []NodeTree
}

type Leaf struct {
	label string
}

func (leaf Leaf) Display() string {
	return leaf.label
}

func (leaf Leaf) Components() []NodeTree {
	return nil
}

type Branch struct {
	label      string
	components []NodeTree
}

func (branch Branch) Display() string {
	componentString := branch.label
	for _, component := range branch.Components() {
		componentString += fmt.Sprintf("-> %s", component.Display())
	}
	return componentString
}

func (branch Branch) Components() []NodeTree {
	return branch.components
}

func Composite() {
	l0 := Leaf{label: "l0"}
	l1 := Leaf{label: "l1"}
	l2 := Leaf{label: "l2"}
	l3 := Leaf{label: "l3"}

	b0 := Branch{
		label:      "b0",
		components: []NodeTree{l2, l3},
	}

	b1 := Branch{
		label:      "b1",
		components: []NodeTree{l0, l1, b0},
	}
	log.Print(b1.Display())
}
