package slide_window

import "testing"

func TestLongestSubString(t *testing.T) {
	type args struct {
		inString string
	}
	tests := []struct {
		name  string
		args  args
		wantR string
	}{
		// TODO: Add test cases.
		{
			name:  "",
			args:  args{inString: "abcdeeeee"},
			wantR: "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := LongestSubString(tt.args.inString); gotR != tt.wantR {
				t.Errorf("LongestSubString() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
