package peakCalculation

import (
	"testing"
	"fmt"
)

func TestGetHighestPrice(t *testing.T) {
	result := GetHighestPrice(12)
	fmt.Println(result)
}
