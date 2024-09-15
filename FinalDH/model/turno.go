package models

// Turno represents a dental appointment
type Turno struct {
    ID          int    `json:"id"`
    FechaHora   string `json:"fecha_hora"`
    Descripcion string `json:"descripcion"`
    DentistaID  int    `json:"dentista_id"`
    PacienteID  int    `json:"paciente_id"`
}