package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "нулевой размер", size: 0, wantLen: 0},
		{name: "отрицательный размер", size: -10, wantLen: 0},
		{name: "один элемент", size: 1, wantLen: 1},
		{name: "много элементов", size: 10_000, wantLen: 10_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if got == nil {
				t.Fatal("ожидался слайс, получен nil")
			}
			if len(got) != tt.wantLen {
				t.Fatalf("длина слайса: ожидалось %d, получено %d", tt.wantLen, len(got))
			}
			for i, value := range got {
				if value <= 0 {
					t.Fatalf("элемент %d: ожидалось положительное число, получено %d", i, value)
				}
			}
		})
	}
}

// maxTestCases — общие случаи для maximum и maxChunks:
// обе функции должны возвращать одинаковый результат.
var maxTestCases = []struct {
	name string
	data []int
	want int
}{
	{name: "nil-слайс", data: nil, want: 0},
	{name: "пустой слайс", data: []int{}, want: 0},
	{name: "один элемент", data: []int{42}, want: 42},
	{name: "все элементы равны", data: []int{7, 7, 7, 7, 7, 7, 7, 7, 7}, want: 7},
	{name: "максимум в начале", data: []int{100, 1, 2, 3, 4, 5, 6, 7, 8, 9}, want: 100},
	{name: "максимум в конце", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}, want: 100},
	{name: "максимум в середине", data: []int{1, 2, 3, 4, 100, 5, 6, 7, 8, 9}, want: 100},
	{name: "элементов меньше, чем частей", data: []int{3, 9, 2}, want: 9},
	{name: "элементов ровно столько, сколько частей", data: []int{1, 2, 3, 99, 5, 6, 7, 8}, want: 99},
	// 17 элементов: при делении на 8 частей последний элемент
	// попадает в остаток и не должен потеряться.
	{name: "максимум в остатке после деления на части", data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 500}, want: 500},
}

func TestMaximum(t *testing.T) {
	for _, tt := range maxTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	for _, tt := range maxTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	for _, size := range []int{1, 7, 8, 9, 1_000, 100_003} {
		data := generateRandomElements(size)

		want := maximum(data)
		if got := maxChunks(data); got != want {
			t.Errorf("размер %d: maxChunks() = %d, maximum() = %d", size, got, want)
		}
	}
}
