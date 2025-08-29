# Koopi Events API - Documentación

## Descripción General

Koopi Events API es una aplicación backend desarrollada en Go para la gestión de eventos. Permite crear, buscar y gestionar eventos de diferentes tipos con integración geográfica.

## Tecnologías Utilizadas

- **Backend**: Go (Golang) 1.21+
- **Framework**: Gin
- **Base de Datos**: PostgreSQL
- **ORM**: GORM
- **Autenticación**: JWT
- **Documentación**: Swagger/OpenAPI
- **Migraciones**: golang-migrate

## Estructura del Proyecto

```
koopi-backend/
├── cmd/
│   └── server/          # Punto de entrada de la aplicación
├── internal/
│   ├── config/          # Configuración de la aplicación
│   ├── database/        # Configuración de base de datos
│   ├── handlers/        # Manejadores HTTP
│   └── middleware/      # Middlewares
├── pkg/
│   ├── models/          # Modelos de datos
│   ├── dto/             # Data Transfer Objects
│   └── utils/           # Utilidades
├── migrations/          # Migraciones SQL
├── docker/              # Configuración Docker
└── docs/                # Documentación
```

## Instalación y Configuración

### Prerrequisitos

- Go 1.21+
- Docker y Docker Compose
- PostgreSQL

### Pasos de Instalación

1. **Clonar el repositorio**
```bash
git clone <repository-url>
cd koopi-backend
```

2. **Configurar variables de entorno**
```bash
cp env.example .env
# Editar .env con tus configuraciones
```

3. **Instalar dependencias y herramientas**
```bash
make setup
```

4. **Levantar servicios Docker**
```bash
make docker-up
```

5. **Ejecutar migraciones**
```bash
make migrate-up
```

6. **Ejecutar la aplicación**
```bash
make run
```

## Endpoints Principales

### Autenticación
- `POST /api/v1/auth/register` - Registro de usuario
- `POST /api/v1/auth/login` - Inicio de sesión
- `GET /api/v1/auth/profile` - Obtener perfil (protegido)

### Eventos
- `GET /api/v1/events` - Listar eventos
- `GET /api/v1/events/:id` - Obtener evento específico
- `POST /api/v1/events` - Crear evento (protegido)
- `PUT /api/v1/events/:id` - Actualizar evento (protegido)
- `DELETE /api/v1/events/:id` - Eliminar evento (protegido)
- `GET /api/v1/events/nearby` - Eventos cercanos
- `GET /api/v1/events/search` - Búsqueda de eventos

### Organizaciones
- `GET /api/v1/organizations` - Listar organizaciones
- `GET /api/v1/organizations/:id` - Obtener organización
- `POST /api/v1/organizations` - Crear organización (protegido)
- `PUT /api/v1/organizations/:id` - Actualizar organización (protegido)

### Reservas
- `GET /api/v1/bookings` - Mis reservas (protegido)
- `POST /api/v1/bookings` - Crear reserva (protegido)
- `PUT /api/v1/bookings/:id/cancel` - Cancelar reserva (protegido)

## Modelos de Datos

### User
```json
{
  "id": 1,
  "email": "user@example.com",
  "name": "Usuario Ejemplo",
  "role": "user",
  "is_active": true,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### Event
```json
{
  "id": 1,
  "title": "Concierto de Rock",
  "description": "Gran concierto de rock en vivo",
  "category_id": 1,
  "organizer_id": 1,
  "location_id": 1,
  "start_date": "2024-02-01T20:00:00Z",
  "end_date": "2024-02-01T23:00:00Z",
  "price": 25.00,
  "capacity": 500,
  "is_active": true,
  "is_free": false,
  "image_url": "https://example.com/image.jpg",
  "website": "https://example.com",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## Autenticación

La API utiliza JWT (JSON Web Tokens) para la autenticación. Para acceder a endpoints protegidos, incluye el token en el header Authorization:

```
Authorization: Bearer <your-jwt-token>
```

## Códigos de Respuesta

- `200` - OK
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `500` - Internal Server Error

## Desarrollo

### Comandos Útiles

```bash
# Ejecutar en desarrollo
make run

# Ejecutar con hot reload
make dev

# Ejecutar tests
make test

# Ejecutar tests con cobertura
make test-coverage

# Generar documentación Swagger
make swagger

# Crear nueva migración
make migrate-create name=nombre_migracion

# Ejecutar migraciones
make migrate-up

# Revertir migraciones
make migrate-down
```

### Estructura de Respuestas

Todas las respuestas siguen un formato estándar:

```json
{
  "success": true,
  "data": {
    // Datos de la respuesta
  },
  "message": "Operación exitosa",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

En caso de error:

```json
{
  "success": false,
  "error": "Descripción del error",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

## Próximos Pasos

1. Implementar handlers para todos los endpoints
2. Agregar validación de datos
3. Implementar sistema de caché con Redis
4. Agregar tests unitarios y de integración
5. Configurar CI/CD
6. Implementar logging estructurado
7. Agregar métricas y monitoreo 