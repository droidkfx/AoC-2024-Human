package main

import (
	"slices"
	"testing"
)

func Test_calculateDistance(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example case",
			args: args{
				l1: []int{3, 4, 2, 1, 3, 3},
				l2: []int{4, 3, 5, 3, 9, 3},
			},
			want: 11,
		},
		{
			name: "example case invert",
			args: args{
				l1: []int{4, 3, 5, 3, 9, 3},
				l2: []int{3, 4, 2, 1, 3, 3},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateDistance(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("calculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_similarityScore(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example case",
			args: args{
				l1: []int{3, 4, 2, 1, 3, 3},
				l2: []int{4, 3, 5, 3, 9, 3},
			},
			want: 31,
		},
		{
			name: "example case invert",
			args: args{
				l1: []int{4, 3, 5, 3, 9, 3},
				l2: []int{3, 4, 2, 1, 3, 3},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slices.Sort(tt.args.l1)
			slices.Sort(tt.args.l2)
			if got := similarityScore(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("similarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
