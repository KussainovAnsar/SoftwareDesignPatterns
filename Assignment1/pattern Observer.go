package main

import (
	"fmt"
	"time"
)

type Observer interface {
	Update(score string)
}

type SportsSubject struct {
	observers []Observer
	score     string
}

func (s *SportsSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *SportsSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.score)
	}
}

func (s *SportsSubject) UpdateScore(score string) {
	s.score = score
	s.NotifyObservers()
}

type SportsFan struct {
	name string
}

func NewSportsFan(name string) *SportsFan {
	return &SportsFan{name}
}

func (f *SportsFan) Update(score string) {
	fmt.Printf("%s received score update: %s\n", f.name, score)
}

func main() {
	footballMatch := &SportsSubject{}
	fan1 := NewSportsFan("Fan 1")
	fan2 := NewSportsFan("Fan 2")
	fan3 := NewSportsFan("Fan 3")
	footballMatch.Attach(fan1)
	footballMatch.Attach(fan2)
	footballMatch.Attach(fan3)
	go func() {
		for i := 1; i <= 5; i++ {
			score := fmt.Sprintf("Home Team %d - %d Away Team", i, i)
			footballMatch.UpdateScore(score)
			time.Sleep(2 * time.Second)
		}
	}()
	select {}
}
