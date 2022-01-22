package ac

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	testCases := [][]int{
		{5, 10, -5},
		{-2, 2, -1, -2},
		{-2, -1, 1, 2},
		{8, -8},
		{10, 2, -5},
	}

	expected := [][]int{
		{5, 10},
		{-2},
		{-2, -1, 1, 2},
		{},
		{10},
	}

	for index, asteroids := range testCases {
		if res := asteroidCollision(asteroids); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}
