package main

import "testing"

func Test_countXmas(t *testing.T) {
	type args struct {
		data    [][]byte
		counter func(int, int, [][]byte) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example xmas",
			args: args{
				counter: countXmasInstance,
				data: [][]byte{
					{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
					{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
					{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
					{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
					{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
					{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
					{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
					{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
					{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
					{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
				},
			},
			want: 18,
		},
		{
			name: "example xmas",
			args: args{
				counter: countXmasCross,
				data: [][]byte{
					{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
					{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
					{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
					{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
					{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
					{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
					{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
					{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
					{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
					{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanInput(tt.args.data, tt.args.counter); got != tt.want {
				t.Errorf("scanInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
