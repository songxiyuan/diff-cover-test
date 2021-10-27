package main

import "testing"

func Test_print(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		//{
		//	name: "Print 1",
		//	args: args{
		//		n: 1,
		//	},
		//},
		//{
		//	name: "Print 4",
		//	args: args{
		//		n: 4,
		//	},
		//},
		{
			name: "Print 2",
			args: args{
				n: 2,
			},
		},
		{
			name: "Print 3",
			args: args{
				n: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.args.n)
		})
	}
}
