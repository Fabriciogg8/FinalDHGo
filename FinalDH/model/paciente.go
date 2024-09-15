package models

// Paciente representa a un paciente en el sistema.
type Paciente struct {
	ID        int    `json:"id"`
	Apellido  string `json:"apellido"`
	Nombre    string `json:"nombre"`
	DNI       string `json:"dni"`
	Domicilio string `json:"domicilio"`
	FechaAlta string `json:"fecha_alta"`
}
