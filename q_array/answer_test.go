package q_array

import "testing"

func TestMaxMate(t *testing.T) {
	type args struct {
		inputArray []int32
	}
	tests := []struct {
		name    string
		args    args
		wantMax int32
	}{
		// TODO: Add test cases.
		{
			"max",
			args{
				[]int32{2, 3, -4, 5, 1},
			},
			6,
		}, {
			"max",
			args{
				[]int32{2, -3, -4, 5, 1},
			},
			12,
		}, {
			"max",
			args{
				[]int32{2, 3, -4, 5, 2},
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := MaxMate(tt.args.inputArray); gotMax != tt.wantMax {
				t.Errorf("MaxMate() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
