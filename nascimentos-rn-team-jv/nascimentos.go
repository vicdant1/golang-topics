package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Year struct {
	Year  int
	Borns int
}

type City struct {
	Code  string
	Name  string
	Years []Year
	Total int
}

type Statistics struct {
	Cities      []City
	TotalBorns  int
	BornsByYear []Year
}

func main() {
	data := readFile()
	statistics := interpretData(data)
	validateFileIntegrity(statistics)

	mostBornNumberByYear := getMostBornNumberByYear(statistics)
	lessBornNumberByYear := getMostLessNumberByYear(statistics)
	bornRateByYear := getBornRateByYear(statistics)
	deviation := getStandartDeviationByYear(statistics)

	createAndSeedCsvFile(statistics, mostBornNumberByYear, lessBornNumberByYear, bornRateByYear, deviation)
	createAndSeedDatFile(statistics)
}

func createAndSeedDatFile(statistics Statistics) {

	file, err := os.Create("totais.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data := ""

	for _, v := range statistics.BornsByYear {
		data += fmt.Sprintf("%d %d\n", v.Year, v.Borns)
	}

	_, err = file.Write([]byte(data))

	if err != nil {
		fmt.Println("Falha ao criar arquivo .dat")
		os.Exit(1)
	}

}

func createAndSeedCsvFile(statistics Statistics, mostBornNumberByYear map[int]int, lessBornNumberByYear map[int]int, bornRateByYear map[int]float64, deviation map[int]float64) {
	csvFile, err := os.Create("estatisticas.csv")
	defer csvFile.Close()
	if err != nil {
		fmt.Println("Falha ao criar arquivo .CSV")
		os.Exit(1)
	}

	csvWritter := csv.NewWriter(csvFile)
	defer csvWritter.Flush()

	csvWritter.Write([]string{"ano", "maior_numero_nascidos", "menor_numero_nascidos", "media_de_nascimentos", "desvio_padrao", "total_de_nascimentos"})
	for _, year := range statistics.BornsByYear {
		csvWritter.Write([]string{strconv.Itoa(year.Year), strconv.Itoa(mostBornNumberByYear[year.Year]), strconv.Itoa(lessBornNumberByYear[year.Year]), fmt.Sprintf("%.f", bornRateByYear[year.Year]), fmt.Sprintf("%.f", deviation[year.Year]), strconv.Itoa(year.Borns)})
	}

}

// A maneira como estruturei os dados não me ajudou nesse método.
func getStandartDeviationByYear(statistics Statistics) map[int]float64 {
	standartDeviationByYear := make(map[int]float64)
	totalCities := len(statistics.Cities)
	bornRate := getBornRateByYear(statistics)

	years := make([]int, 0, len(statistics.BornsByYear))
	for _, key := range statistics.BornsByYear {
		years = append(years, key.Year)
	}

	var accumulate float64
	for _, year := range years {
		for _, city := range statistics.Cities {
			for _, cityYear := range city.Years {
				if year == cityYear.Year {
					accumulate += math.Pow(float64(cityYear.Borns)-bornRate[year], 2)
					break
				}
			}
		}
		accumulate = math.Sqrt(accumulate / float64(totalCities))
		standartDeviationByYear[year] = accumulate

		accumulate = 0
	}

	return standartDeviationByYear
}

func getBornRateByYear(statistics Statistics) map[int]float64 {
	bornRateByYear := make(map[int]float64)
	totalCities := len(statistics.Cities)

	for _, year := range statistics.BornsByYear {
		bornRateByYear[year.Year] = float64(year.Borns) / float64(totalCities)
	}

	return bornRateByYear
}

func getMostLessNumberByYear(statistics Statistics) map[int]int {
	lessBornNumberByYear := make(map[int]int)

	for _, city := range statistics.Cities {
		for _, year := range city.Years {
			if lessBornNumberByYear[year.Year] >= year.Borns {
				lessBornNumberByYear[year.Year] = year.Borns
			}
		}
	}

	return lessBornNumberByYear
}

func getMostBornNumberByYear(statistics Statistics) map[int]int {
	mostBornNumberByYear := make(map[int]int)

	for _, city := range statistics.Cities {
		for _, year := range city.Years {
			if mostBornNumberByYear[year.Year] < year.Borns {
				mostBornNumberByYear[year.Year] = year.Borns
			}
		}
	}

	return mostBornNumberByYear
}

func readFile() [][]string {
	filename := os.Args[1]
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	return data
}

func interpretData(data [][]string) Statistics {

	var cities []City
	totalBorns := 0
	var bornsByYear []Year

	totalYears := 0
	var headerData []string
	for _, row := range data {

		// skipando header
		if totalYears == 0 {
			headerData = row
			totalYears = len(row) - 2
			continue
		}

		isTotalRow := strings.ToLower(row[0]) == "total"
		var cityCode string
		var cityName string
		if !isTotalRow {
			cityCode = strings.Split(row[0], " ")[0]
			cityName = strings.ReplaceAll(row[0], cityCode, "")
		}

		var years []Year
		for i := 1; i <= totalYears; i++ {
			year, _ := strconv.Atoi(headerData[i])
			borns, _ := strconv.Atoi(row[i])

			years = append(years, Year{Year: year, Borns: borns})
		}

		totalYearBorns, _ := strconv.Atoi(row[totalYears+1])

		if isTotalRow {
			totalBorns = totalYearBorns
			bornsByYear = years
			continue
		}

		cityRegister := City{
			Code:  cityCode,
			Name:  cityName,
			Years: years,
			Total: totalYearBorns,
		}

		cities = append(cities, cityRegister)
	}

	statistics := Statistics{
		Cities:      cities,
		TotalBorns:  totalBorns,
		BornsByYear: bornsByYear,
	}

	return statistics
}

func validateFileIntegrity(statistics Statistics) {
	var totalBorns int

	contrastData := make(map[int]int)
	for _, city := range statistics.Cities {
		totalBorns += city.Total

		for _, year := range city.Years {
			contrastData[year.Year] += year.Borns
		}
	}

	for _, year := range statistics.BornsByYear {
		if contrastData[year.Year] != year.Borns {
			fmt.Printf("Inconsistência na leitura: %d nascidos em %d de acordo com o arquivo e %d de acordo com a leitura.\n", year.Borns, year.Year, contrastData[year.Year])
			os.Exit(1)
		}
	}

	if totalBorns != statistics.TotalBorns {
		fmt.Printf("Inconsistência na leitura: %d nascidos de acordo com o arquivo e %d de acordo com a leitura.\n", statistics.TotalBorns, totalBorns)
		os.Exit(1)
	}
}
