package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/VadimSavintsev/GoLab5/worker"
)

func TestRunWorkers(t *testing.T) {
	// Засекаем время начала выполнения
	start := time.Now()

	// Вызываем функцию RunWorkers
	worker.RunWorkers(5)

	// Засекаем время окончания выполнения
	elapsed := time.Since(start)

	// Выводим время выполнения в консоль
	fmt.Printf("RunWorkers took: %s\n", elapsed)
}
