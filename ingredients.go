package main

// Everything is in terms of "per 100 g" (except for g, which is the amount of
// that thing that you're eating).

// perPound takes the price per pound and returns the price per 100 grams.
func perPound(per100g float32) float32 {
	return per100g * 0.220462
}

////////////////////////////////////////////////////////////////////////////////

type Grams float32

const (
	kg = 1000
)

func (g Grams) Of(f *food) *food {
	r := float32(g / f.g)
	return &food{
		name:     f.name,
		g:        g,
		dollars:  r * f.dollars,
		carbs:    r * f.carbs,
		protein:  r * f.protein,
		fat:      r * f.fat,
		fiber:    r * f.fiber,
		calories: r * f.calories,
	}
}

var water = &food{
	name:     "water",
	g:        1,
	dollars:  0,
	carbs:    0,
	protein:  0,
	fat:      0,
	fiber:    0,
	calories: 0,
}

var lentils = &food{
	name:     "lentils",
	g:        0.1 * kg,
	dollars:  perPound(2.39) / 5,
	carbs:    20.13,
	protein:  9.02,
	fat:      0.38,
	fiber:    7.9,
	calories: 116,
}

var spinach = &food{
	name:     "spinach",
	g:        0.1 * kg,
	dollars:  perPound(4.29),
	carbs:    3.63,
	protein:  2.86,
	fat:      0.39,
	fiber:    2.2,
	calories: 23,
}

var strawberries = &food{
	name:     "strawberries",
	g:        0.1 * kg,
	dollars:  perPound(1.67),
	carbs:    9.13,
	protein:  0.43,
	fat:      0.11,
	fiber:    2.1,
	calories: 35,
}

var almonds = &food{
	name:     "almonds",
	g:        0.1 * kg,
	dollars:  perPound(5.63),
	carbs:    21.55,
	protein:  21.15,
	fat:      49.93,
	fiber:    12.5,
	calories: 579,
}

var hempSeeds = &food{
	name:     "hemp seeds",
	g:        0.1 * kg,
	dollars:  (13.69 / 793) * 100,
	carbs:    8.67,
	protein:  31.56,
	fat:      48.75,
	fiber:    4,
	calories: 553,
}

var applesGala = &food{
	name:     "apples, gala",
	g:        0.1 * kg,
	dollars:  perPound(2.50),
	carbs:    13.68,
	protein:  0.25,
	fat:      0.12,
	fiber:    2.3,
	calories: 57,
}
