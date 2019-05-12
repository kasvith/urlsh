package math

import "testing"

func TestIntPow(t *testing.T) {
	type args struct {
		base     int
		exponent int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "should return correct value for 10*0", args{ 10, 0 }, 1 },
		{ "should return correct value for 10*1", args{ 10, 1 }, 10 },
		{ "should return correct value for 10*2", args{ 10, 2 }, 100 },
		{ "should return correct value for 2*4", args{ 2, 4 }, 16 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntPow(tt.args.base, tt.args.exponent); got != tt.want {
				t.Errorf("IntPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
