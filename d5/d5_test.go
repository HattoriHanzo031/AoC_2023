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
		in           singleRange
		wantMapped   []singleRange
		wantUnmapped []singleRange
	}{
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{1, 5},
			wantMapped:   nil,
			wantUnmapped: []singleRange{{1, 5}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{11, 19},
			wantMapped:   []singleRange{{11, 19}},
			wantUnmapped: nil,
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{17, 21},
			wantMapped:   []singleRange{{17, 20}},
			wantUnmapped: []singleRange{{21, 21}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{7, 17},
			wantMapped:   []singleRange{{10, 17}},
			wantUnmapped: []singleRange{{7, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{9, 21},
			wantMapped:   []singleRange{{10, 20}},
			wantUnmapped: []singleRange{{9, 9}, {21, 21}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{10, 20},
			wantMapped:   []singleRange{{10, 20}},
			wantUnmapped: nil,
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{1, 9},
			wantMapped:   nil,
			wantUnmapped: []singleRange{{1, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{1, 10},
			wantMapped:   []singleRange{{10, 10}},
			wantUnmapped: []singleRange{{1, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{20, 25},
			wantMapped:   []singleRange{{20, 20}},
			wantUnmapped: []singleRange{{21, 25}},
		},
		{
			name:         "",
			args:         args{10, 20, 0},
			in:           singleRange{21, 25},
			wantMapped:   nil,
			wantUnmapped: []singleRange{{21, 25}},
		},

		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{1, 5},
			wantMapped:   nil,
			wantUnmapped: []singleRange{{1, 5}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{11, 19},
			wantMapped:   []singleRange{{21, 29}},
			wantUnmapped: nil,
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{17, 21},
			wantMapped:   []singleRange{{27, 30}},
			wantUnmapped: []singleRange{{21, 21}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{7, 17},
			wantMapped:   []singleRange{{20, 27}},
			wantUnmapped: []singleRange{{7, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{9, 21},
			wantMapped:   []singleRange{{20, 30}},
			wantUnmapped: []singleRange{{9, 9}, {21, 21}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{10, 20},
			wantMapped:   []singleRange{{20, 30}},
			wantUnmapped: nil,
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{1, 9},
			wantMapped:   nil,
			wantUnmapped: []singleRange{{1, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{1, 10},
			wantMapped:   []singleRange{{20, 20}},
			wantUnmapped: []singleRange{{1, 9}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{20, 25},
			wantMapped:   []singleRange{{30, 30}},
			wantUnmapped: []singleRange{{21, 25}},
		},
		{
			name:         "",
			args:         args{10, 20, 10},
			in:           singleRange{21, 25},
			wantMapped:   nil,
			wantUnmapped: []singleRange{{21, 25}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMapped, gotUnmapped := mappingRangesFn(tt.args.from, tt.args.to, tt.args.offset)(tt.in)
			assert.Equal(t, tt.wantMapped, gotMapped)
			assert.Equal(t, tt.wantUnmapped, gotUnmapped)
			// t.Log(gotMapped, gotUnmapped)
			// t.Fail()
		})
	}
}
