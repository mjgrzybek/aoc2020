package day9

import (
	"reflect"
	"testing"
)

func TestCyclicBuffer_Push(t *testing.T) {
	type fields struct {
		buffer  []int
		lastIdx int
		size    int
	}
	type args struct {
		val []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				buffer:  make([]int, 6, 6),
				lastIdx: 0,
				size:    6,
			},
			args: args{
				val: []int{2},
			},
		},
		{
			name: "",
			fields: fields{
				buffer:  make([]int, 6, 6),
				lastIdx: 0,
				size:    6,
			},
			args: args{
				val: []int{1, 2, 3, 4, 5},
			},
		},
		{
			name: "",
			fields: fields{
				buffer:  make([]int, 6, 6),
				lastIdx: 0,
				size:    6,
			},
			args: args{
				val: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &CyclicBuffer{
				buffer:  tt.fields.buffer,
				lastIdx: tt.fields.lastIdx,
				size:    tt.fields.size,
			}

			for _, v := range tt.args.val {
				b.Push(v)
			}

			for i, expected := range tt.args.val[Max(len(tt.args.val)-b.size, 0):] {
				actual := b.Get(i % b.size)
				if actual != expected {
					t.Fail()
				}
			}
		})
	}
}

func TestSortedSlice_Update(t *testing.T) {
	type fields struct {
		Array []int
	}
	type args struct {
		removed int
		added   int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected []int
	}{
		{
			name: "",
			fields: fields{
				Array: []int{0, 1, 2, 3, 4, 5, 6, 7},
			},
			args: args{
				removed: 3,
				added:   5,
			},
			expected: []int{0, 1, 2, 4, 5, 5, 6, 7},
		},
		{
			name: "remove nonexisting",
			fields: fields{
				Array: []int{0, 1, 2, 3, 4, 5, 6, 7},
			},
			args: args{
				removed: 10,
				added:   5,
			},
			expected: []int{0, 1, 2, 3, 4, 5, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sum2Map{
				Array: tt.fields.Array,
			}

			s.Update(tt.args.removed, tt.args.added)

			if reflect.DeepEqual(s.Array, tt.expected) == false {
				t.Fail()
			}
		})
	}
}
