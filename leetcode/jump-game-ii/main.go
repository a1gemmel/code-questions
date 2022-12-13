package main

/*
	Idea: Simple breadth-first search, treating each i as a node with vertices to the next nums[i] nodes
	Problem: Many many vertices makes this very memory-inefficient
*/

func jumpBreadthFirstSearch(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	visited := map[int]bool{}
	toCheck := []struct {
		pos   int
		jumps int
	}{
		{
			pos:   0,
			jumps: 0,
		},
	}

	for {
		v := toCheck[0]
		for i := nums[v.pos]; i >= 1; i-- {
			nextPos := v.pos + i
			if nextPos >= len(nums)-1 {
				return v.jumps + 1
			}
			if !visited[v.pos+i] {
				toCheck = append(toCheck, struct {
					pos   int
					jumps int
				}{pos: v.pos + i, jumps: v.jumps + 1})
			}
		}
		toCheck = toCheck[1:]
	}
}

/*
	Keep track of the reachable interval at each jump index (starting at [0] for index 0)
	At each interval, find the max distance reachable
*/

func jump(nums []int) int {
	maxReachable := 0
	start := 0
	jumps := 0

	for {
		for i := maxReachable; i >= start; i-- {
			if i >= len(nums)-1 {
				return jumps
			}
			maxReachable = max(maxReachable, i+nums[i])
		}
		jumps += 1
		start += 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	jump([]int{2, 3, 1, 1, 4})
}
