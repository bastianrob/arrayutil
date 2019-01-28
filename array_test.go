package arrayutil

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type anon struct {
		prop string
	}
	anon1 := anon{prop: "anon 1"}
	anon2 := anon{prop: "anon 2"}

	type args struct {
		arr    interface{}
		clause interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "String exists",
			args: args{
				arr:    []interface{}{"this", "that"},
				clause: "this",
			},
			want: true,
		},
		{
			name: "Int exists",
			args: args{
				arr:    []interface{}{1, 2},
				clause: 2,
			},
			want: true,
		},
		{
			name: "Referred object exists",
			args: args{
				arr:    []interface{}{anon1, anon2},
				clause: anon2,
			},
			want: true,
		},
		{
			name: "Copy object exists",
			args: args{
				arr:    []interface{}{anon1, anon2},
				clause: anon{prop: "anon 1"},
			},
			want: true,
		},
		{
			name: "Does not exists",
			args: args{
				arr:    []interface{}{anon1, anon2},
				clause: anon{prop: "not anon"},
			},
			want: false,
		},
		{
			name: "Argument is not an array",
			args: args{arr: "something", clause: "search"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.arr, tt.args.clause); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type Person struct {
		Name       string
		Birthplace string
	}
	type PersonGroup map[string][]string
	type args struct {
		arr          interface{}
		initialValue interface{}
		transform    interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "Input must be an array",
			args:    args{arr: "something"},
			wantErr: true,
		},
		{
			name:    "Transform must not be nil",
			args:    args{arr: []int{1, 2, 3}, transform: nil},
			wantErr: true,
		},
		{
			name:    "Transform must be a function",
			args:    args{arr: []int{1, 2, 3}, transform: "something"},
			wantErr: true,
		},
		{
			name: "Sum of array",
			args: args{
				arr:          []int{1, 2, 3},
				initialValue: 0,
				transform: func(accumulator, entry, idx int) int {
					return accumulator + entry
				},
			},
			wantErr: false,
			want:    6,
		},
		{
			name: "Group by person's name",
			args: args{
				arr: []Person{
					Person{"John Doe", "Jakarta"},
					Person{"John Doe", "Depok"},
					Person{"John Doe", "Medan"},
				},
				initialValue: make(PersonGroup),
				transform: func(accumulator PersonGroup, entry Person, idx int) PersonGroup {
					birthplaces, exists := accumulator[entry.Name]
					if !exists {
						birthplaces = []string{entry.Birthplace}
					} else {
						birthplaces = append(birthplaces, entry.Birthplace)
					}
					accumulator[entry.Name] = birthplaces
					return accumulator
				},
			},
			wantErr: false,
			want:    PersonGroup{"John Doe": []string{"Jakarta", "Depok", "Medan"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reduce(tt.args.arr, tt.args.initialValue, tt.args.transform)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reduce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
