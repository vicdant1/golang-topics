/*TODOS

- VALIDATE CODE
- FORMAT OUTPUT LIKE PROFESSOR DID
- TEST CONDITIONS
- TEST CASE INSENSITIVE - CHECK IF THIS IS CAUSING ERRORS

*/

package main

import (
	"fmt"
	"os"
	"strings"
)

type Entry[K int | string, V comparable] struct {
	Key   K
	Value V
}

type Map[K int | string, V comparable] struct {
	Entries []Entry[K, V]
}

func (m *Map[K, V]) addEntry(key K, value V) {
	defer func() {
		for _, v := range m.Entries {
			fmt.Printf("[%v] %v\n", v.Key, v.Value)
		}
	}()

	newEntry := Entry[K, V]{
		Key:   key,
		Value: value,
	}

	for i, v := range m.Entries {
		if v.Key == key {
			m.Entries[i].Value = value
			return
		}
	}

	m.Entries = append(m.Entries, newEntry)
}

func (m *Map[K, V]) getEntry(key K) {
	for _, v := range m.Entries {
		if v.Key == key {
			fmt.Printf("[%v] %v\n", v.Key, v.Value)
		}
	}
}

func (m *Map[K, V]) mapSize() {
	fmt.Println(len(m.Entries))
}

func (m *Map[K, V]) printMap() {
	fmt.Println(m.Entries)
	m.mapSize()
}

func main() {
	genericMap := Map[int, string]{}

	for {

		var command string
		var key int
		var value string
		fmt.Scanf("%s %d %s", &command, &key, &value)

		switch strings.ToLower(command) {
		case "add":
			genericMap.addEntry(key, value)
			break

		case "get":
			genericMap.getEntry(key)
			break

		case "size":
			genericMap.mapSize()
			break

		case "print":
			genericMap.printMap()
			break

		default:
			fmt.Println("Comando não reconhecido, tente novamente.")
			break
		case "exit":
			os.Exit(0)
		}

	}
}
