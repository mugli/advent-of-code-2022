package main

import "path"

type nodeType int

const (
	fileNode nodeType = iota
	dirNode
)

const (
	prompt = "$"
	cdCmd  = "cd"
	dirCmd = "dir"
	//lsCmd  = "ls" // not used
)

type node struct {
	nodeType nodeType
	name     string
	path     string
	size     int64
	parent   *node
	children []*node
}

func (n *node) readDirs() (dirs []*node) {
	for _, node := range n.children {
		if node.nodeType == dirNode {
			dirs = append(dirs, node)
		}
	}

	return
}

func (n *node) addChildFile(name string, size int64) *node {
	file := &node{
		nodeType: fileNode,
		name:     name,
		path:     path.Join(n.path, name),
		size:     size,
		parent:   n,
	}
	n.children = append(n.children, file)

	// update sizes in the parent directories
	parent := n
	for parent != nil {
		parent.size += size
		parent = parent.parent
	}

	return file
}

func (n *node) addChildDir(name string) *node {
	dir := &node{
		nodeType: dirNode,
		name:     name,
		path:     path.Join(n.path, name),
		parent:   n,
	}
	n.children = append(n.children, dir)

	return dir
}

func newFS() *node {
	return &node{
		nodeType: dirNode,
		path:     "/",
	}
}

type shell struct {
	rootDir *node
	curDir  *node
}

func newShell(rootFS *node) *shell {
	return &shell{
		rootDir: rootFS,
		curDir:  rootFS,
	}
}

func (s *shell) cd(dir string) {
	switch dir {
	case "..":
		parent := s.curDir.parent
		if parent != nil {
			s.curDir = parent
		}
	case "/":
		s.curDir = s.rootDir
	default:
		dirs := s.curDir.readDirs()

		var foundDir *node
		for _, d := range dirs {
			if d.name == dir {
				foundDir = d
			}
		}

		if foundDir != nil {
			s.curDir = foundDir
		}
	}
}

func (s *shell) mkdir(name string) {
	s.curDir.addChildDir(name)
}

// https://man7.org/linux/man-pages/man1/fallocate.1.html
func (s *shell) fallocate(name string, size int64) {
	s.curDir.addChildFile(name, size)
}

type walkFunc func(node *node)

func walkDir(startNode *node, fn walkFunc) {
	fn(startNode)

	dirs := startNode.readDirs()
	for _, d := range dirs {
		walkDir(d, fn)
	}
}
