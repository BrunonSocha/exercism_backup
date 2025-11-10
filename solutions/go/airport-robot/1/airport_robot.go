package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
}

type Italian struct {
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(visitor string) string {
	return "Ciao " + visitor + "!"
}

type Portuguese struct {
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(visitor string) string {
	return "Olá " + visitor + "!"
}
func SayHello(name string, g Greeter) string {
	return "I can speak " + g.LanguageName() + ": " + g.Greet(name)
}


