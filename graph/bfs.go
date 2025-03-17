package graph

import "math"

/*
算法图解
广度优先搜索用于非加权图查找最短路劲
查找姓名以M结尾的朋友
*/
func FindM() bool {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"tom", "jonny"}
	graph["tom"] = []string{}
	graph["jonny"] = []string{}
	graph["peggy"] = []string{}
	graph["anuj"] = []string{}

	record := make(map[string]int)
	queue := make([]string, 0, len(graph))
	queue = append(queue, graph["you"]...)
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		if record[name] > 0 {
			continue
		}
		if string(name[len(name)-1]) == "m" {
			return true
		}
		record[name]++
		queue = append(queue, graph[name]...)
	}
	return false
}

/*
狄克斯特拉算法用于加权图查找最短路径(不可用于有环图,负权图)
*/
func FindShortest() int {
	//节点散列表
	graph := make(map[string]map[string]int)
	graph["start"] = map[string]int{"A": 6, "B": 2}
	graph["A"] = map[string]int{"end": 1}
	graph["B"] = map[string]int{"A": 3, "end": 5}

	//节点花费散列表
	cost := make(map[string]int)
	cost["A"] = 6
	cost["B"] = 2
	cost["end"] = math.MaxInt

	//父类节点散列表
	parent := make(map[string]string)
	parent["A"] = "start"
	parent["B"] = "start"
	parent["end"] = ""

	//已处理过节点
	processed := make(map[string]bool)

	//获取最低开销节点
	node := getLowest(cost, processed)
	for node != "" {
		//起点至node节点的开销
		fee := cost[node]
		//遍历node节点的相邻节点
		for k, v := range graph[node] {
			//起点经过node节点抵达相邻节点的开销
			newCost := fee + v
			//若小于之前的开销则说明找了花销更小的途径
			if newCost < cost[k] {
				//变更起点至该节点的花销
				cost[k] = newCost
				//更新该节点的父节点为node节点
				parent[k] = node
			}
		}
		//加入已处理过的节点
		processed[node] = true
		node = getLowest(cost, processed)
	}
	return cost["end"]
}

// 获取最低花费节点
func getLowest(cost map[string]int, processed map[string]bool) string {
	lowest := math.MaxInt
	key := ""
	for k, v := range cost {
		if v < lowest && !processed[k] {
			lowest = v
			key = k
		}
	}
	return key
}
