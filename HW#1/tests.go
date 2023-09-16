package main

import (
	"fmt"
	"testing"
)

func main(t *testing.T) {
	type args struct {
		a []int
		b []int
	}

	tests := []struct{
		name string
		args args
		want bool
	}{
		{
			"two equal slices", 
			args{
				a: []int{1,2,3},
				b: []int{1,2,3},
			}, 
			true
		},
		{
			"now equal slices", 
			args{
				a: []int{2,2,3},
				b: []int{1,2,3},
			}, 
			true
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareTwoSlices(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("compareTwoSlices() = %v, want %v" got, tt.want)
			}
		})
	}
}
