package tests

import (
	"testing"
	"time"

	"github.com/VadimSavintsev/GoLab5" // Импортируем пакет main
)

func TestRunWorkers(t *testing.T) {
	// Засекаем время начала выполнения
	start := time.Now()

	// Вызываем функцию RunWorkers
	main.RunWorkers(5)

	// Засекаем время окончания выполнения
	elapsed := time.Since(start)

	// Проверяем, что время выполнения не превышает ожидаемое
	expectedDuration := 3 * time.Second
	if elapsed > expectedDuration {
		t.Errorf("RunWorkers took too long: %s, expected less than %s", elapsed, expectedDuration)
	}
}
