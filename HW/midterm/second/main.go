package main
import "fmt"
func main() {
    myArray := [5]int{1, 2, 3, 4, 5}
    for _, value := range myArray {
        fmt.Println(value)
    }
}