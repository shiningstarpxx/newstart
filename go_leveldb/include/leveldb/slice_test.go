/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: slice_test.go
 * @Version: 1.0.0
 * @Data: 2022/3/12 4:10 PM
 */
package leveldb

import (
	"reflect"
	"testing"
)

func TestNewEmptySlice(t *testing.T) {
	tests := []struct {
		name string
		want *LSlice
	}{
		{
			"empty slice",
			&LSlice{0, ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmptySlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSlice(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want *LSlice
	}{
		{
			"sssss-5",
			args{"sssss", 5},
			&LSlice{5, "sssss"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSlice(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSliceWithStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *LSlice
	}{
		{
			"sssss-5",
			args{"sssss"},
			&LSlice{5, "sssss"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSliceWithStr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSliceWithStr() = %v, want %v", got, tt.want)
			}
		})
	}
}