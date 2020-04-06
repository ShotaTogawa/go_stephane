package iterations

import "testing"

const repeatCount = 5

// Repect iterates argument char repeatCount times
func Repeat(char string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expect := "aaaaa"

	if repeated != expect {
		t.Errorf("expected %q but got %q", repeated, expect)
	}
}

// 実行方法: go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
