package stringmatch_test

import (
	"testing"

	stringmatch "github.com/lixiaoqing18/gobyexample/datastructure_and_algorithm/string_match"
)

func Test_RK(t *testing.T) {
	data := []struct {
		name     string
		str      string
		match    string
		expected int
	}{
		{"found1", "abcdefghijkljavagolangpython", "java", 12},
		{"notfound1", "abcdefghijkljavagolangpython", "cplusplus", -1},
		{"notfound2", "abcdefghijkljavagolangpython", "pytorch", -1},
		{"found2", "abcdefghijkljavagolangpython", "golang", 16},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result := stringmatch.RK(d.str, d.match)
			if result != d.expected {
				t.Errorf("%s must be in %s at %d, but get %d", d.match, d.str, d.expected, result)
			}
		})
	}
}

func BenchmarkRK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := stringmatch.RK(stringmatch.MainForTest, stringmatch.PatternForTest)
		blackhole = result
	}
}
