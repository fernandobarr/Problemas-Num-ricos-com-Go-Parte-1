package main

import "fmt"

func prob_Num() {
	for i := 1; i <= 100; i++ {
		var divisivel int
		var pim int = 3
		divisivel = i % pim
		if divisivel == 0 {
			fmt.Println(i)
			//fmt.Printf("Pim - %v e %v = %v ", i, pim, divisivel)
			//fmt.Println("")
		}
		//else {
		//	fmt.Println(i)
		//}
	}
}

func main() {
	prob_Num()
}
