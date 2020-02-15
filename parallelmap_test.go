package arrayutil

import (
	"reflect"
	"testing"
)

func TestParallelMap(t *testing.T) {
	type args struct {
		arr       interface{}
		transform interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "Argument is not an array",
			args:    args{arr: 1, transform: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Transform function is nil",
			args:    args{arr: []int{1, 2, 3}, transform: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Transform is not a function",
			args:    args{arr: []int{1, 2, 3}, transform: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid transform",
			args: args{arr: []int{1, 2, 3}, transform: func(num int) int {
				return num + 1
			}},
			want:    []int{2, 3, 4},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParallelMap(tt.args.arr, tt.args.transform)
			if (err != nil) != tt.wantErr {
				t.Errorf("Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
