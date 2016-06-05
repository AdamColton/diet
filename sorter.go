package main

import "sort"

type sortByFn func(*food, *food) bool

type foodSorter struct {
	foods []*food
	by    sortByFn
}

func Sort(foods []*food) *foodSorter {
	return &foodSorter{foods: foods}
}

func (s *foodSorter) By(sortBy sortByFn) {
	s.by = sortBy
	sort.Sort(s)
}

func (s *foodSorter) Len() int           { return len(s.foods) }
func (s *foodSorter) Swap(i, j int)      { s.foods[i], s.foods[j] = s.foods[j], s.foods[i] }
func (s *foodSorter) Less(i, j int) bool { return s.by(s.foods[i], s.foods[j]) }

func Name(food1, food2 *food) bool     { return food1.name < food2.name }
func Dollars(food1, food2 *food) bool  { return food1.dollars < food2.dollars }
func G(food1, food2 *food) bool        { return food1.g < food2.g }
func NetCarbs(food1, food2 *food) bool { return food1.NetCarbs() < food2.NetCarbs() }
func Protein(food1, food2 *food) bool  { return food1.protein < food2.protein }
func Fat(food1, food2 *food) bool      { return food1.fat < food2.fat }
func Fiber(food1, food2 *food) bool    { return food1.fiber < food2.fiber }
func Calories(food1, food2 *food) bool { return food1.calories < food2.calories }

func Decending(sortBy sortByFn) sortByFn {
	return func(food1, food2 *food) bool {
		return !sortBy(food1, food2)
	}
}
