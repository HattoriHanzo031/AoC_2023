package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mappingRangesFn(t *testing.T) {
	type args struct {
		from   int
		to     int
		offset int
	}
	tests := []struct {
		name         string
		args         args
		in           Range
		wantMapped   []Range
		wantUnmapped []Range
	}{
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{1, 5},
			wantMapped:   nil,
			wantUnmapped: []Range{{1, 5}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{11, 19},
			wantMapped:   []Range{{11, 19}},
			wantUnmapped: nil,
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{17, 21},
			wantMapped:   []Range{{17, 20}},
			wantUnmapped: []Range{{21, 21}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{7, 17},
			wantMapped:   []Range{{10, 17}},
			wantUnmapped: []Range{{7, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{9, 21},
			wantMapped:   []Range{{10, 20}},
			wantUnmapped: []Range{{9, 9}, {21, 21}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{10, 20},
			wantMapped:   []Range{{10, 20}},
			wantUnmapped: nil,
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{1, 9},
			wantMapped:   nil,
			wantUnmapped: []Range{{1, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{1, 10},
			wantMapped:   []Range{{10, 10}},
			wantUnmapped: []Range{{1, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{20, 25},
			wantMapped:   []Range{{20, 20}},
			wantUnmapped: []Range{{21, 25}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           Range{21, 25},
			wantMapped:   nil,
			wantUnmapped: []Range{{21, 25}},
		},

		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{1, 5},
			wantMapped:   nil,
			wantUnmapped: []Range{{1, 5}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{11, 19},
			wantMapped:   []Range{{21, 29}},
			wantUnmapped: nil,
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{17, 21},
			wantMapped:   []Range{{27, 30}},
			wantUnmapped: []Range{{21, 21}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{7, 17},
			wantMapped:   []Range{{20, 27}},
			wantUnmapped: []Range{{7, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{9, 21},
			wantMapped:   []Range{{20, 30}},
			wantUnmapped: []Range{{9, 9}, {21, 21}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{10, 20},
			wantMapped:   []Range{{20, 30}},
			wantUnmapped: nil,
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{1, 9},
			wantMapped:   nil,
			wantUnmapped: []Range{{1, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{1, 10},
			wantMapped:   []Range{{20, 20}},
			wantUnmapped: []Range{{1, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{20, 25},
			wantMapped:   []Range{{30, 30}},
			wantUnmapped: []Range{{21, 25}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           Range{21, 25},
			wantMapped:   nil,
			wantUnmapped: []Range{{21, 25}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMapped, gotUnmapped := rangeMappingFn(tt.args.from, tt.args.to, tt.args.offset)(tt.in)
			assert.Equal(t, tt.wantMapped, gotMapped)
			assert.Equal(t, tt.wantUnmapped, gotUnmapped)
			// t.Log(gotMapped, gotUnmapped)
			// t.Fail()
		})
	}
}
