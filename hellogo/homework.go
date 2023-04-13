package main

import "fmt"

type Entry[K int | string, V comparable] struct {
	Key   K
	Value V
}

type Map[K int | string, V comparable] struct {
	Entries []Entry[K, V]
}

func main() {
	fmt.Println('d' > 'z')
}
