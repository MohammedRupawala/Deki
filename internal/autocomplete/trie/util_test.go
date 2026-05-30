package trie

import "testing"

func TestFindLongestMatch(t *testing.T) {
	tests := []struct {
		name    string
		matches []string
		want    string
	}{
		{
			name:    "empty input",
			matches: []string{},
			want:    "",
		},
		{
			name:    "single string",
			matches: []string{"single"},
			want:    "single",
		},
		{
			name:    "identical strings",
			matches: []string{"abc", "abc"},
			want:    "abc",
		},
		{
			name:    "common prefix",
			matches: []string{"flower", "flow", "flight"},
			want:    "fl",
		},
		{
			name:    "no common prefix",
			matches: []string{"dog", "racecar", "car"},
			want:    "",
		},
		{
			name:    "common prefix length four",
			matches: []string{"interview", "interrupt", "integrate"},
			want:    "inte",
		},
		{
			name:    "prefix equals shortest string",
			matches: []string{"abc", "abcd", "abcde"},
			want:    "abc",
		},
		{
			name:    "common prefix one character",
			matches: []string{"apple", "ant", "axe"},
			want:    "a",
		},
		{
			name:    "match until last character",
			matches: []string{"aaaa", "aaaa", "aaaa"},
			want:    "aaaa",
		},
		{
			name:    "fails at last character",
			matches: []string{"aaaa", "aaab", "aaac"},
			want:    "aaa",
		},
		{
			name:    "two strings",
			matches: []string{"prefix", "prelude"},
			want:    "pre",
		},
	}

	tr := &Trie{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tr.FindLongestMatch(tc.matches)
			if got != tc.want {
				t.Fatalf(
					"FindLongestMatch(%v) = %q, want %q",
					tc.matches,
					got,
					tc.want,
				)
			}
		})
	}
}