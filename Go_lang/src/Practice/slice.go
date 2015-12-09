package main
import "fmt"
func main () {
    var a []int
    ar := [10]int{1,2,3,4}
    a = ar[0:3]
    fmt.Println(a)
    var test = []int{1,2,3,4,5}
    fmt.Println(test)
}
