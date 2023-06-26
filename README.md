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


  - Comparable
    In GO, we can use 'comparable' as an identifier of primitive types that can be both compared and used as generic type.
    
  - Modularization
    In Go, we implement packages as directories:
    package university -> ./university
    things we want to export in package should begin with capital first letter
    
    ```go
    package temperature
    
    // It begins with capital letter
    func CelsiusToFahrenheit(t float64) float64 {
      return (t*9/5) + 32
    }
    
    //------------ TO USE IT ------------
    package main
    import (
      "fmt"  
      "temperature"
    )
    
    func main() {
      temp := temperature.CelsiusToFahrenheit(10.2)
      
      fmt.Println(temp)
    }
    
    ```
    
    Modules simplifies dependecies
    
    go.mod - lists every package program needs to work
    
    go mod init module-name
    
    // generates it on ./go.mod
    
    module module-name
    
    go 1.20 // go version
    
    go mod tidy // check program and insert dependency requirements on go.mod -> generates a go.sum (it checks integrity)
    
   
  - Testing
    Go has a native tool focused on tests (mainly unity tests);
    There are two main tests - basic (one input produces one output) and tabled (set of test cases submitted to the same function \[similar to theory]\);
    Testing should be apart from application;
    
    Go test files should:
    1. belong to the same directory and package as the code they test
    2. be named ending with _test.go
    3. import package testing
    
    test functions names should begin with Test:
    
    ```go
    // *testing.T is a generic testing pointer, it provides many possibilities
    func TestCelsiusToGahrenheit(t *testing.T) {
      cTemp := 29.0
      fTemp := 84.2
      
      // Fatal kills execution
      // Error prints error and keeps going with test cases
      if(CelsiusToFahrenheit(cTemp) != fTemp) {
        t.Fatal("\nError on converting from Celsius to Fahrenheit")
      } else {
        t.Log("\nConverting from Celsius to Fahrenheit: correct")
      }
    }
    
    // In order to test many cases, we create a slice to store param and expected value and loop over it
    ```
    
    Running:
    ```
    - run -> recieves a function to test 
             (if we dont specify this flag, all test functions will be runned)
             run="none" // disable test functions in test files (useful when working with benchmark)
    -bench -> recieves a function to test 
             (if we dont specify this flag, all test functions will be runned)
             bench="." // runs all benchmarks
             
    - benchmem -> show information about memmory allocation
    - v -> verbose test mode
        
    go run test -run=TestCelsiusToFahrenheit
    go run test -run=TestCelsiusToFahrenheit -v
    ```
    
    Go allow us to run tests in parallel (by default it runs sequentially):
    
    ```go
    func TestCelsiusToFahrenheit(t *testing.T) {
      t.Parallel() // <====================== it will make this test run parallel to next function
      // Test code
    }    
    ```
    
  - Benchmarking
    Go benchmark files should:
    1. belong to the same directory and package as the code they test
    2. be named ending with _test.go
    3. import package testing
    4. recieve b *testing.B pointer (actually this is not mandatory)

    OBS.: 
      - b.N < as much as possible in 1 second (comes from pointer)
      - b.ResetTimer() should be called at the beginning to reset in each function (generate a fair result)
   
  - HTTP Under GO
    https://pkg.go.dev/net/http

    We can create servers and clients

    ```go
    // SERVER
    
    package main
    import (
      "fmt"
      "log"
      "net/http"
    )
    func sayHello(res http.ResponseWriter, req *http.Request) {
      res.Write([]byte("Hello, stranger"))
      // fmt.Fprint(res, "Hello, stranger") -> mesmo comportamento
    }

    func printNow(res http.ResponseWriter, req *http.Request) {
      currentTime := time.Now()
      fmt.Fprint(res, currentTime.Format("02/01/2006 15:04:05"))
    }
    
    func main() {

      // We can custom http server through http object
      server := &http.Server {
        Addr: "localhost:4000",
        WriteTimeout: 10 * time.Second,
      }
      
      http.HandleFunc("/", sayHello)
      http.HandleFunc("/now", printNow)
      
      //log.Fatal(http.ListenAndServe("localhost:4000", nil)) // Listen to default http server
      log.Fatal(server.ListenAndServe("localhost:4000", nil)) // Listen to my custom http server
    }
    
    ```


    ```go
    // CLIENT

    package main
    import (
      "fmt"
      "io"
      "log"
      "net/http"
    )
    
    const url = "https://www.boredapi.com/api/activity"

    // Using package method
    func main() {
      // GET, HEAD, POST, PUT, DELETE

      if res, err := http.Get(url); err != nil {
        log.Fatal(err)
      } else {
        contents, _ := io.ReadAll(res.Body)
        defer res.Body.Close()
        fmt.Println(string(contents))
      }
    }

    // Creating request manually
    func main() {
      req, _ := http.NewRequest("GET", url, nil)
      if res, err := http.DefaultClient.Do(req); err != nil {
        log.Fatal(err)
      } else {
        contents, _ := io.ReadAll(res.Body)
        defer res.Body.Close()
        fmt.Println(string(contents))
      }
    }

    // We can create custom clients as well overloading config
    func main(){
      client := &http.Client{Timeout: 3 * time.Second} // OVERLOAD CLIENT CONFIG
      if res, err := client.Get(url); err != nil {
        log.Fatal(err)
      } else {
        contents, _ := io.ReadAll(res.Body)
        defer res.Body.Close()
        fmt.Println(string(contents))
      }
    }

    // Sending a POST request
    const endpoint = "https://httpbin.org/post"
    func main() {
      data := strings.NewReader(`{ "activity": "Learn to play a new instrument", "type": "music" }`)
      if res, err := http.Post(endpoint, "application/json", data); err != nil {
        log.Fatal(err)
      } else {
        contents, _ := io.ReadAll(res.Body)
        defer res.Body.Close()
        fmt.Printf("%s", contents)
      }
    }

    // Dealing with http errors
    res, err := http.Get(url)
    switch {
    case 400 <= res.StatusCode && res.StatusCode < 500:
      fmt.Println("Client error")
    case 500 <= res.StatusCode && res.StatusCode < 600:
      fmt.Println("Server error")
    }

    // Mapping json response to Structured Data
    const url = "https://www.boredapi.com/api/activity"
    type Activity struct {
      Activity string
      Type string
      Participants int
    }

    contents, _ := io.ReadAll(res.Body)
    defer res.Body.Close()
    var a Activity
    err := json.Unmarshal(contents, &a) // Maps to data structure
    if err != nil {
      log.Fatal(err)
    } else {
      fmt.Printf("What? %s\n", a.Activity)
      fmt.Printf("What kind? %s\n", a.Type)
      fmt.Printf("How many? %d\n", a.Participants)
    }
    
    ```







    
