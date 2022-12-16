package main

import "github.com/montanaflynn/stats"

func uniquePaths(m int, n int) int {
	return stats.Ncr(m+n-2, n-1)
}
