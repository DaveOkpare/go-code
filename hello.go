// hello.go
package main // <-- Make sure it doesn't say 'main'.

import "fmt"

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello("David"))
}
