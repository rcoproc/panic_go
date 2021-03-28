package main
import (
  "fmt"
)

func divide(a,b int) (int, error){
    if b == 0 {
      return 0, fmt.Errorf("Can not divide by zero")
    }

    return (a/b), nil
}


func main() {
    a, b := 10, -2

    c, err := divide(a,b)

    if err != nil {
      fmt.Println("Erro na divisão", err)
      panic(err)
    }

    fmt.Println("Division is", c)
}
