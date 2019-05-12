package base62

import (
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(100000000000000000)
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Correctly encode negative values", args{-10}, ""},
		{"Correctly encode 10", args{10}, "b"},
		{"Correctly encode 100", args{100}, "2B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.n); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("dyDxDu1234yupiIdyDxDu1234yupiI")
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Correctly decode a", args{"b"}, 10},
		{"Correctly decode 1C", args{"2B"}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.s); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
