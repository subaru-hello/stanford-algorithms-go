package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
    fmt.Println("Hello, ")
    x,y := 87654567890, 65456789876;
    // x,y := 123, 4567;
    // x,y := 1234, 5678;
    result := karatsuba(x, y)
    fmt.Println("Result:", result)
    fmt.Println("Result2:", x * y)
}



func getMaxLength(x,y int) int {
  if(x > y) {
    return x
  }
  return y
}

func karatsuba(x,y int) int{
    // base point
    if(x < 10 || y < 10){
        return x * y;
    }
    // get max length
    n := getMaxLength(len(strconv.Itoa(x)), len(strconv.Itoa(y)))/2

    firstHalfOfX := x / int(math.Pow(10, float64(n)))
    lastHalfOfX := x % int(math.Pow(10, float64(n)))
    firstHalfOfY := y / int(math.Pow(10, float64(n)))
    lastHalfOfY := y % int(math.Pow(10, float64(n)))

    stepOne := karatsuba(firstHalfOfX,firstHalfOfY)
    stepTwo := karatsuba(lastHalfOfX,lastHalfOfY)
    stepFour := karatsuba(firstHalfOfX, lastHalfOfY) + karatsuba(lastHalfOfX, firstHalfOfY)
    // star function = ac * 10 ** n + bd + 10 ** n/2 * (ad + bc)
    result := stepOne * int(math.Pow(10,float64(n * 2))) + stepTwo + stepFour * int(math.Pow(10, float64(n)))
return result;

}