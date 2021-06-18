package main

import "fmt"

func main() {
	var foD []string
	foS := make([]string, 3)
	foD = append(foD, "Костя", "Семён", "Таня")
	foS = append(foS, "Валера", "Таня", "Дима")
	friends := map[string][]string{
		"Dima":   foD,
		"Semyon": foS,
		"Костя":  nil,
	}
	_, value := friends["Костя"]
	fmt.Print(value, " ")
}
