package graph

import (
	"bytes"
	"fmt"

	"github.com/zond/godip/common"
)

func New() *Graph {
	return &Graph{
		Nodes: make(map[common.Province]*Node),
	}
}

type Graph struct {
	Nodes map[common.Province]*Node
}

func (self *Graph) String() string {
	buf := new(bytes.Buffer)
	for _, n := range self.Nodes {
		fmt.Fprintf(buf, "%v", n)
	}
	return string(buf.Bytes())
}

func (self *Graph) Has(n common.Province) (result bool) {
	p, c := n.Split()
	if node, ok := self.Nodes[p]; ok {
		if _, ok := node.Subs[c]; ok {
			result = true
		}
	}
	return
}

func (self *Graph) AllFlags(n common.Province) (result map[common.Flag]bool) {
	result = map[common.Flag]bool{}
	p, _ := n.Split()
	if node, ok := self.Nodes[p]; ok {
		for _, sub := range node.Subs {
			for flag, _ := range sub.Flags {
				result[flag] = true
			}
		}
	}
	return
}

func (self *Graph) Flags(n common.Province) (result map[common.Flag]bool) {
	p, c := n.Split()
	if node, ok := self.Nodes[p]; ok {
		if sub, ok := node.Subs[c]; ok {
			result = sub.Flags
		}
	}
	return
}

func (self *Graph) SC(n common.Province) (result *common.Nation) {
	if node, ok := self.Nodes[n.Super()]; ok {
		result = node.SC
	}
	return
}

func (self *Graph) SCs(n common.Nation) (result []common.Province) {
	for name, node := range self.Nodes {
		if node.SC != nil && *node.SC == n {
			result = append(result, name)
		}
	}
	return
}

func (self *Graph) SuperProvinces(sc bool) []common.Province {
	result := []common.Province{}
	for _, node := range self.Nodes {
		if (sc && node.SC != nil) || (!sc && node.SC == nil) {
			result = append(result, node.Name)
		}
	}
	return result
}

func (self *Graph) Edges(n common.Province) (result map[common.Province]map[common.Flag]bool) {
	result = map[common.Province]map[common.Flag]bool{}
	for p, edge := range self.edges(n) {
		result[p] = edge.Flags
	}
	return
}

func (self *Graph) edges(n common.Province) (result map[common.Province]*edge) {
	p, c := n.Split()
	if node, ok := self.Nodes[p]; ok {
		if sub, ok := node.Subs[c]; ok {
			result = sub.Edges
		}
	}
	return
}

type pathStep struct {
	path []common.Province
	pos  common.Province
}

func (self *Graph) pathHelper(dst common.Province, queue []pathStep, filter common.PathFilter, seen map[common.Province]bool) []common.Province {
	var newQueue []pathStep
	for _, step := range queue {
		seen[step.pos] = true
		for name, edge := range self.edges(step.pos) {
			if !seen[name] {
				if filter == nil || filter(name, edge.Flags, edge.sub.Flags, edge.sub.node.SC) {
					thisPath := append(append([]common.Province{}, step.path...), name)
					if name == dst {
						return thisPath
					}
					newQueue = append(newQueue, pathStep{
						path: thisPath,
						pos:  name,
					})
				}
			}
		}
	}
	if len(newQueue) > 0 {
		return self.pathHelper(dst, newQueue, filter, seen)
	}
	return nil
}

func (self *Graph) Path(src, dst common.Province, filter common.PathFilter) []common.Province {
	queue := []pathStep{
		pathStep{
			path: nil,
			pos:  src,
		},
	}
	return self.pathHelper(dst, queue, filter, make(map[common.Province]bool))
}

func (self *Graph) Coasts(prov common.Province) (result []common.Province) {
	if node, ok := self.Nodes[prov.Super()]; ok {
		for _, sub := range node.Subs {
			result = append(result, sub.getName())
		}
	}
	return
}

func (self *Graph) Prov(n common.Province) *SubNode {
	p, c := n.Split()
	if self.Nodes[p] == nil {
		self.Nodes[p] = &Node{
			Name:  p,
			Subs:  make(map[common.Province]*SubNode),
			graph: self,
		}
	}
	return self.Nodes[p].sub(c)
}

func (self *Graph) Provinces() (result []common.Province) {
	for _, node := range self.Nodes {
		for _, sub := range node.Subs {
			result = append(result, sub.getName())
		}
	}
	return
}

type Node struct {
	Name common.Province
	Subs map[common.Province]*SubNode
	SC   *common.Nation

	graph *Graph
}

func (self *Node) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v", self.Name)
	if self.SC != nil {
		fmt.Fprintf(buf, " %v", *self.SC)
	}
	if sub, ok := self.Subs[""]; ok {
		fmt.Fprintf(buf, " %v\n", sub)
	}
	for _, s := range self.Subs {
		if s.Name != "" {
			fmt.Fprintf(buf, "  %v\n", s)
		}
	}
	return string(buf.Bytes())
}

func (self *Node) sub(n common.Province) *SubNode {
	if self.Subs[n] == nil {
		self.Subs[n] = &SubNode{
			Name:  n,
			Edges: make(map[common.Province]*edge),
			node:  self,
			Flags: make(map[common.Flag]bool),
		}
	}
	return self.Subs[n]
}

type edge struct {
	Flags map[common.Flag]bool

	sub *SubNode
}

type SubNode struct {
	Name  common.Province
	Edges map[common.Province]*edge
	Flags map[common.Flag]bool

	node *Node
}

func (self *SubNode) String() string {
	buf := new(bytes.Buffer)
	if self.Name != "" {
		fmt.Fprintf(buf, "%v ", self.Name)
	}
	flags := make([]common.Flag, 0, len(self.Flags))
	for flag, _ := range self.Flags {
		flags = append(flags, flag)
	}
	if len(flags) > 0 {
		fmt.Fprintf(buf, "%v ", flags)
	}
	dests := make([]string, 0, len(self.Edges))
	for n, edge := range self.Edges {
		var flags []common.Flag
		for f, _ := range edge.Flags {
			flags = append(flags, f)
		}
		if len(flags) > 0 {
			dests = append(dests, fmt.Sprintf("%v %v", n, flags))
		} else {
			dests = append(dests, string(n))
		}
	}
	fmt.Fprintf(buf, "=> %v", dests)
	return string(buf.Bytes())
}

func (self *SubNode) getName() common.Province {
	return self.node.Name.Join(self.Name)
}

func (self *SubNode) Conn(n common.Province, flags ...common.Flag) *SubNode {
	target := self.node.graph.Prov(n)
	flagMap := make(map[common.Flag]bool)
	for _, flag := range flags {
		flagMap[flag] = true
	}
	self.Edges[target.getName()] = &edge{
		sub:   target,
		Flags: flagMap,
	}
	return self
}

func (self *SubNode) SC(n common.Nation) *SubNode {
	self.node.SC = &n
	return self
}

func (self *SubNode) Flag(flags ...common.Flag) *SubNode {
	for _, flag := range flags {
		self.Flags[flag] = true
	}
	return self
}

func (self *SubNode) Prov(n common.Province) *SubNode {
	return self.node.graph.Prov(n)
}

func (self *SubNode) Done() *Graph {
	return self.node.graph
}
