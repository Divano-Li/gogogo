package main

/*
[[1,1,0,0,1,2],
 [0,0,1,1,2,2]]
*/

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	count := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == byte(1) {
				count++
				dfss(i, j, m, n, grid)
			}
		}
	}
	return count
}

func dfss(i, j, m, n int, grid [][]byte) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == byte(0) {
		return
	} else {
		grid[i][j] = byte(0)
		dfss(i-1, j, m, n, grid)
		dfss(i+1, j, m, n, grid)
		dfss(i, j-1, m, n, grid)
		dfss(i, j+1, m, n, grid)
	}
}
