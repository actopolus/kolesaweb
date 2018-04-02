package main

func main() {
    calc()
}

// go:noinline
func calc() int {
	b := 2
	c := 3
	b  = c

	return b + c
}
