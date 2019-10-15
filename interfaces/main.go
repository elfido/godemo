package main

type Program struct {
	ID          string
	Name        string
	ReleaseYear int
}

type Storage interface {
	Add(entityID string, program Program) error
}

func main() {
	var storage Storage
	storage = &FileStorage{}
	program := Program{
		ID:          "xzy",
		Name:        "some program",
		ReleaseYear: 2019,
	}
	storage.Add(program.ID, program)
}
