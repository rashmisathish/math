package main

import (
	"fmt"
	"github.com/shopspring/decimal"

	//	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("hello world")
res:=Oper(decimal.NewFromFloat(11.8),decimal.NewFromFloat(12))
fmt.Println(res)
}
func Oper(a,b decimal.Decimal)decimal.Decimal{
	return decimal.Sum(a,b)
}

