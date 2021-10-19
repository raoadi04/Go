package fields

import (
	"reflect"
	"testing"
)

func TestNeoSplit(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Test 0", args: args{s: "A B C D\nE", n: 0}, want: []string{"A B C D\nE"}},
		{name: "Test 1", args: args{s: "A B C D\nE", n: 1}, want: []string{"A", "B C D\nE"}},
		{name: "Test 2", args: args{s: "A B C D\nE", n: 2}, want: []string{"A", "B", "C D\nE"}},
		{name: "Test 3", args: args{s: "A B C D\nE", n: 3}, want: []string{"A", "B", "C", "D\nE"}},
		{name: "Test 4", args: args{s: "A B C D\nE", n: 4}, want: []string{"A", "B", "C", "D", "E"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldsN(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldsN() = %v, want %v", got, tt.want)
			}
		})
	}
}
