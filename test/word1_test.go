package test

import "testing"

func TestIsPalindrome1(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"Kayak", true},
		{"detartrated", true},
		{"A man,a plan,a canal:panama", true},
		{"desserts", false},
	}

	for _, test := range tests {
		if got := IsPalindrome1(test.input); !got == test.want {
			t.Errorf("IsPalindrome1(%q)= %v", test.input, got)
		}
	}
}
