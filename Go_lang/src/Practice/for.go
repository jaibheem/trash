package main
import "fmt"
func main() {
    var v int
    for i := 0; i < 5; i++{
        //var v int
        fmt.Printf("%d\n", v)
        v = 5
	defer fmt.Printf("I am done with all Iterations%d\n", i)
    }
}
//godoc -http=:1989 -index=true
