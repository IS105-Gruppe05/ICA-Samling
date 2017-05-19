package log

import (
  "math"
  "os"
  "strconv"
  "fmt"
)

func Log2(){
  number, _ := strconv.ParseFloat(os.Args[1], 64)

    fmt.Printf("That number is %f \n", math.Log2(number))
}
