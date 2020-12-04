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

func Test_calculateCollisions(t *testing.T) {
	type args struct {
		lines [][]bool
		y     int
		x     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				lines: [][]bool{{false,false,false,false},{true,true,true,true},{false,false,false,false},{true,true,true,true}},
				y:     2,
				x:     1,
			},
			want: 0,
		},
		{
			name: "b",
			args: args{
				lines: [][]bool{{true,false,false,false},{true,true,true,true},{false,false,false,false},{true,true,true,true}},
				y:     2,
				x:     1,
			},
			want: 1,
		},
		{
			name: "c",
			args: args{
				lines: [][]bool{{true,false,false,false},{true,true,true,true},{false,false,false,false},{true,true,true,true},{false,false,true,false}},
				y:     2,
				x:     1,
			},
			want: 2,
		},
		{
			name: "d",
			args: args{
				lines: [][]bool{{false,false,false,false},{true,true,true,true},{false,false,false,false},{true,true,true,true}},
				y:     1,
				x:     1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCollisions(tt.args.lines, tt.args.y, tt.args.x); got != tt.want {
				t.Errorf("calculateCollisions() = %v, want %v", got, tt.want)
			}
		})
	}
}