package arrayutil

import "testing"

func TestIntersect(t *testing.T) {
	type args struct {
		arr1 interface{}
		arr2 interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{{
		name: "intersects",
		args: args{
			arr1: []int{1, 2, 3, 4, 5},
			arr2: []int{6, 1},
		},
		want: true,
	}, {
		name: "does not intersects",
		args: args{
			arr1: []int{1, 2, 3, 4, 5},
			arr2: []int{6, 8},
		},
		want: false,
	}, {
		name: "not an array",
		args: args{
			arr1: "I am not an array",
			arr2: "I too, am not an array",
		},
		want: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
