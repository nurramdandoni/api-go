package resource

import (
	"fmt"
	"math"
)

func Cal(x, y int, z string) {
	var hasil = math.Pow(float64(x*y), 2)
	fmt.Println(hasil)
}
