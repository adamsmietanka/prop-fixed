package main

import (
	"testing"
)

func TestBarycentricX(t *testing.T) {
	type args struct {
		X []float64
		Y []float64
		Z [][]float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "lower left",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 10,
				z: 0,
			},
			want: 0,
		},
		{
			name: "yOutside",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 2,
				z: 0.5,
			},
			want: 0,
		},
		{
			name: "zOutside",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 0,
				z: 9,
			},
			want: 0,
		},
		{
			name: "middle",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 35,
				z: 0.5,
			},
			want: 0.5,
		},
		{
			name: "left edge middle",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 35,
				z: 0.25,
			},
			want: 0,
		},
		{
			name: "right edge middle",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 35,
				z: 0.75,
			},
			want: 1,
		},
		{
			name: "down edge middle",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 10,
				z: 0.25,
			},
			want: 0.5,
		},
		{
			name: "up edge middle",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 60,
				z: 0.75,
			},
			want: 0.5,
		},
		{
			name: "lower middle",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 22.5,
				z: 0.5,
			},
			want: 0.75,
		},
		{
			name: "upper middle",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				y: 47.5,
				z: 0.5,
			},
			want: 0.25,
		},
		{
			name: "big_clipped",
			args: args{
				X: []float64{1, 2},
				Y: []float64{10, 20},
				Z: [][]float64{{0.1, 0.2}, {1.1, 1.2}},
				y: 15,
				z: 0.65,
			},
			want: 1.5,
		},
		{
			name: "big",
			args: args{
				X: []float64{0, 1, 2, 3, 4, 5},
				Y: []float64{10, 20, 30, 40, 50, 60},
				Z: [][]float64{{0, 0.1, 0.2, 0.3, 0.4, 0.5},
					{1, 1.1, 1.2, 1.3, 1.4, 1.5},
					{2, 2.1, 2.2, 2.3, 2.4, 2.5},
					{3, 3.1, 3.2, 3.3, 3.4, 3.5},
					{4, 4.1, 4.2, 4.3, 4.4, 4.5},
					{5, 5.1, 5.2, 5.3, 5.4, 5.5}},
				y: 15,
				z: 0.65,
			},
			want: 1.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BarycentricX(tt.args.X, tt.args.Y, tt.args.Z, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("BarycentricX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBarycentricZ(t *testing.T) {
	type args struct {
		X []float64
		Y []float64
		Z [][]float64
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Basic",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				x: 0,
				y: 10,
			},
			want: 0,
		},
		{
			name: "xOutside",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				x: 2,
				y: 10,
			},
			want: 0,
		},
		{
			name: "yOutside",
			args: args{
				X: []float64{0, 1},
				Y: []float64{10, 60},
				Z: [][]float64{{0, 0.5}, {0.5, 1}},
				x: 0,
				y: 9,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				[]float64{0, 1},
				[]float64{10, 60},
				[][]float64{{0, 0.5}, {0.5, 1}},
				0.5,
				35,
			},
			want: 0.5,
		},
		{
			name: "3",
			args: args{
				[]float64{0, 1},
				[]float64{10, 60},
				[][]float64{{0, 0.5}, {0.5, 1}},
				0,
				35,
			},
			want: 0.25,
		},
		{
			name: "3.5",
			args: args{
				[]float64{0, 1},
				[]float64{10, 60},
				[][]float64{{0, 0.5}, {0.5, 1}},
				0.5,
				10,
			},
			want: 0.25,
		},
		{
			name: "4",
			args: args{
				[]float64{0, 1},
				[]float64{10, 60},
				[][]float64{{0, 0.5}, {0.5, 1}},
				0.5,
				10,
			},
			want: 0.25,
		},
		{
			name: "5",
			args: args{
				[]float64{0, 1},
				[]float64{10, 60},
				[][]float64{{0, 0.5}, {0.5, 1}},
				1,
				60,
			},
			want: 1,
		},
		{
			name: "6pre",
			args: args{
				[]float64{1, 2},
				[]float64{10, 20},
				[][]float64{{0.1, 0.2}, {1.1, 1.2}},
				1.5,
				15,
			},
			want: 0.65,
		},
		{
			name: "6",
			args: args{
				[]float64{0, 1, 2, 3, 4, 5},
				[]float64{10, 20, 30, 40, 50, 60},
				[][]float64{{0, 0.1, 0.2, 0.3, 0.4, 0.5},
					{1, 1.1, 1.2, 1.3, 1.4, 1.5},
					{2, 2.1, 2.2, 2.3, 2.4, 2.5},
					{3, 3.1, 3.2, 3.3, 3.4, 3.5},
					{4, 4.1, 4.2, 4.3, 4.4, 4.5},
					{5, 5.1, 5.2, 5.3, 5.4, 5.5}},
				1.5,
				15,
			},
			want: 0.65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BarycentricZ(tt.args.X, tt.args.Y, tt.args.Z, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("BarycentricZ() = %v, want %v", got, tt.want)
			}
		})
	}
}
