package issubsequence

import "testing"

type args struct {
	s string
	t string
}

func TestAssignCookies(t *testing.T) {
	testData := []args{
		args{
			s: "abc",
			t: "ahbgdc",
		},
		args{
			s: "acb",
			t: "ahbgdc",
		},
	}
	expectedData := []bool{true, false}

	for index, data := range testData {
		if res := isSubsequence(data.s, data.t); res != expectedData[index] {
			t.Errorf("expected %t, got %t", expectedData[index], res)
		}
	}
}
