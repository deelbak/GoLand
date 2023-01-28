package Hello

import "fmt"

var s string

func Say(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
func main() {
	fmt.Scanln(&s)
	fmt.Println(Say(s))
}
