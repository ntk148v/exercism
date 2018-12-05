package tree

import (
	"fmt"
	"sort"
)

// Recorad is a database record.
type Record struct {
	ID, Parent int
}

// Node is a node in a tree structure.
type Node struct {
	ID       int
	Children []*Node
}

// Interfaces for storing records.
type byID []Record

func (ids byID) Len() int           { return len(ids) }
func (ids byID) Swap(i, j int)      { ids[i], ids[j] = ids[j], ids[i] }
func (ids byID) Less(i, j int) bool { return ids[i].ID < ids[j].ID }

// Build creates a tree from a given set of records.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Sort(byID(records))
	nodes := make([]*Node, len(records))
	for i, r := range records {
		nodes[i] = &Node{ID: r.ID}
		if i == 0 {
			if r.ID != 0 || r.Parent != 0 {
				return nil, fmt.Errorf("Invalid root record")
			}
			continue
		} else if i != r.ID || r.ID <= r.Parent || r.ID >= len(records) {
			return nil, fmt.Errorf("Invalid record")
		}

		if parent := nodes[r.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[i])
		}
	}
	return nodes[0], nil
}
