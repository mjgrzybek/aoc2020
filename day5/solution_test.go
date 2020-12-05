package day5

import "testing"

func TestBinarySpacePartitioning(t *testing.T) {
	type args struct {
		inputValue string
		loKey      rune
		hiKey      rune
		span       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a",
			args: args{
				inputValue: "FFFFFFF",
				loKey:      'B',
				hiKey:      'F',
			},
			want: 127,
		},
		{
			name: "b",
			args: args{
				inputValue: "FFFFFFB",
				loKey:      'B',
				hiKey:      'F',
			},
			want: 126,
		},
		{
			name: "c",
			args: args{
				inputValue: "BBBBBBB",
				loKey:      'B',
				hiKey:      'F',
			},
			want: 0,
		},
		{
			name: "d",
			args: args{
				inputValue: "BBBBBBF",
				loKey:      'B',
				hiKey:      'F',
			},
			want: 1,
		},
		{
			name: "e",
			args: args{
				inputValue: "BFFFBBF",
				loKey:      'B',
				hiKey:      'F',
			},
			want: 44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySpacePartitioning(tt.args.inputValue, tt.args.loKey, tt.args.hiKey); got != tt.want {
				t.Errorf("BinarySpaceParticioning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoord_getSeatID(t *testing.T) {
	type fields struct {
		Col string
		Row string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "a",
			fields: fields{
				Row: "BFFFBBF",
				Col: "RRR",
			},
			want: 567,
		},
		{
			name: "b",
			fields: fields{
				Row: "FFFBBBF",
				Col: "RRR",
			},
			want: 119,
		},
		{
			name: "c",
			fields: fields{
				Row: "BBFFBBF",
				Col: "RLL",
			},
			want: 820,
		},
		{
			name: "FBFBBFFRLR",
			fields: fields{
				Col: "LR",
				Row: "FBFBBFFR",
			},
			want: 35,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Coord{
				Col: tt.fields.Col,
				Row: tt.fields.Row,
			}
			if got := c.getSeatID(); got != tt.want {
				t.Errorf("getSeatID() = %v, want %v", got, tt.want)
			}
		})
	}
}