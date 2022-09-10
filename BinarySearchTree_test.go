package stl4go

import "testing"

func Test_isLess(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testing",
			args: args{
				a: 1,
				b: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLess(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isLess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEqual(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testing",
			args: args{
				a: 2,
				b: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
