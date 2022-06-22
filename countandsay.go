package main

import "fmt"

func countAndSay(n int) string {
	dp := []string{"", "1", "11", "21", "1211"}
	for i := 5; i <= n; i++ {
		prev := dp[i-1]
		var currByte byte
		var currFreq int
		s := ""
		for j := 0; j < len(prev); j++ {
			if j == 0 {
				currByte = prev[j]
				currFreq = 1
			} else {
				if prev[j] == currByte {
					currFreq++
				} else {
					s = fmt.Sprintf("%s%d%c", s, currFreq, currByte)
					currByte = prev[j]
					currFreq = 1
				}
			}
		}
		s = fmt.Sprintf("%s%d%c", s, currFreq, currByte)
		dp = append(dp, s)
	}
	return dp[n]
}
