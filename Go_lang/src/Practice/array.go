package main
import "fmt"
func f(a [3]int) { fmt.Println(a)}
func fp(a *[3]int) { fmt.Println(a)}
func main() {
    var arr [3]int
    fmt.Println(arr)
    fmt.Println(len(arr))
    f(arr)
    fp(&arr)
    one := [3]int{ 1,2,3 }
    fmt.Println(one)
    two := [10]int{1,2,3}
    fmt.Println(two)
    three := [...]int{1,2,3}
    fmt.Println(three)
    four := [10]int{2:1, 3:1, 5:1, 7:1}
    fmt.Println(four)
}
//godoc -http=:1989 -index=true
