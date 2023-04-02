package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type State struct {
	Name            string
	Acronymn        string
	Region          string
	CitiesCount     int
	PopulationCount int
}

func main() {

	states := make(map[string]State)

	f, err := os.Open("../static/estados.csv")
	defer f.Close()

	if err != nil {
		fmt.Println("Erro: " + err.Error())
	}

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Erro: " + err.Error())
	}

	interpretData(data, &states)

	for state, data := range states {
		fmt.Printf("%s - %s (%d Habitantes)\n", state, data.Acronymn, data.PopulationCount)
	}
}

func interpretData(data [][]string, states *map[string]State) {

	for index, row := range data {
		if index == 0 {
			continue
		}

		name := row[0]
		citiesCount, _ := strconv.Atoi(row[3])
		populationCount, _ := strconv.Atoi(row[4])

		state := State{
			Name:            name,
			Acronymn:        row[1],
			Region:          row[2],
			CitiesCount:     citiesCount,
			PopulationCount: populationCount,
		}

		(*states)[name] = state
	}
}
