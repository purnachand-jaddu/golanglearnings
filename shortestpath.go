package main

import (
	"fmt"
	"math"
)

type MyNode struct {
	Val   int
	State int
}

type Entry struct {
	Node    MyNode
	Length  int
	Visited map[MyNode]bool
}

func shortestPathLength(graph [][]int) int {
	length := len(graph)
	all := (1 << length) - 1
	q := []Entry{}
	for i := 0; i < len(graph); i++ {
		node := MyNode{
			Val:   i,
			State: 1 << i,
		}
		visited := map[MyNode]bool{}
		visited[node] = true
		q = append(q, Entry{
			Node:    node,
			Visited: visited,
			Length:  1,
		})
	}
	min := math.MaxInt32
	for len(q) > 0 {
		// fmt.Println(len(q))
		aux := []Entry{}
		for _, entry := range q {
			if entry.Node.State == all {
				min = Min(min, entry.Length-1)
				continue
			}
			for _, adj := range graph[entry.Node.Val] {
				node := MyNode{
					Val:   adj,
					State: entry.Node.State | (1 << adj),
				}
				_, ok := entry.Visited[node]
				if !ok {
					newEntry := Entry{
						Node:    node,
						Visited: map[MyNode]bool{},
					}
					for k, v := range entry.Visited {
						newEntry.Visited[k] = v
					}
					newEntry.Visited[node] = true
					newEntry.Length = entry.Length + 1
					aux = append(aux, newEntry)
				}
			}
		}
		q = aux
	}
	return min
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(shortestPathLength([][]int{{2, 3, 7}, {3, 6}, {0, 4}, {0, 1, 4, 5}, {3, 7, 2, 6}, {3}, {4, 1}, {4, 0}}))
}
