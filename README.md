# ARGolang

Proyecto base de API en Go con PostgreSQL.

## Estructura de Carpetas

- `cmd/api`: Punto de entrada de la aplicación.
- `internal/`: Lógica interna de la aplicación.
    - `database/`: Conexión y migraciones.
    - `handlers/`: Controladores HTTP.
    - `models/`: Definición de datos.
    - `repository/`: Capa de persistencia.
- `scripts/`: Scripts de base de datos.
- `tests/`: Pruebas de integración.

## Requisitos

- Go 1.21+
- PostgreSQL

## Inicio Rápido

1. Copia el archivo `.env.example` a `.env` y configura tus variables.
2. Ejecuta `make run` para iniciar el servidor.
3. Ejecuta `make test` para correr las pruebas.
