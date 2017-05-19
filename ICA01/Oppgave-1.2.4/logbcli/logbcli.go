package main

import (
  "math"
  "os"
  "strconv"
  "fmt"
)

func main(){
  base, _ := strconv.ParseFloat(os.Args[1], 64)
  number, _ := strconv.ParseFloat(os.Args[2], 64)

    fmt.Printf("That number is %f \n", math.Log(number) / math.Log(base))
}
