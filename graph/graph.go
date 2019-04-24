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

func (self *Graph) AllSCs() (result []godip.Province) {
	for name, node := range self.Nodes {
		if node.SC != nil {
			result = append(result, name)
		}
	}
	return
}

// Edges returns the edges leading away from the specified province, or if reverse
// is set to true then it instead returns the edges leading to it.
func (self *Graph) Edges(n godip.Province, reverse bool) (result map[godip.Province]map[godip.Flag]bool) {
	result = map[godip.Province]map[godip.Flag]bool{}
	for p, edge := range self.edges(n, reverse) {
		result[p] = edge.Flags
	}
	return
}

// When reverse is true then the returned edges will lead to the province n; when
// false then they will lead away from it.
func (self *Graph) edges(n godip.Province, reverse bool) (result map[godip.Province]*edge) {
	p, c := n.Split()
	if node, ok := self.Nodes[p]; ok {
		if sub, ok := node.Subs[c]; ok {
			if reverse {
				result = sub.ReverseEdges
			} else {
				result = sub.Edges
			}
		}
	}
	return
}

type pathStep struct {
	path []godip.Province
	src  godip.Province
	dst  godip.Province
}

// pathHelper returns a path to or from a target province satisfying the given
// pathFilter and starting with the given path steps. When reverse is true then the
// paths will lead to the province; when false they will lead away from it. By
// seeding this with a step from nowhere ("") to a starting point then this method
// can be used to obtain a path between two points. The filter can be used to
// specify the type of provinces that the path can go through, but it can also be
// used as a callback function to allow extracting information about all potential
// matching paths.
func (self *Graph) pathHelper(target godip.Province, reverse bool, queue []pathStep, seen map[[2]godip.Province]bool, filter godip.PathFilter) []godip.Province {
	var newQueue []pathStep
	for _, step := range queue {
		key := [2]godip.Province{step.src, step.dst}
		if seen[key] {
			continue
		}
		seen[key] = true
		stepTarget := step.dst
		if reverse {
			stepTarget = step.src
		}
		for name, edge := range self.edges(stepTarget, reverse) {
			if filter == nil || filter(name, edge.Flags, edge.sub.Flags, edge.sub.node.SC, step.path) {
				thisPath := append(append([]godip.Province{}, step.path...), name)
				if name == target {
					return thisPath
				}
				step := pathStep{path: thisPath, src: step.dst, dst: name}
				if reverse {
					step = pathStep{path: thisPath, src: name, dst: step.src}
				}
				newQueue = append(newQueue, step)
			}
		}
	}
	if len(newQueue) > 0 {
		return self.pathHelper(target, reverse, newQueue, seen, filter)
	}
	return nil
}

// Path returns a list of provinces that go from first to last.  When reverse is
// set to true then all edges are traversed backwards (i.e. to find a root that a
// unit at last could use to get to first). The filter can be used to specify the
// type of provinces that the path can go through, but it can also be used as a
// callback function to allow extracting information about all potential matching
// paths (nb. last can be set to "" to use this callback with all reachable
// provinces from first).
func (self *Graph) Path(first, last godip.Province, reverse bool, filter godip.PathFilter) []godip.Province {
	queue := []pathStep{
		pathStep{path: nil, src: "", dst: first},
	}
	if reverse {
		queue = []pathStep{
			pathStep{path: nil, src: first, dst: ""},
		}
	}
	return self.pathHelper(last, reverse, queue, map[[2]godip.Province]bool{}, filter)
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
			Name:         n,
			Edges:        make(map[godip.Province]*edge),
			ReverseEdges: make(map[godip.Province]*edge),
			node:         self,
			Flags:        make(map[godip.Flag]bool),
		}
	}
	return self.Subs[n]
}

type edge struct {
	Flags map[godip.Flag]bool

	sub *SubNode
}

type SubNode struct {
	Name         godip.Province
	Edges        map[godip.Province]*edge
	ReverseEdges map[godip.Province]*edge
	Flags        map[godip.Flag]bool

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
	target.ReverseEdges[self.getName()] = &edge{
		sub:   self,
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
