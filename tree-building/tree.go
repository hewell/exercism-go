// Package tree builds tree structure from flat records
package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Build builds array into tree
func Build(records []Record) (*Node, error) {
	n := len(records)
	if n == 0 {
		return nil, nil
	}
	// sort by id
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	// sanity check before loop
	if records[0].ID != 0 {
		return nil, errors.New("Root not found")
	}
	if records[0].Parent != 0 {
		return nil, errors.New("Root has parent")
	}
	if records[n-1].ID != n-1 {
		return nil, errors.New("non-continuous records")
	}

	// slice for holding all the nodes
	nodes := make([]*Node, n)

	for i, record := range records {
		if i > 0 && record.ID <= record.Parent {
			return nil, errors.New("Child node ID has to be bigger than its Parent")
		}

		// new node added
		nodes[i] = &Node{ID: record.ID}
		// if parent exists, append the new node to its parent's Children
		if record.ID != 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[i])
		}
	}
	// return root, root contains the whole tree required
	return nodes[0], nil
}
