package binary

import "testing"

type comparison struct {
	name string
	age  int
}

// Less is a Less interface implementation
func (c *comparison) Less() func(l, r interface{}) bool {
	return func(l, r interface{}) bool {
		lhs := l.(comparison)
		rhs := r.(comparison)
		return lhs.age < rhs.age
	}
}

var comparisonLess = (&comparison{}).Less()

// TestSearch is test for binary search function
func TestSearch(t *testing.T) {
	type args struct {
		input interface{}
		elem  interface{}
		less  func(l, r interface{}) bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				input: []int{6, 7, 8, 9},
				elem:  8,
				less:  IntLess(),
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "not found int",
			args: args{
				input: []int{6, 7, 8, 9},
				elem:  10,
				less:  IntLess(),
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "one to ten",
			args: args{
				input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				elem:  2,
				less:  IntLess(),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "one to ten strings",
			args: args{
				input: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
				elem:  "6",
				less:  StringLess(),
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "comparison",
			args: args{
				input: []comparison{
					{
						name: "qwer",
						age:  1,
					},
					{
						name: "ty",
						age:  2,
					},
					{
						name: "asdfg",
						age:  3,
					},
					{
						name: "zxcvbn",
						age:  4,
					},
				},
				elem: comparison{
					name: "ty",
					age:  2,
				},
				less: comparisonLess,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Search(tt.args.input, tt.args.elem, tt.args.less)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}
