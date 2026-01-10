package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct{}

func (greeter Italian) LanguageName() string {
	return "Italian"
}

func (greeter Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

func (greeter Portuguese) LanguageName() string {
	return "Portuguese"
}

func (greeter Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
