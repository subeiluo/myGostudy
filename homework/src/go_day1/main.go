package main

import "fmt"

func main() {
	var (
		a = "为了心中不变的梦!"
		b = 20
		c = 40
		d = 60
		e = 3.1415926
	)
	fmt.Printf("this is string= %s\n", a)
	fmt.Printf("this is binary= %b\n", b)
	fmt.Printf("this is Decimal= %d\n", c)
	fmt.Printf("this is Hexadecimal= %x\n", d)
	fmt.Printf("this is Hexadecimal= %X\n", d)
	fmt.Printf("this is float= %f\n", e)
}
