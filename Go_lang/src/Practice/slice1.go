package main
import "fmt"
//var sl = make([]int, 0, 100)
func appendToSlice(i int, sl []int) []int {
    if len(sl) == cap(sl){ 
        fmt.Println("test") 
    }
    n := len(sl)
    sl = sl[0:n+1]
    sl[n] = i
    return sl
}
func main () {
    var sl = make([]int, 0, 100)
    fmt.Println(sl)
    fmt.Println(appendToSlice(1, sl))
}

