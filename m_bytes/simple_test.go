package m_bytes

import "testing"

func Test_simple(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "simple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Simple()
		})
	}
}

func TestCountBit(t *testing.T) {
	type args struct {
		n uint32
	}
	tests := []struct {
		name  string
		args  args
		wantC uint32
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				1,
			},
			wantC: 1,
		},
		{
			name: "",
			args: args{
				2,
			},
			wantC: 1,
		},
		{
			name: "",
			args: args{
				3,
			},
			wantC: 2,
		},
		{
			name: "",
			args: args{
				8,
			},
			wantC: 1,
		},
		{
			name: "",
			args: args{
				9,
			},
			wantC: 2,
		}, {
			name: "",
			args: args{
				10,
			},
			wantC: 2,
		},
		{
			name: "",
			args: args{
				11,
			},
			wantC: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := CountBit(tt.args.n); gotC != tt.wantC {
				t.Errorf("CountBit() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
