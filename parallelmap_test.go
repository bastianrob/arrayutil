package arrayutil

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestParallelMap(t *testing.T) {
	type args struct {
		arr       interface{}
		transform interface{}
		t         reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "Argument is not an array",
			args:    args{arr: 1, transform: nil, t: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Transform function is nil",
			args:    args{arr: []int{1, 2, 3}, transform: nil, t: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Transform is not a function",
			args:    args{arr: []int{1, 2, 3}, transform: 1, t: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "T is not supplied",
			args:    args{arr: []int{1, 2, 3}, transform: 1, t: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid transform",
			args: args{arr: []int{1, 2, 3}, transform: func(num int) int {
				return num + 1
			}, t: reflect.TypeOf(1)},
			want:    []int{2, 3, 4},
			wantErr: false,
		},
		{
			name: "Valid transform",
			args: args{arr: []int{1, 2, 3}, transform: func(num int) string {
				return strconv.Itoa(num)
			}, t: reflect.TypeOf("")},
			want:    []string{"1", "2", "3"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParallelMap(tt.args.arr, tt.args.transform, tt.args.t)
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

func BenchmarkParallelMap(b *testing.B) {
	source := [100]int{}
	for i := 0; i < len(source); i++ {
		source[i] = i + 1
	}
	transf := func(num int) int {
		//fake, long running operations
		time.Sleep(20 * time.Millisecond)
		return num + 1
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ParallelMap(source, transf, reflect.TypeOf(1))
	}
}

func BenchmarkImperative(b *testing.B) {
	source := [100]int{}
	for i := 0; i < len(source); i++ {
		source[i] = i + 1
	}
	transf := func(num int) int {
		//fake, long running operations
		time.Sleep(20 * time.Millisecond)
		return num + 1
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, num := range source {
			transf(num)
		}
	}
}
