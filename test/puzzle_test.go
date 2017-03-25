package test

import (
	"fmt"
	"testing"

	"github.com/toshim45/jajak/handlers"
)

func TestCalculate(t *testing.T) {
	p := handlers.NewPuzzleHandler()
	result := p.Calculate("artikow")
	fmt.Print(result)
	if result == "" {
		t.Errorf("wrong result")
	}
}
