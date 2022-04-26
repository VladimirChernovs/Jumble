package main

import "testing"

func TestFindJumble(t *testing.T) {
	type args struct {
		jumbleArray [][]string
		str         string
	}
	jumbleArray := [][]string{{"A", "B", "C"}, {"X", "Y", "Z"}}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{jumbleArray, "CZ"}, 0, 2},
		{"2", args{jumbleArray, "cz"}, 0, 2},
		{"3", args{jumbleArray, "AB"}, 0, 0},
		{"4", args{jumbleArray, "Bc"}, 0, 1},
		{"5", args{jumbleArray, "BY"}, 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindJumble(tt.args.jumbleArray, tt.args.str)
			if got != tt.want {
				t.Errorf("FindJumble() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindJumble() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
