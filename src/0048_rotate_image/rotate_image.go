/*
48. Rotate Image
https://leetcode.com/problems/rotate-image/

You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).

Note:
You have to rotate the image in-place,
which means you have to modify the input 2D matrix directly.
DO NOT allocate another 2D matrix and do the rotation.
*/
// time: 2019-01-02

package ri

// time complexity: O(n^2)
// space complexity: O(1)
func rotate(matrix [][]int) {
	/*
			[
		  		[1,2,3],
		  		[4,5,6],
		  		[7,8,9]
			]
	*/
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	/*
			[
		  		[1,4,7],
		  		[2,5,8],
		  		[3,6,9]
			]
	*/
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	/*
			[
		  		[7,4,1],
		  		[8,5,2],
		  		[9,6,3]
			]
	*/
}
