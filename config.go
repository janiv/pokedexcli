package main

type config struct {
	locationOffset int
	locationLimit  int
}

func NewConfig() config {
	//These are just the defaults when you hit the api
	//Need to add stuff to increment, decrement
	return config{
		locationOffset: 0,
		locationLimit:  20,
	}
}
