package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{"Pizza": 3990, "Coke": 500, "Chicken Nuggets": 900, "French fries": 1100, "PiskaSiska": 999}

	tempmap := map[string]float64{"Pizza": 0, "Coke": 0, "Chicken Nuggets": 0, "French fries": 0, "PiskaSiska": 0}

	var choice string
	fmt.Println("Добро пожаловать в меню")
	for {

		fmt.Println("Доступные блюда:\n")

		for i, v := range menu {
			fmt.Printf("- %s: %f тг\n", i, v)

		}
		fmt.Println("\nВведите название блюда или 'exit' для выхода")
		fmt.Scan(&choice)
		if choice == "exit" {
			break
		}
		_, ok := menu[choice]
		if ok {
			tempmap[choice] += menu[choice]
		} else {
			fmt.Println("К сожалению, этого блюда нет в меню")
		}
	}
	summ := 0.0
	for product, value := range tempmap {
		if value != 0 {
			fmt.Printf("%s x%d - %f", product, int(value/menu[product]), value)
			summ += value
		}

	}
	fmt.Println()
	fmt.Println(summ)
	discount := 0.0
	if summ >= 5000 {

		discount := summ * 0.10
		fmt.Println(discount)
		summ = summ - discount
		fmt.Println("Итого", summ)
	}

	tax := summ * (12.0 / 100)
	total := (summ + tax) - discount
	fmt.Println("Итого с ндс", total)

}
