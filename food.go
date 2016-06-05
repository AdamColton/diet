package main

import "fmt"

type food struct {
	name     string
	dollars  float32
	g        Grams
	carbs    float32
	protein  float32
	fat      float32
	fiber    float32
	calories float32
}

func (f *food) String() string {
	return fmt.Sprintf("%15s %11.0f %9.0f %10.0f %7.2f %6.0f\n",
		f.name,
		f.NetCarbs(),
		f.protein,
		f.calories,
		f.dollars,
		f.g,
	)
}

func Sum(name string, foods ...*food) *food {
	total := &food{name: name}
	for _, f := range foods {
		total.dollars += f.dollars
		total.g += f.g
		total.carbs += f.carbs
		total.protein += f.protein
		total.fat += f.fat
		total.fiber += f.fiber
		total.calories += f.calories
	}
	return total
}

func (f *food) NetCarbs() float32 {
	return f.carbs - f.fiber
}
