package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  r := float64(x/2)
  for i := 0; i < 10; i++{
    r = r -(r*r-x)/(2*r)
  }
  return r
}

func SqrtDelta(x float64) float64 {
  r := float64(x/2)
  limDelta := float64(1e-6)
  delta := float64(0);
  for i := 0; i < 10000; i++{
    delta = r - (r -(r*r-x)/(2*r))
    if (delta < limDelta){
      fmt.Println("Número de ciclos con el delta: ", i)
      break
    }
    r = r -(r*r-x)/(2*r)
  }
  return r
}

func main() {
  fmt.Println("resultado con delta: ",SqrtDelta(10))
  fmt.Println("resultado con 10 ciclos: ",Sqrt(10))
  fmt.Println("resultado con la libreria math: ",math.Sqrt(10))
}
