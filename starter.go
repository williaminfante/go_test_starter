package starter

import (
	"fmt"
	"math"
	"net/http"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

func OddOrEven(num int) string {
	criteria := math.Mod(float64(num), 2)
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", num)
	}
	return fmt.Sprintf("%v is an even number", num)
}

func Checkhealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "health check passed")
}
