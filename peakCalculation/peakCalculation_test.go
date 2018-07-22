package peakCalculation

import (
	"testing"
	"fmt"
)

func TestGetHighestPrice(t *testing.T) {
	result1 := GetHiLowPrice(12,1,2)
	test := 0.0
	for _,v:= range result1 {
		test += v
	}
	result2 := GetHiLowPrice(12,1,2)
	fmt.Println(test/10)
	fmt.Println(result1)
	fmt.Println(result2)
}
