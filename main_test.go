package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Run("positive size", func(t *testing.T) {
		data := generateRandomElements(100)

		if len(data) != 100 {
			t.Fatalf("expected length 100, got %d", len(data))
		}

		for _, number := range data {
			if number < 1 || number > 1_000_000 {
				t.Fatalf("number out of range: %d", number)
			}
		}
	})

	t.Run("zero size", func(t *testing.T) {
		data := generateRandomElements(0)

		if len(data) != 0 {
			t.Fatalf("expected empty slice, got length %d", len(data))
		}
	})

	t.Run("negative size", func(t *testing.T) {
		data := generateRandomElements(-5)

		if len(data) != 0 {
			t.Fatalf("expected empty slice, got length %d", len(data))
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "one element",
			data: []int{42},
			want: 42,
		},
		{
			name: "several elements",
			data: []int{3, 9, 2, 15, 7},
			want: 15,
		},
		{
			name: "negative numbers",
			data: []int{-8, -3, -12, -1},
			want: -1,
		},
		{
			name: "duplicate maximum",
			data: []int{5, 10, 10, 3},
			want: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maximum(test.data)

			if got != test.want {
				t.Errorf("maximum(%v) = %d, want %d",
					test.data, got, test.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "one element",
			data: []int{42},
			want: 42,
		},
		{
			name: "fewer elements than chunks",
			data: []int{3, 8, 2},
			want: 8,
		},
		{
			name: "length divisible by chunks",
			data: []int{1, 2, 3, 4, 5, 6, 7, 99},
			want: 99,
		},
		{
			name: "length not divisible by chunks",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 100},
			want: 100,
		},
		{
			name: "negative numbers",
			data: []int{-8, -3, -12, -1},
			want: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxChunks(test.data)

			if got != test.want {
				t.Errorf("maxChunks(%v) = %d, want %d",
					test.data, got, test.want)
			}
		})
	}
}
