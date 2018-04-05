package graph

import (
	"bytes"
	"fmt"

	"github.com/zond/godip"
)

func New() *Graph {
	return &Graph{
		Nodes: make(map[godip.Province]*Node),
	}
}

type Graph struct {
	Nodes map[godip.Province]*Node
}

func (self *Graph) String() string {
	buf := new(bytes.Buffer)
	for _, n := range self.Nodes {
		fmt.Fprintf(buf, "%v", n)
	}
	return string(buf.Bytes())
}

func (self *Graph) Has(n godip.Province) (result bool) {
	p, c := n.Split()
	if node, ok := self.Nodes[p]; ok {
		if _, ok := node.Subs[c]; ok {
			result = true
		}
	}
	return
}

func (self *Graph) AllFlags(n godip.Province) (result map[godip.Flag]bool) {
	result = map[godip.Flag]bool{}
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

func (self *Graph) Flags(n godip.Province) (result map[godip.Flag]bool) {
	p, c := n.Split()
	if node, ok := self.Nodes[p]; ok {
		if sub, ok := node.Subs[c]; ok {
			result = sub.Flags
		}
	}
	return
}

func (self *Graph) SC(n godip.Province) (result *godip.Nation) {
	if node, ok := self.Nodes[n.Super()]; ok {
		result = node.SC
	}
	return
}

func (self *Graph) SCs(n godip.Nation) (result []godip.Province) {
	for name, node := range self.Nodes {
		if node.SC != nil && *node.SC == n {
			result = append(result, name)
		}
	}
	return
}

func (self *Graph) Edges(n godip.Province) (result map[godip.Province]map[godip.Flag]bool) {
	result = map[godip.Province]map[godip.Flag]bool{}
	for p, edge := range self.edges(n) {
		result[p] = edge.Flags
	}
	return
}

func (self *Graph) edges(n godip.Province) (result map[godip.Province]*edge) {
	p, c := n.Split()
	if node, ok := self.Nodes[p]; ok {
		if sub, ok := node.Subs[c]; ok {
			result = sub.Edges
		}
	}
	return
}

type pathStep struct {
	path []godip.Province
	src  godip.Province
	dst  godip.Province
}

func (self *Graph) pathHelper(dst godip.Province, queue []pathStep, seen map[[2]godip.Province]bool, filter godip.PathFilter) []godip.Province {
	var newQueue []pathStep
	for _, step := range queue {
		key := [2]godip.Province{step.src, step.dst}
		if seen[key] {
			continue
		}
		seen[key] = true
		for name, edge := range self.edges(step.dst) {
			if filter == nil || filter(name, edge.Flags, edge.sub.Flags, edge.sub.node.SC, step.path) {
				thisPath := append(append([]godip.Province{}, step.path...), name)
				if name == dst {
					return thisPath
				}
				newQueue = append(newQueue, pathStep{
					path: thisPath,
					src:  step.dst,
					dst:  name,
				})
			}
		}
	}
	if len(newQueue) > 0 {
		return self.pathHelper(dst, newQueue, seen, filter)
	}
	return nil
}

func (self *Graph) Path(src, dst godip.Province, filter godip.PathFilter) []godip.Province {
	queue := []pathStep{
		pathStep{
			path: nil,
			src:  "",
			dst:  src,
		},
	}
	return self.pathHelper(dst, queue, map[[2]godip.Province]bool{}, filter)
}

func (self *Graph) Coasts(prov godip.Province) (result []godip.Province) {
	if node, ok := self.Nodes[prov.Super()]; ok {
		for _, sub := range node.Subs {
			result = append(result, sub.getName())
		}
	}
	return
}

func (self *Graph) Prov(n godip.Province) *SubNode {
	p, c := n.Split()
	if self.Nodes[p] == nil {
		self.Nodes[p] = &Node{
			Name:  p,
			Subs:  make(map[godip.Province]*SubNode),
			graph: self,
		}
	}
	return self.Nodes[p].sub(c)
}

func (self *Graph) Nations() (result []godip.Nation) {
	found := map[godip.Nation]bool{}
	for _, node := range self.Nodes {
		if node.SC != nil && *node.SC != godip.Neutral {
			found[*node.SC] = true
		}
	}
	result = make([]godip.Nation, 0, len(found))
	for nation := range found {
		result = append(result, nation)
	}
	return
}

func (self *Graph) Provinces() (result []godip.Province) {
	for _, node := range self.Nodes {
		for _, sub := range node.Subs {
			result = append(result, sub.getName())
		}
	}
	return
}

type Node struct {
	Name godip.Province
	Subs map[godip.Province]*SubNode
	SC   *godip.Nation

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

func (self *Node) sub(n godip.Province) *SubNode {
	if self.Subs[n] == nil {
		self.Subs[n] = &SubNode{
			Name:  n,
			Edges: make(map[godip.Province]*edge),
			node:  self,
			Flags: make(map[godip.Flag]bool),
		}
	}
	return self.Subs[n]
}

type edge struct {
	Flags map[godip.Flag]bool

	sub *SubNode
}

type SubNode struct {
	Name  godip.Province
	Edges map[godip.Province]*edge
	Flags map[godip.Flag]bool

	node *Node
}

func (self *SubNode) String() string {
	buf := new(bytes.Buffer)
	if self.Name != "" {
		fmt.Fprintf(buf, "%v ", self.Name)
	}
	flags := make([]godip.Flag, 0, len(self.Flags))
	for flag, _ := range self.Flags {
		flags = append(flags, flag)
	}
	if len(flags) > 0 {
		fmt.Fprintf(buf, "%v ", flags)
	}
	dests := make([]string, 0, len(self.Edges))
	for n, edge := range self.Edges {
		var flags []godip.Flag
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

func (self *SubNode) getName() godip.Province {
	return self.node.Name.Join(self.Name)
}

func (self *SubNode) Conn(n godip.Province, flags ...godip.Flag) *SubNode {
	target := self.node.graph.Prov(n)
	flagMap := make(map[godip.Flag]bool)
	for _, flag := range flags {
		flagMap[flag] = true
	}
	self.Edges[target.getName()] = &edge{
		sub:   target,
		Flags: flagMap,
	}
	return self
}

func (self *SubNode) SC(n godip.Nation) *SubNode {
	self.node.SC = &n
	return self
}

func (self *SubNode) Flag(flags ...godip.Flag) *SubNode {
	for _, flag := range flags {
		self.Flags[flag] = true
	}
	return self
}

func (self *SubNode) Prov(n godip.Province) *SubNode {
	return self.node.graph.Prov(n)
}

func (self *SubNode) Done() *Graph {
	return self.node.graph
}
