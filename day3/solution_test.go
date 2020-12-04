package day3

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_loadData(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{
			name: "a",
			args: args{
				reader: func() io.Reader {
					return strings.NewReader(".#")
				}(),
			},
			want: [][]bool{{false, true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadData(tt.args.reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadData() = %v, want %v", got, tt.want)
			}
		})
	}
}
