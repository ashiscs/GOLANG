package main

import "fmt"

func cal(x int) int {
	return x
}

func main() {
	fmt.Println(cal(501))
	fmt.Println(fac("Ashis", "Singh"))
}
func fac(f, l string) (string, string) {
	return f, l
}
