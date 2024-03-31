package main

import (
	"math/big"
	"testing"
)

// TestKaratsuba tests the karatsuba function and compare with the direct multiplication result.
func TestKaratsuba(t *testing.T) {
    testCases := []struct {
        xStr, yStr   string
    }{
        {"1234", "5678"},
        {"8765456782562562456245290", "6545672562562562589876"},
        {"9876545678952562", "9876544567898765432123453663367"},
        {"98765456789524949483838562", "9876544567898765432120049820289493453663367"},
        {"3141592653589793238462643383279502884197169399375105820974944592", "2718281828459045235360287471352662497757247093699959574966967627"},
    }

    for _, tc := range testCases {
        x, _ := new(big.Int).SetString(tc.xStr, 10)
        y, _ := new(big.Int).SetString(tc.yStr, 10)
        
        result := karatsuba(x, y).String()
		directMultiplication := new(big.Int).Mul(x,y).String()

        if result != directMultiplication {
            t.Errorf("Karatsuba(%s, %s) = %s; want %s", tc.xStr, tc.yStr, result, directMultiplication)
        }
    }
}
