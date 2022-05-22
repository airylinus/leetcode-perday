package time_loop

import (
	"reflect"
	"testing"
)

func TestAssembleTime(t *testing.T) {
	type args struct {
		inputNums []int32
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				[]int32{2, 1, 3, 3},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssembleTime(tt.args.inputNums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssembleTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
