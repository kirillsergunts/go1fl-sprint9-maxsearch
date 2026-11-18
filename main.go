package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements возвращает слайс длины size, заполненный
// случайными целыми положительными числами из диапазона [1, math.MaxInt32].
// Если size меньше или равен нулю, возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(math.MaxInt32) + 1
	}

	return data
}

// maximum возвращает максимальный элемент слайса.
// Для пустого слайса возвращает 0: все элементы положительные,
// поэтому 0 не может совпасть с настоящим максимумом.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks делит слайс на CHUNKS частей, ищет максимум в каждой
// части в отдельной горутине и возвращает наибольший из найденных.
// Для пустого слайса возвращает 0.
func maxChunks(data []int) int {
	// Если элементов меньше, чем частей, делить нечего:
	// размер части получился бы равным нулю.
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	chunkMaximums := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		// Последняя часть забирает остаток, если длина слайса
		// не делится на CHUNKS без остатка.
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			// Каждая горутина пишет только в свой элемент слайса,
			// поэтому гонки данных нет и мьютекс не нужен.
			chunkMaximums[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}
