package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, class := range topoSort(prereqs) {
		fmt.Printf("%2d : %s \n", i, class)
	}
}

//这个算法似乎有点拗口
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	//定义了一个匿名函数	其实定义的是一个深度优先搜索	树叶节点被首先访问，这个节点没有先修课，接下来不断深入，完成拓扑排序
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			//如果这个没有被访问过，就去访问
			if !seen[item] {
				seen[item] = true
				//访问子节点
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
