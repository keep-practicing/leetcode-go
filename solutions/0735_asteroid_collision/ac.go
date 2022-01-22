/*
735. Asteroid Collision
https://leetcode.com/problems/asteroid-collision/

We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size,
and the sign represents its direction (positive meaning right, negative meaning left).
Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions.
If two asteroids meet, the smaller one will explode. If both are the same size, both will explode.
Two asteroids moving in the same direction will never meet.

Note:
The length of asteroids will be at most 10000.
Each asteroid will be a non-zero integer in the range [-1000, 1000]..
*/
// time: 2019-01-14

package ac

// stack
// time complexity:  O(N), where NN is the number of asteroids. Our stack pushes and pops each asteroid at most once.
// space complexity: O(N), the size of stack.
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, asteroid := range asteroids {
		flag := true
		for len(stack) > 0 && asteroid < 0 && stack[len(stack)-1] > 0 {
			if stack[len(stack)-1] == -asteroid {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack)-1] < -asteroid {
				stack = stack[:len(stack)-1]
				continue
			}
			flag = false
			break
		}
		if flag {
			stack = append(stack, asteroid)
		}
	}
	return stack
}
