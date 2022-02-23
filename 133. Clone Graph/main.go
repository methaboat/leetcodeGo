package main

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func dfs(node *Node, mp map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := mp[node.Val]; !ok {
		copy := &Node{Val: node.Val}
		mp[node.Val] = copy
		for _, nx := range node.Neighbors {
			copy.Neighbors = append(copy.Neighbors, dfs(nx, mp))
		}
	}
	return mp[node.Val]
}

func cloneGraph(node *Node) *Node {
	mp := make(map[int]*Node)
	return dfs(node, mp)
}
