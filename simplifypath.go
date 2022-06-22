package main

import (
	"fmt"
)

func simplifyPath(path string) string {
	subs := ""
	parts := []string{}
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if len(subs) > 0 {
				parts = append(parts, subs)
			}
			subs = ""
		} else {
			subs = fmt.Sprintf("%s%c", subs, path[i])
		}
	}
	if len(subs) > 0 {
		parts = append(parts, subs)
	}
	simplifiedParts := []string{}
	for _, part := range parts {
		if part == "." {
			//ignore
		} else if part == ".." {
			currLength := len(simplifiedParts)
			if currLength > 0 {
				simplifiedParts = simplifiedParts[:currLength-1]
			}
		} else {
			simplifiedParts = append(simplifiedParts, part)
		}
	}
	result := ""
	for _, part := range simplifiedParts {
		result = fmt.Sprintf("%s%c", result, '/')
		result = fmt.Sprintf("%s%s", result, part)
	}
	if len(result) == 0 {
		result = "/"
	}
	return result
}
