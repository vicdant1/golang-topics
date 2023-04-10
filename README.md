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
  - *Reader | *Scanner
    - Reader reads a single value (I have to use a specific function depending on what kind of value I'm trying to read)
      ```go
      
      reader := bufio.NewReader(os.Stdin)
      input, _ := reader.ReadLine("\n")      
      
      ```
    - Scanner can read many values (or a single as well):
      
      ```go
      
      scanner := bufio.NewScanner(os.Stdin)
      
      scanner.Scan()
      
      text := scanner.Text()
      
      ``` 
      
  - Implement something with Flags ([package](https://pkg.go.dev/flag))
  
  - Methods:
  
    ```go
      type Rectangle struct {}
      // This r is called "reciever" || it can be a pointer as well
      func(r Rectangle) printArea() {
        fmt.Println("Test")
      }
    ```
  - Interfaces
  
    - In order to implement this interface my struct should implement all of the methods
    ```go
    type IFigure {
      area() float32
    }  
    ```
    - Go has a concept of empty interface, anything automatically implements it:
    
    ```go
    func main(){
      // implements any type - kindof dynamic typing in go
      var i interface {}
      
      i = 42
      fmt.Println("Value:", i)
      
      i = "Hello"
      fmt.Println("Value:", i)
    }
    
    // can only use type assertion on interfaces
    
    var x interface{} = "foo"

    var s string = x.(string)
    fmt.Println(s)     // "foo"

    s, ok := x.(string)
    fmt.Println(s, ok) // "foo true"

    n, ok := x.(int)
    fmt.Println(n, ok) // "0 false"

    fmt.Println(n.(type))
    
    n = x.(int)        // ILLEGAL
    ```
    
  - Generics
  ```go
    // K may be integer, V may be int, float64, string
    type IndexTable[K int, V int|float64|string] struct {
      ket K
      values []V
    }
    
    // I am passing this explicit type because compiler can't understand it sometimes
    table := IndexTable[int, int] {1, []int {1,2,3,4}}
    
    // In functions:
    
    func DoSomething[T int|float64](x, y T) T{
      return x + y
    }
    
    // I can create an interface containing the types I admit (type set)
    type Number interface {
      int | float64
    }
    
    func DoSomething[Number](x, y T) T{
      return x + y
    }
  ```
  

    
    
