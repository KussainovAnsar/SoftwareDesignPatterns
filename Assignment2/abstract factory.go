package main

import "fmt"

type AbstractFactory interface {
	CreateStarter() Starter
	CreateMainCourse() MainCourse
}

type Starter interface {
	Description() string
}

type MainCourse interface {
	Description() string
}

type BreakfastFactory struct{}

func (bf BreakfastFactory) CreateStarter() Starter {
	return &BreakfastStarter{}
}

func (bf BreakfastFactory) CreateMainCourse() MainCourse {
	return &BreakfastMainCourse{}
}

type LunchFactory struct{}

func (lf LunchFactory) CreateStarter() Starter {
	return &LunchStarter{}
}

func (lf LunchFactory) CreateMainCourse() MainCourse {
	return &LunchMainCourse{}
}

type DinnerFactory struct{}

func (df DinnerFactory) CreateStarter() Starter {
	return &DinnerStarter{}
}

func (df DinnerFactory) CreateMainCourse() MainCourse {
	return &DinnerMainCourse{}
}

type BreakfastStarter struct{}

func (bs BreakfastStarter) Description() string {
	return "Breakfast Starter: Cereal"
}

type BreakfastMainCourse struct{}

func (bm BreakfastMainCourse) Description() string {
	return "Breakfast Main Course: Omelette"
}

type LunchStarter struct{}

func (ls LunchStarter) Description() string {
	return "Lunch Starter: Soup"
}

type LunchMainCourse struct{}

func (lm LunchMainCourse) Description() string {
	return "Lunch Main Course: Sandwich"
}

type DinnerStarter struct{}

func (ds DinnerStarter) Description() string {
	return "Dinner Starter: Salad"
}

type DinnerMainCourse struct{}

func (dm DinnerMainCourse) Description() string {
	return "Dinner Main Course: Steak"
}

func main() {
	breakfastFactory := BreakfastFactory{}
	breakfastStarter := breakfastFactory.CreateStarter()
	breakfastMainCourse := breakfastFactory.CreateMainCourse()

	fmt.Println(breakfastStarter.Description())
	fmt.Println(breakfastMainCourse.Description())

	lunchFactory := LunchFactory{}
	lunchStarter := lunchFactory.CreateStarter()
	lunchMainCourse := lunchFactory.CreateMainCourse()

	fmt.Println(lunchStarter.Description())
	fmt.Println(lunchMainCourse.Description())

	dinnerFactory := DinnerFactory{}
	dinnerStarter := dinnerFactory.CreateStarter()
	dinnerMainCourse := dinnerFactory.CreateMainCourse()

	fmt.Println(dinnerStarter.Description())
	fmt.Println(dinnerMainCourse.Description())
}
