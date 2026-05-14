# 🧹 CRUD Golang (Arquitectura Hexagonal)

Este proyecto es un CRUD de usuarios desarrollado en Go, siguiendo los principios de arquitectura hexagonal, con buenas prácticas de validación, logging, documentación y manejo de dependencias.

---

## 🚀 Tecnologías Utilizadas

* **Go** 1.21+
* **Echo**: Framework web
* **PostgreSQL**: Base de datos relacional
* **Zerolog**: Logging estructurado
* **Swaggo**: Generador de documentación Swagger
* **Go-playground/validator**: Validación de structs
* **Uber/dig**: Inyección de dependencias

---

## 📦 Instalación

```bash
cd crud_golang
go mod tidy
```

---

## 🥪 Ejecutar Proyecto

Asegúrate de tener PostgreSQL ejecutando y un archivo `.env` con las credenciales. Luego:

```bash
go run main.go
```

Por defecto, el servicio escuchará en `http://localhost:8081`.

---

## 💃 Endpoints

| Método | Endpoint     | Descripción            |
| ------ | ------------ | ---------------------- |
| GET    | `/orders`     | Listar ordenes        |
| GET    | `/orders/:id` | Obtener orden por ID |
| POST   | `/orders`     | Crear nueva order    |
| PUT    | `/orders/:id` | Actualizar orden     |
| DELETE | `/orders/:id` | Eliminar orden       |

---

## 📓 Ejemplo de Orden

```json
{
  "name": "Juan Pérez",
  "email": "juan@example.com"
}
```

---

## 📘 Documentación Swagger

Después de compilar los docs con:

```bash
swag init
```

La documentación estará disponible en:

```
http://localhost:8081/swagger/index.html
```

---

## 🔍 Estructura del Proyecto

```
internal/
├── application/         # Casos de uso (servicios)
├── domain/              # Modelos y puertos
├── infrastructure/
│   ├── db/              # Acceso a datos con SQL
│   ├── http/            # Controladores y middlewares
│   ├── di/              # Inyección de dependencias
│   ├── kit/             # Utilidades y constantes
cmd/                     # Entry point
docs/                    # Archivos Swagger generados
```
