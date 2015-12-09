package main
import "fmt"
func main() {
    m := make(map[string]float64)
    fmt.Println(m)
    m1 := map[string]float64{"1":1, "pi":3.141}
    fmt.Println(m1)
    var one float64
    one = m1["1"]
    fmt.Println(one)
    pi := m1["pi"]
    fmt.Println(pi)
    m1["two"] = 2
    m1["two"] = 3
    fmt.Println(m1["two"])
    var value float64
    var present bool
    value, present = m["test"]
    fmt.Println(value, present)
    m2 := make(map[string]int)
    m2["ten"] = 10   
    m2["Eleven"] = 11
    m2["Twelve"] = 12
    fmt.Println(m2)
    for key, value := range m2 {
        fmt.Printf("Keys %s, Value %d\n", key, value)
    }
    yes, no := m2["ten"]
    fmt.Println(yes, no)
}
