package main

import "fmt"

const (
	_ = iota
	// kb = 1024
	kb = 1 << (iota * 10) 
	mb = 1 << (iota * 10) 
	gb = 1 << (iota * 10) 
)

func main() {
	x := 2
	fmt.Printf("%d\t\t%b", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
	z := y << 1
	fmt.Printf("%d\t\t%b", z, z)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}