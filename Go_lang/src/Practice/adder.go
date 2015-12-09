package main
import "fmt"
func adder() (func(int) int){
    var x int
    return func(delta int) int {
        x += delta
        return x
    }
}
func main () {
f := adder()
fmt.Println(f(1))
fmt.Println(f(20))
fmt.Println(f(300))
}
//godoc -http=:1989 -index=true
