package main

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		operators    []operator
		concatActive bool
	}
	type want struct {
		result    bool
		operators []operator
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "single",
			args: args{
				operators: []operator{ADD},
			},
			want: want{
				operators: []operator{MULT},
				result:    true,
			},
		},
		{
			name: "single end of cycle",
			args: args{
				operators: []operator{MULT},
			},
			want: want{
				operators: []operator{ADD},
				result:    false,
			},
		},
		{
			name: "double no carry",
			args: args{
				operators: []operator{ADD, ADD},
			},
			want: want{
				operators: []operator{MULT, ADD},
				result:    true,
			},
		},
		{
			name: "double carry",
			args: args{
				operators: []operator{MULT, ADD},
			},
			want: want{
				operators: []operator{ADD, MULT},
				result:    true,
			},
		},
		{
			name: "double v2",
			args: args{
				operators: []operator{ADD, MULT},
			},
			want: want{
				operators: []operator{MULT, MULT},
				result:    true,
			},
		},
		{
			name: "double carry carry",
			args: args{
				operators: []operator{MULT, MULT},
			},
			want: want{
				operators: []operator{ADD, ADD},
				result:    false,
			},
		},
		{
			name: "triple",
			args: args{
				operators: []operator{MULT, MULT, ADD},
			},
			want: want{
				operators: []operator{ADD, ADD, MULT},
				result:    true,
			},
		},
		{
			name: "triple with concat",
			args: args{
				concatActive: true,
				operators:    []operator{MULT, MULT, ADD},
			},
			want: want{
				operators: []operator{CONCAT, MULT, ADD},
				result:    true,
			},
		},
		{
			name: "triple with concat 2",
			args: args{
				concatActive: true,
				operators:    []operator{CONCAT, MULT, ADD},
			},
			want: want{
				operators: []operator{ADD, CONCAT, ADD},
				result:    true,
			},
		},
		{
			name: "triple with concat 3",
			args: args{
				concatActive: true,
				operators:    []operator{ADD, CONCAT, ADD},
			},
			want: want{
				operators: []operator{MULT, CONCAT, ADD},
				result:    true,
			},
		},
		{
			name: "triple with concat 4",
			args: args{
				concatActive: true,
				operators:    []operator{MULT, CONCAT, ADD},
			},
			want: want{
				operators: []operator{CONCAT, CONCAT, ADD},
				result:    true,
			},
		},
		{
			name: "triple with concat 5",
			args: args{
				concatActive: true,
				operators:    []operator{CONCAT, CONCAT, ADD},
			},
			want: want{
				operators: []operator{ADD, ADD, MULT},
				result:    true,
			},
		},
		{
			name: "triple with concat Z",
			args: args{
				concatActive: true,
				operators:    []operator{CONCAT, CONCAT, CONCAT},
			},
			want: want{
				operators: []operator{ADD, ADD, ADD},
				result:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPermutation(tt.args.operators, tt.args.concatActive); got != tt.want.result {
				t.Errorf("nextPermutation() = %v, want %v", got, tt.want)
			}
			if eq := reflect.DeepEqual(tt.args.operators, tt.want.operators); !eq {
				t.Errorf("nextPermutation() array incorrect, got %v, want %v", tt.args.operators, tt.want.operators)
			}
		})
	}
}

func Test_sumValidCalibrationData(t *testing.T) {
	type args struct {
		dataList     []calibrationData
		concatActive bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				concatActive: false,
				dataList: []calibrationData{
					{result: 190, numbers: []int{10, 19}},
					{result: 3267, numbers: []int{81, 40, 27}},
					{result: 83, numbers: []int{17, 5}},
					{result: 156, numbers: []int{15, 6}},
					{result: 7290, numbers: []int{6, 8, 6, 15}},
					{result: 161011, numbers: []int{16, 10, 13}},
					{result: 192, numbers: []int{17, 8, 14}},
					{result: 21037, numbers: []int{9, 7, 18, 13}},
					{result: 292, numbers: []int{11, 6, 16, 20}},
				},
			},
			want: 3749,
		},
		{
			name: "example",
			args: args{
				concatActive: true,
				dataList: []calibrationData{
					{result: 190, numbers: []int{10, 19}},
					{result: 3267, numbers: []int{81, 40, 27}},
					{result: 83, numbers: []int{17, 5}},
					{result: 156, numbers: []int{15, 6}},
					{result: 7290, numbers: []int{6, 8, 6, 15}},
					{result: 161011, numbers: []int{16, 10, 13}},
					{result: 192, numbers: []int{17, 8, 14}},
					{result: 21037, numbers: []int{9, 7, 18, 13}},
					{result: 292, numbers: []int{11, 6, 16, 20}},
				},
			},
			want: 11387,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumValidCalibrationData(tt.args.dataList, tt.args.concatActive); got != tt.want {
				t.Errorf("sumValidCalibrationData() = %v, want %v", got, tt.want)
			}
		})
	}
}
