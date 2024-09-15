package models

// Dentista representa a un odontólogo en el sistema.
type Dentista struct {
	ID        int    `json:"id"`
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
	Matricula string `json:"matricula"`
}
