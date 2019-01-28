package arrayutil

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		arr     interface{}
		filterf FilterFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []interface{}
	}{
		{"Success", args{
			arr: []int{1, 2, 3, 4},
			filterf: func(entry interface{}) bool {
				return entry == 1
			}}, false, []interface{}{1}},
		{"Failed", args{
			arr:     "[]int{1, 2, 3, 4}",
			filterf: nil}, true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Filter(tt.args.arr, tt.args.filterf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Filter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
