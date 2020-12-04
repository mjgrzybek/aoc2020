package day2

import "testing"

func TestEntry_isValid(t *testing.T) {
	type fields struct {
		Min      int
		Max      int
		Letter   rune
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "a",
			fields: fields{
				Min:      0,
				Max:      1,
				Letter:   'a',
				Password: "a",
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				Min:      0,
				Max:      5,
				Letter:   'a',
				Password: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := entry{
				Min:      tt.fields.Min,
				Max:      tt.fields.Max,
				Letter:   tt.fields.Letter,
				Password: tt.fields.Password,
			}
			if got := e.isValid1(); got != tt.want {
				t.Errorf("isValid1() = %v, want %v", got, tt.want)
			}
		})
	}
}
