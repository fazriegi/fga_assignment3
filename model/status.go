package model

type Status struct {
	Water int
	Wind  int
}

type Weather struct {
	Status Status
}
