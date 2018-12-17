package combinatorics

import (
  "fmt"
)

func ExamplePermutations() {
  ch:=cPermutations([]interface{}{"A", "B", "C"})
  
  for p:=range ch {
    fmt.Println(p)
  }
}

func ExampleDerangements() {
  ch:=Derangements([]interface{}{"A", "B", "C"})
  
  for d:=range ch {
    fmt.Println(d)
  }
}


func ExampleChoose() {
  ch:=Choose(2, []interface{}{"A", "B", "C"})
  
  for c:=range ch {
    fmt.Println(c)
  }
}