package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Давлеканово", "ЧИШМЫ"}, // города с населением 10-99 тыс. человек
		100:  []string{"Стерлитамак", "Пермь"}, // города с населением 100-999 тыс. человек
		1000: []string{"Москва", "Уфа"},        // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Стерлитамак": 824123,
		"Пермь":       13000000,
	}
	for cityName, _ := range cityPopulation {
		flag := false
		for _, elemFromGroup := range groupCity[100] {
			if cityName == elemFromGroup && !flag {
				flag = true
				break
			}
		}
		if !flag {
			delete(cityPopulation, cityName)
		}
	}
	fmt.Print(cityPopulation)

}
