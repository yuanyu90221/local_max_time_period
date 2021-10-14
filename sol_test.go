package sol

import (
	"reflect"
	"testing"
)

func TestFindBestInOut(t *testing.T) {
	type args struct {
		history []int
	}
	tests := []struct {
		name string
		args args
		want ReachDataFormat
	}{
		{
			name: "[60, 100, 58, 20, 56, 42, 20, 250, 194, 10, 1000]",
			args: args{history: []int{60, 100, 58, 20, 56, 42, 20, 250, 194, 10, 1000}},
			want: ReachDataFormat{
				990,
				9,
				10,
			},
		},
		{
			name: "[60, 100, 58, 20, 56, 42, 40, 250, 194, 1000, 10]",
			args: args{history: []int{60, 100, 58, 20, 56, 42, 40, 250, 194, 1000, 10}},
			want: ReachDataFormat{
				980,
				3,
				9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBestInOut(tt.args.history); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBestInOut() = %v, want %v", got, tt.want)
			}
		})
	}
}
