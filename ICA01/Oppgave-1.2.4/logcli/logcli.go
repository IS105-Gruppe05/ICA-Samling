package main

import (
  "math"
  "os"
  "strconv"
  "fmt"
)

func main(){
  number, _ := strconv.ParseFloat(os.Args[1], 64)

    fmt.Printf("That number is %f \n", math.Log2(number))
}
