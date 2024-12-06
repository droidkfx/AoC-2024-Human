package main

import "testing"

func Test_countMiddleOfCorrectOrders(t *testing.T) {
	type args struct {
		rules  map[int][]int
		orders [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				rules: map[int][]int{
					47: {53, 13, 61, 29},
					97: {13, 61, 47, 29, 53, 75},
					75: {29, 53, 47, 61, 13},
					61: {13, 53, 29},
					29: {13},
					53: {29, 13},
				},
				orders: [][]int{
					{75, 47, 61, 53, 29},
					{97, 61, 53, 29, 13},
					{75, 29, 13},
					{75, 97, 47, 61, 53},
					{61, 13, 29},
					{97, 13, 75, 29, 47},
				},
			},
			want: 143,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMiddleOfCorrectOrders(tt.args.rules, tt.args.orders); got != tt.want {
				t.Errorf("countMiddleOfCorrectOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countMiddleOfCorrectedOrders(t *testing.T) {
	type args struct {
		rules  map[int][]int
		orders [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				rules: map[int][]int{
					47: {53, 13, 61, 29},
					97: {13, 61, 47, 29, 53, 75},
					75: {29, 53, 47, 61, 13},
					61: {13, 53, 29},
					29: {13},
					53: {29, 13},
				},
				orders: [][]int{
					{75, 47, 61, 53, 29},
					{97, 61, 53, 29, 13},
					{75, 29, 13},
					{75, 97, 47, 61, 53},
					{61, 13, 29},
					{97, 13, 75, 29, 47},
				},
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMiddleOfCorrectedOrders(tt.args.rules, tt.args.orders); got != tt.want {
				t.Errorf("countMiddleOfCorrectedOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
