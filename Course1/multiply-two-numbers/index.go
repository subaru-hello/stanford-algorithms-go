package main

import (
	"fmt"
	"math/big"
)

func main() {
    fmt.Println("Hello, ")
    xStr,yStr := "8765456782562562456245290", "6545672562562562589876";
    // x,y := 123, 4567;
    x,_ := new(big.Int).SetString(xStr, 10);
    y,_ := new(big.Int).SetString(yStr, 10);;
    result := karatsuba(x, y)
    fmt.Println("KaratsubaResult:", result.String())
    fmt.Println("DirectMultiplicationResult:", new(big.Int).Mul(x, y));
}


// if the sizes of n is an odd, it should extract a bigger of half of n
func getMaxLength(x,y *big.Int) int {
    xLen := len(x.String());
    yLen := len(y.String());

  if(xLen > yLen) {
    return xLen
  }
  return yLen
}

// use pointer so as not to eat redundant memories by copying passed numbers(x,y) each times;
func karatsuba(x,y *big.Int) *big.Int{
    ten := big.NewInt(10)
    // base case
    if(x.Cmp(ten) < 0 || y.Cmp(ten) < 0){
        return new(big.Int).Mul(x , y);
    }
    // get max length
    n := getMaxLength(x,y)
    halfN := n /2

    // calculate 10^halfN
    tenToHalfN := new(big.Int).Exp(ten, big.NewInt(int64(halfN)), nil)

    firstHalfOfX :=  new(big.Int).Div(x, tenToHalfN)
    lastHalfOfX := new(big.Int).Mod(x, tenToHalfN)
    firstHalfOfY := new(big.Int).Div(y, tenToHalfN)
    lastHalfOfY :=  new(big.Int).Mod(y, tenToHalfN)

    stepOne := karatsuba(firstHalfOfX,firstHalfOfY)
    stepTwo := karatsuba(lastHalfOfX,lastHalfOfY)
    stepFour := new(big.Int).Add(karatsuba(firstHalfOfX, lastHalfOfY) , karatsuba(lastHalfOfX, firstHalfOfY))
    
    // star function = ac * 10^(2*halfN) + bd + 10^halfN * (ad + bc)
    result := new(big.Int).Mul(stepOne, new(big.Int).Exp(ten, big.NewInt(int64(2*halfN)), nil))
    result.Add(result, new(big.Int).Mul(stepFour, tenToHalfN)).Add(result, stepTwo)
return result;

}