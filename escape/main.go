package main

type A struct {
	title string
	desc  string
}

func main() {
	dontEscape()
	escapes()
	escapesTwo()
}

func dontEscape() A {
	return A{"test1", "test"}
}

func escapes() *A {
	return &A{"test2", "test"}
}

func escapesTwo() {
	nonsized := make([]byte, 100000)
	nonsized = append(nonsized, 10)
}