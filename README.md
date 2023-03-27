# Golang Topics

This is a repo to store my golang progress on golang.

I wish someday I'd be able to work with this great language

- Class annotations:
  - Switches allows us to evaluate expressions and constant values
  
    ```go
    package main

    import "fmt"

    func main() {
      a := 1

      switch a { // can only use const expressions to evaluate
      case 10:
        fmt.Println("equals to 10")
      }

      switch { // can use dynamic expressions to evaluate
      case a < 10:
        fmt.Println("less than 10")
      case a == 10:
        fmt.Println("equals to 10")
      default:
        fmt.Println("more than 10")
      }

    }
    ```
    
  - Maps
    
    
    ```go
    
    package main
    import "fmt"

    func main() {
      myMap := make(map[string]int)

      myMap["aluno0"] = 10
      myMap["aluno1"] = 40
      myMap["aluno3"] = 50
      myMap["aluno2"] = 60

      for key, value := range myMap { // it doesn't prints sorted
        fmt.Printf("%s: %d\n", key, value)
      }

      fmt.Println(myMap["aluno5"]) // prints 0
      fmt.Println(myMap)
    }
    
    ```
