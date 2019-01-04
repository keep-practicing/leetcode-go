/*
79. Word Search
https://leetcode.com/problems/word-search/

Given a 2D board and a word, find if the word exists in the grid.
The word can be constructed from letters of sequentially adjacent cell,
where "adjacent" cells are those horizontally or vertically neighboring.
The same letter cell may not be used more than once.
*/
// 2019-01-04

package wordsearch

var (
	d       = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	m, n    int
	visited [][]bool
)

// backtracking
// time complexity: O(m*n*m*n)
// space complexity: O(m*n)
func exist(board [][]byte, word string) bool {
	m = len(board)
	n = len(board[0])
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if searchWord(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func searchWord(board [][]byte, word string, index, startX, startY int) bool {
	if len(word)-1 == index {
		return board[startX][startY] == word[index]
	}
	if board[startX][startY] == word[index] {
		visited[startX][startY] = true
		// 从startX， startY开始，向四个方向寻找
		for i := 0; i < 4; i++ {
			newX := startX + d[i][0]
			newY := startY + d[i][1]
			if inArea(newX, newY) && !visited[newX][newY] && searchWord(board, word, index+1, newX, newY) {
				return true
			}
		}
		visited[startX][startY] = false
	}
	return false
}

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}
