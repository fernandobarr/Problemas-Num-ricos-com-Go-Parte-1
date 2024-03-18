package main

import "fmt"

func prob_Num() {
	for i := 1; i <= 100; i++ {
		var divisivel int
		var pin int = 3
		divisivel = i % pin
		if divisivel == 0 {
			fmt.Println(i)
			//fmt.Printf("Pin - %v e %v = %v ", i, pin, divisivel)
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
