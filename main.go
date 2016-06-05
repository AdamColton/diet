package main

import "fmt"

const (
	header = "INGREDIENT        NET CARBS   PROTEIN   CALORIES       $      G\n"
	breaks = "---------------------------------------------------------------\n"
)

func main() {
	fmt.Print(header, breaks)

	Sort(smoothie).By(Decending(Protein))
	for _, ingredient := range smoothie {
		fmt.Print(ingredient)
	}

	var smoothieSum = Sum("SMOOTHIE", smoothie...)
	fmt.Print(breaks, smoothieSum)
}
