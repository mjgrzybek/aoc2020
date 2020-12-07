package day7

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_parsePackage(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 Package
	}{
		{
			name: "",
			args: args{
				s: "2 posh indigo bags",
			},
			want:  2,
			want1: "posh indigo",
		},
		{
			name: "",
			args: args{
				s: "1 posh indigo bag",
			},
			want:  1,
			want1: "posh indigo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parsePackage(tt.args.s)
			if got != tt.want {
				t.Errorf("parsePackage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parsePackage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_loadRelations(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want Package2ContainersMap
	}{
		{
			name: "",
			args: args{
				r: strings.NewReader(`vibrant aqua bags contain 1 shiny gold bag.
drab bronze bags contain 3 vibrant aqua bags, 4 light turquoise bags, 5 light magenta bags, 1 vibrant aqua bag.
vibrant lavender bags contain 2 dim salmon bags, 5 muted violet bags, 5 light magenta bags, 1 vibrant aqua bag.`),
			},
			want: func() Package2ContainersMap {
				p2c := Package2ContainersMap{}
				fillP2C(p2c, "dim salmon", "vibrant lavender")
				fillP2C(p2c, "light magenta", "drab bronze", "vibrant lavender")
				fillP2C(p2c, "light turquoise", "drab bronze")
				fillP2C(p2c, "muted violet", "vibrant lavender")
				fillP2C(p2c, "shiny gold", "vibrant aqua")
				fillP2C(p2c, "vibrant aqua", "drab bronze", "vibrant lavender")
				return p2c
			}(),
		},
		{
			name: "",
			args: args{
				r: strings.NewReader(`a a bags contain 1 b b bag.
b b bags contain 3 c c bags, 4 d d bags, 1 X X bag.
e e bags contain 2 f f bags, 5 g g bags, 2 X X bags.`),
			},
			want: func() Package2ContainersMap {
				p2c := Package2ContainersMap{}
				fillP2C(p2c, "b b", "a a")
				fillP2C(p2c, "c c", "b b")
				fillP2C(p2c, "d d", "b b")
				fillP2C(p2c, "X X", "b b", "e e")
				fillP2C(p2c, "f f", "e e")
				fillP2C(p2c, "g g", "e e")
				return p2c
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadRelations(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadRelations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func fillP2C(p2c Package2ContainersMap, k string, vs ...string) {
	p2c[k] = PackagesSet{}
	for _, v := range vs {
		p2c[k][v] = struct{}{}
	}
}

var predefinedP2C = func() Package2ContainersMap {
	p2c := Package2ContainersMap{}
	fillP2C(p2c, "b b", "a a")
	fillP2C(p2c, "c c", "b b")
	fillP2C(p2c, "d d", "b b")
	fillP2C(p2c, "X X", "b b", "e e")
	fillP2C(p2c, "f f", "e e")
	fillP2C(p2c, "g g", "e e")
	return p2c
}()

func Test_outermost(t *testing.T) {
	type args struct {
		root string
		p2c  Package2ContainersMap
	}
	tests := []struct {
		name string
		args args
		want PackagesSet
	}{
		{
			name: "",
			args: args{
				root: "a a",
				p2c: predefinedP2C,
			},
			want: func() PackagesSet {
				ps := PackagesSet{}
				return ps
			}(),
		},
		{
			name: "",
			args: args{
				root: "b b",
				p2c: predefinedP2C,
			},
			want: func() PackagesSet {
				ps := PackagesSet{}
				ps["a a"] = struct{}{}
				return ps
			}(),
		},
		{
			name: "",
			args: args{
				root: "X X",
				p2c: predefinedP2C,
			},
			want: func() PackagesSet {
				ps := PackagesSet{}
				ps["a a"] = struct{}{}
				ps["b b"] = struct{}{}
				ps["e e"] = struct{}{}
				return ps
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outermost(tt.args.root, tt.args.p2c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outermost() = %v, want %v", got, tt.want)
			}
		})
	}
}