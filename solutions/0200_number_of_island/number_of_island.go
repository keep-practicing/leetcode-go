/*
200. Number of Islands
https://leetcode.com/problems/number-of-islands/

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
*/
// time: 2019-01-07

package noi

var (
	d       = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n    int
	visited [][]bool
)

// flood-fill
// time complexity: O(n*m)
// space complexity: O(n*m)
func numIslands(grid [][]byte) int {
	m = len(grid)
	if m <= 0 {
		return 0
	}
	n = len(grid[0])

	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var res int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

// 从grid[x][y]的位置开始，进行floodfill。
func dfs(grid [][]byte, x, y int) {
	visited[x][y] = true
	for i := 0; i < len(d); i++ {
		newX := x + d[i][0]
		newY := y + d[i][1]
		if inArea(newX, newY) && !visited[newX][newY] && grid[newX][newY] == '1' {
			dfs(grid, newX, newY)
		}
	}
}

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
