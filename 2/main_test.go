package main

import "testing"

func Test_calculateSafeReports(t *testing.T) {
	type args struct {
		reports     [][]int
		dampenLevel int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example case",
			args: args{
				dampenLevel: 0,
				reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
			},
			want: 2,
		},
		{
			name: "example case with dampen",
			args: args{
				dampenLevel: 1,
				reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
			},
			want: 4,
		},
		{
			name: "adversarial case",
			args: args{
				dampenLevel: 1,
				reports: [][]int{
					{57, 56, 57, 59, 60, 63, 64, 65},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSafeReports(tt.args.reports, tt.args.dampenLevel); got != tt.want {
				t.Errorf("calculateSafeReports() = %v, want %v", got, tt.want)
			}
		})
	}
}
