package stringmatch_test

import (
	"testing"

	stringmatch "github.com/lixiaoqing18/gobyexample/datastructure_and_algorithm/string_match"
)

func Test_KMP(t *testing.T) {
	data := []struct {
		name     string
		str      string
		match    string
		expected int
	}{
		{"found1", "abcabcabcabd", "abcabd", 6},
		{"notfound1", "abcdefghijkljavagolangpython", "cplusplus", -1},
		{"notfound2", "abcdefghijkljavagolangpython", "pytorch", -1},
		{"found2", "abcdefghijkljavagolangpython", "golang", 16},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := stringmatch.KMP(d.str, d.match)
			if result != d.expected {
				t.Errorf("%s must be in %s at %d, but get %d", d.match, d.str, d.expected, result)
			}
		})
	}
}

func BenchmarkKMP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := stringmatch.KMP(stringmatch.MainForTest, stringmatch.PatternForTest)
		blackhole = result
	}
}
