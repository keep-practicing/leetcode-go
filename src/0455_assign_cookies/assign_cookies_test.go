package assigncookies

import "testing"

type args struct {
	g []int
	s []int
}

func TestAssignCookies(t *testing.T) {
	testData := []args{
		{
			g: []int{1, 2, 3},
			s: []int{1, 1},
		},
		{
			g: []int{1, 2},
			s: []int{1, 2, 3},
		},
	}
	expectedData := []int{1, 2}

	for index, data := range testData {
		if res := findContentChildren(data.g, data.s); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
