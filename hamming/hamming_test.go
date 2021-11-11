package hamming

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name    string
		args    args
		want    int
	}{
		{"aa/aa", args{"aa", "aa"}, 0},
		{"ab/aa", args{"ab", "aa"}, 1},
		{"ab/ba", args{"ab", "ba"}, 2},
		{"ab/aaa", args{"ab", "aaa"}, 0},
		{"bbb/a", args{"bbb", "a"}, 0},
		{"bbb/b", args{"bbb", "b"}, 1},
		{"🙂😄🙂😄/😄🙂😄🙂", args{"🙂😄🙂😄", "😄🙂😄🙂"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EditDistance(tt.args.str1, tt.args.str2)
			if got != tt.want {
				t.Errorf("EditDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
