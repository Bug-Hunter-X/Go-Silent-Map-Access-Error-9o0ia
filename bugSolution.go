```go
func main() {
    var m map[string]int
    val, ok := m["a"]
    if ok {
        fmt.Println("Value found:", val)
    } else {
        fmt.Println("Key not found") // Handle the case when the key is missing
    }

    //Alternative using the comma ok idiom
    if val, ok := m["a"]; ok {
        fmt.Println("Value found:", val)
    } else {
        fmt.Println("Key not found")
    }

    //Initializing the map
    m = make(map[string]int)
    m["a"] = 1
    fmt.Println(m["a"]) //Prints 1
}
```