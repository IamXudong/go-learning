// graph 建立一个从字符串到字符串集合的映射
package graph

var graph = make(map[string]map[string]bool)

// AddEdge 添加
func AddEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

// HasEdge 存在性检查
func HasEdge(from, to string) bool {
	return graph[from][to]
}
