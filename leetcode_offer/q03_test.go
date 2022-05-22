package leetcode_offer

import "testing"

func TestFindRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "[3, 4, 2, 0, 0, 1]",
			args: args{[]int{3, 4, 2, 0, 0, 1}},
			want: 0,
		},
		//{
		//	name: "simple",
		//	args: args{[]int{0, 1, 3, 3}},
		//	want: 3,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("FindRepeatNumber() = %v, want %v", got, tt.want)
			}
			//Fuckup()
		})
	}
}
