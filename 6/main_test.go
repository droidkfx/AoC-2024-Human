package main

import "testing"

func Test_countGuardSpaces(t *testing.T) {
	type args struct {
		guardMap [][]bool
		guard    guardPosition
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				guard: guardPosition{x: 4, y: 6, d: Dir_UP},
				guardMap: [][]bool{
					{false, false, false, false, true, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, true},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, true, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, true, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, true, false, false, false /* guard pos */, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, true, false},
					{true, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, true, false, false, false},
				},
			},
			want: 41,
		},
		{
			name: "straight off",
			args: args{
				guard: guardPosition{x: 1, y: 2, d: Dir_UP},
				guardMap: [][]bool{
					{false, false, false},
					{false, false, false},
					{false, false, false},
				},
			},
			want: 3,
		},
		{
			name: "one turn off",
			args: args{
				guard: guardPosition{x: 1, y: 2, d: Dir_UP},
				guardMap: [][]bool{
					{false, true, false},
					{false, false, false},
					{false, false, false},
				},
			},
			want: 3,
		},
		{
			name: "compound turn backtrack",
			args: args{
				guard: guardPosition{x: 1, y: 2, d: Dir_UP},
				guardMap: [][]bool{
					{false, true, false},
					{false, false, true},
					{false, false, false},
				},
			},
			want: 2,
		},
		{
			name: "compound turn backtrack",
			args: args{
				guard: guardPosition{x: 1, y: 2, d: Dir_UP},
				guardMap: [][]bool{
					{false, true, false},
					{false, false, true},
					{false, false, false},
				},
			},
			want: 2,
		},
		{
			name: "double back track",
			args: args{
				guard: guardPosition{x: 2, y: 4, d: Dir_UP},
				guardMap: [][]bool{
					{false, true, false, false, true},
					{false, false, false, false, true},
					{false, false, true, false, false},
					{false, false, false, true, false},
					{true, false, false, false, false},
					{false, false, true, false, false},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, loop := countGuardSpaces(tt.args.guardMap, tt.args.guard); got != tt.want || loop {
				t.Errorf("countGuardSpaces() = %v, %v, want %v, false", got, loop, tt.want)
			}
		})
	}
}

func Test_countPossibleLoops(t *testing.T) {
	type args struct {
		guardMap [][]bool
		guard    guardPosition
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				guard: guardPosition{x: 4, y: 6, d: Dir_UP},
				guardMap: [][]bool{
					{false, false, false, false, true, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, true},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, true, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, true, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, true, false, false, false /* guard pos */, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, true, false},
					{true, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, true, false, false, false},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPossibleLoops(tt.args.guardMap, tt.args.guard); got != tt.want {
				t.Errorf("countPossibleLoops() = %v, want %v", got, tt.want)
			}
		})
	}
}
