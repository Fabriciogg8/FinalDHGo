# Examen Final Backend III

**Autor:**  
Fabricio González

**Materia:**  
Este proyecto es parte de la especialización de Backend dentro de la carrera Certified Tech Developer, bajo la instrucción del profesor Pablo López.

**Objetivo:**  
El objetivo de este proyecto es desarrollar una API REST en Go para la gestión de turnos odontológicos, que incluye la implementación de endpoints para la creación, actualización, eliminación y consulta de turnos, pacientes y dentistas.

## Endpoints

### Turnos

- **Crear Turno**
  - **URL:** `/turnos`
  - **Método:** POST
  - **Descripción:** Crea un nuevo turno.
  - **Parámetros:**
    - **Cuerpo:** `Turno` (ID, FechaHora, Descripcion, DentistaID, PacienteID)
  - **Respuesta:** 201 Created

- **Obtener Todos los Turnos**
  - **URL:** `/turnos`
  - **Método:** GET
  - **Descripción:** Obtiene todos los turnos.
  - **Respuesta:** 200 OK (Array de `Turno`)

- **Obtener Turno por ID**
  - **URL:** `/turnos/{id}`
  - **Método:** GET
  - **Descripción:** Obtiene un turno por su ID.
  - **Parámetros:**
    - **ID:** ID del turno
  - **Respuesta:** 200 OK (`Turno`)

- **Actualizar Turno**
  - **URL:** `/turnos/{id}`
  - **Método:** PUT
  - **Descripción:** Actualiza un turno existente.
  - **Parámetros:**
    - **ID:** ID del turno
    - **Cuerpo:** `Turno` (FechaHora, Descripcion, DentistaID, PacienteID)
  - **Respuesta:** 200 OK

- **Actualizar Parcialmente Turno**
  - **URL:** `/turnos/{id}`
  - **Método:** PATCH
  - **Descripción:** Actualiza parcialmente un turno existente.
  - **Parámetros:**
    - **ID:** ID del turno
    - **Cuerpo:** Mapa de campos a actualizar
  - **Respuesta:** 200 OK

- **Eliminar Turno**
  - **URL:** `/turnos/{id}`
  - **Método:** DELETE
  - **Descripción:** Elimina un turno.
  - **Parámetros:**
    - **ID:** ID del turno
  - **Respuesta:** 204 No Content

- **Crear Turno por DNI y Matrícula**
  - **URL:** `/turnos/by-dni-matricula`
  - **Método:** POST
  - **Descripción:** Crea un turno usando el DNI del paciente y la matrícula del dentista.
  - **Parámetros:**
    - **DNI:** DNI del paciente
    - **Matrícula:** Matrícula del dentista
    - **Cuerpo:** `Turno` (FechaHora, Descripcion)
  - **Respuesta:** 201 Created

- **Obtener Turno por DNI**
  - **URL:** `/turnos/by-dni`
  - **Método:** GET
  - **Descripción:** Obtiene un turno por el DNI del paciente.
  - **Parámetros:**
    - **DNI:** DNI del paciente
  - **Respuesta:** 200 OK (`Turno`)

### Pacientes

- **Crear Paciente**
  - **URL:** `/pacientes`
  - **Método:** POST
  - **Descripción:** Crea un nuevo paciente.
  - **Parámetros:**
    - **Cuerpo:** `Paciente` (DNI, Nombre, etc.)
  - **Respuesta:** 201 Created

- **Obtener Todos los Pacientes**
  - **URL:** `/pacientes`
  - **Método:** GET
  - **Descripción:** Obtiene todos los pacientes.
  - **Respuesta:** 200 OK (Array de `Paciente`)

- **Obtener Paciente por ID**
  - **URL:** `/pacientes/{id}`
  - **Método:** GET
  - **Descripción:** Obtiene un paciente por su ID.
  - **Parámetros:**
    - **ID:** ID del paciente
  - **Respuesta:** 200 OK (`Paciente`)

- **Actualizar Paciente**
  - **URL:** `/pacientes/{id}`
  - **Método:** PUT
  - **Descripción:** Actualiza un paciente existente.
  - **Parámetros:**
    - **ID:** ID del paciente
    - **Cuerpo:** `Paciente` (DNI, Nombre, etc.)
  - **Respuesta:** 200 OK

- **Eliminar Paciente**
  - **URL:** `/pacientes/{id}`
  - **Método:** DELETE
  - **Descripción:** Elimina un paciente.
  - **Parámetros:**
    - **ID:** ID del paciente
  - **Respuesta:** 204 No Content

### Dentistas

- **Crear Dentista**
  - **URL:** `/dentistas`
  - **Método:** POST
  - **Descripción:** Crea un nuevo dentista.
  - **Parámetros:**
    - **Cuerpo:** `Dentista` (Matrícula, Nombre, etc.)
  - **Respuesta:** 201 Created

- **Obtener Todos los Dentistas**
  - **URL:** `/dentistas`
  - **Método:** GET
  - **Descripción:** Obtiene todos los dentistas.
  - **Respuesta:** 200 OK (Array de `Dentista`)

- **Obtener Dentista por ID**
  - **URL:** `/dentistas/{id}`
  - **Método:** GET
  - **Descripción:** Obtiene un dentista por su ID.
  - **Parámetros:**
    - **ID:** ID del dentista
  - **Respuesta:** 200 OK (`Dentista`)

- **Actualizar Dentista**
  - **URL:** `/dentistas/{id}`
  - **Método:** PUT
  - **Descripción:** Actualiza un dentista existente.
  - **Parámetros:**
    - **ID:** ID del dentista
    - **Cuerpo:** `Dentista` (Matrícula, Nombre, etc.)
  - **Respuesta:** 200 OK

- **Eliminar Dentista**
  - **URL:** `/dentistas/{id}`
  - **Método:** DELETE
  - **Descripción:** Elimina un dentista.
  - **Parámetros:**
    - **ID:** ID del dentista
  - **Respuesta:** 204 No Content

  ## Instrucciones de Instalación

1. **Clonar el repositorio:**

    ```sh
    git clone <URL_DEL_REPOSITORIO>
    ```

2. **Navegar al directorio del proyecto:**

    ```sh
    cd FinalDH
    ```

3. **Instalar las dependencias:**

    ```sh
    go mod tidy
    ```

4. **Ejecutar la aplicación:**

    ```sh
    go run main.go
    ```

**Contribuciones** 
Las contribuciones al proyecto son bienvenidas. Si encuentras algún error o tienes sugerencias de mejoras, por favor, abre un issue o envía una solicitud de extracción.