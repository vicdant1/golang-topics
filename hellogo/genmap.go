package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Criando struct Entry utilizando generics
// Utilizando tipo comparable para satisfazer condição especificada no readme
type Entry[K int | string, V comparable] struct {
	Key   K
	Value V
}

// Criando struct Map passando K e V dos mesmos tipos usados em Entry
type Map[K int | string, V comparable] struct {
	Entries []Entry[K, V]
}

// Criando função addEntry - passando receiver e chave/valor
func (m *Map[K, V]) addEntry(key K, value V) {

	defer func() { // utilizando defer e anonymous function para printar o slice ao final
		for _, v := range m.Entries {
			fmt.Printf("[%v] %v\n", v.Key, v.Value)
		}
	}()

	newEntry := Entry[K, V]{ // criando nova entry para inserir
		Key:   key,
		Value: value,
	}

	for i, v := range m.Entries { // varrendo slice
		if v.Key == key { // validando se chave ja existe no slice para substituir
			m.Entries[i].Value = value
			return // utilizando return para interromper o fluxo de execução
		}
	}

	m.Entries = append(m.Entries, newEntry) // caso não exista a chave, inserir keyvaluepair
}

// Criando função getEntry - passando receiver e chave
func (m *Map[K, V]) getEntry(key K) {
	for _, v := range m.Entries { // varrendo slice
		if v.Key == key { // validando se a chave existe
			fmt.Printf("[%v] %v\n", v.Key, v.Value) // apresentando informação na saida padrao
			return                                  // utilizando return para interromper o fluxo de execução
		}
	}

	fmt.Println("Chave inexistente, tente novamente.") // printando mensagem de chave inexistente
}

// Criando função mapSize - passando receiver
func (m *Map[K, V]) mapSize() {
	fmt.Println(len(m.Entries)) // printando informação (quantidade de entries)
}

// Criando função printMap - passando receiver
func (m *Map[K, V]) printMap() {
	fmt.Println(m.Entries) // printando todas as entradas
	m.mapSize()            // chamando funcao de map para printar o tamanho
}

func main() {
	// criando mapa de acordo com o especificado ([int]:string)
	genericMap := Map[int, string]{}

	for {
		fmt.Print("> ")

		// fazendo leitura e atribuicao das variaveis
		var command string
		var key int
		var value string

		// lendo entrada do usuário
		reader := bufio.NewReader(os.Stdin)
		userInput, _, _ := reader.ReadLine()
		userInputString := string(userInput)
		userInputSize := len(strings.Fields(userInputString)) - 1 // determinando quantidade de parâmetros informados

		fmt.Sscanf(userInputString, "%s %d %s\n", &command, &key, &value) // passando informação para as variaveis

		// switch para chamar cada funcao de acordo com sua abreviacao
		switch strings.ToLower(command) { // tratando command para ser case insensitive
		case "add":
			if userInputSize > 2 {
				fmt.Println("Quantidade de parâmetros inválida, tente novamente.")
				break
			}

			genericMap.addEntry(key, value)
			break

		case "get":
			if userInputSize > 1 {
				fmt.Println("Quantidade de parâmetros inválida, tente novamente.")
				break
			}

			genericMap.getEntry(key)
			break

		case "size":
			genericMap.mapSize()
			break

		case "print":
			genericMap.printMap()
			break

		case "exit":
			os.Exit(0) // finalizando a execução do programa
			break

		default:
			// exibindo mensagem de comando desconhecido
			fmt.Println("Comando não reconhecido, tente novamente.")
			break

		}

		fmt.Println()

	}
}
