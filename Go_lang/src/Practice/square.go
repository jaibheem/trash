package main
import (
    "fmt"
    "os"
    "strconv"
    "reflect"
)

func square (x int) int{
    square := x * x
    return square
}
func main () {
    fmt.Println("Enter a Number")
    args := os.Args
    one := strconv.Atoi(args[1])
    fmt.Println(reflect.TypeOf(one))
    //fmt.Println("Square of %d is %d", args, square(args))
    //fmt.Println("%d%d", args, square(args))
}
