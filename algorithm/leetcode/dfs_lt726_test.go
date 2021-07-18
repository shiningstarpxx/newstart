/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  2:37 PM  17/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "testing"

func Test_countOfAtoms(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"simple test",
			args{
				"H2O",
			},
			"H2O",
		},
		{
			"Mg(OH)2",
			args{
				"Mg(OH)2",
			},
			"H2MgO2",
		},
		{
			"hard",
			args{
				"K4(ON(SO3)2)2",
			},
			"K4N2O14S4",
		},
		{
			"bad case",
			args{
				"H11He49NO35B7N46Li20",
			},
			"B7H11He49Li20N47O35",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfAtoms(tt.args.formula); got != tt.want {
				t.Errorf("countOfAtoms() = %v, want %v", got, tt.want)
			}
		})
	}
}
