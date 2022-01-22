package reverseinteger

import "testing"

func TestReverse(t *testing.T) {
	testCases := []int{
		123, -123, 10, 0, 12239999999999,
	}
	expected := []int{321, -321, 1, 0, 0}

	for index, data := range testCases {
		if res := reverse(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
