/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: difference_array_lt1109_test.go
 * @Version: 1.0.0
 * @Data: 2021/8/31 4:15 PM
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_corpFlightBookings(t *testing.T) {
	type args struct {
		bookings [][]int
		n        int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"simple test",
			args{
				[][]int{
					{1, 2, 10}, {2, 3, 20}, {2, 5, 25},
				},
				5,
			},
			[]int{10,55,45,25,25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corpFlightBookings(tt.args.bookings, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("corpFlightBookings() = %v, want %v", got, tt.want)
			}
		})
	}
}
