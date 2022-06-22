package main

import "fmt"

type TreeNode struct {
	String   string
	Children []*TreeNode
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := createMapping()
	q := []*TreeNode{}
	root := TreeNode{}
	q = append(q, &root)
	for i := 0; i < len(digits); i++ {
		aux := []*TreeNode{}
		children := m[digits[i]]
		for j := 0; j < len(q); j++ {
			for k := 0; k < len(children); k++ {
				newChild := TreeNode{}
				newChild.String = fmt.Sprintf("%s%c", q[j].String, children[k])
				q[j].Children = append(q[j].Children, &newChild)
				aux = append(aux, &newChild)
			}
		}
		q = aux
	}
	result := []string{}
	for i := 0; i < len(q); i++ {
		result = append(result, q[i].String)
	}
	return result
}

func createMapping() map[byte][]byte {
	m := make(map[byte][]byte)
	m['2'] = []byte{'a', 'b', 'c'}
	m['3'] = []byte{'d', 'e', 'f'}
	m['4'] = []byte{'g', 'h', 'i'}
	m['5'] = []byte{'j', 'k', 'l'}
	m['6'] = []byte{'m', 'n', 'o'}
	m['7'] = []byte{'p', 'q', 'r', 's'}
	m['8'] = []byte{'t', 'u', 'v'}
	m['9'] = []byte{'w', 'x', 'y', 'z'}
	return m
}
