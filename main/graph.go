// 使用graph, 测试从字符串到字符串集合的映射
package main

import (
	"fmt"
	"go-learning/graph"
)

func main() {
	graph.AddEdge("Pinghu", "Nanshan")
	graph.AddEdge("LongGang", "Shanghei")
	graph.AddEdge("Pinghu", "Shanghei")

	fmt.Println("Pinghu, Shanghei", graph.HasEdge("Pinghu", "Shanghei"))
	fmt.Println("Pinghu, Nanshan", graph.HasEdge("Pinghu", "Nanshan"))
	fmt.Println("LongGang, Shanghei", graph.HasEdge("LongGang", "Shanghei"))
	fmt.Println("Pinghu, Lanjing", graph.HasEdge("Pinghu", "Lanjing"))
	fmt.Println("Lanjing, Changsha", graph.HasEdge("Lanjing", "Changsha"))
}
