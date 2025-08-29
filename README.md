# Koopi - Sistema de Gestión de Eventos

Koopi es una aplicación backend para gestionar eventos de toda clase en una zona, ciudad o región. Permite a los usuarios crear, buscar y gestionar eventos como conciertos, mercados navideños, partidos de fútbol, etc., con integración para búsqueda por Maps.

## 🎯 Funcionalidades Principales

### Gestión de Eventos
- **Creación de eventos**: Conciertos, mercados, deportes, culturales, etc.
- **Búsqueda geográfica**: Integración con Maps para búsqueda por ubicación
- **Categorización**: Diferentes tipos de eventos con filtros
- **Gestión de organizadores**: Usuarios y organizaciones que crean eventos
- **Sistema de reservas**: Para eventos que requieren inscripción

### Características Técnicas
- **API RESTful**: Endpoints bien documentados
- **Geolocalización**: Coordenadas GPS para cada evento
- **Búsqueda avanzada**: Por fecha, ubicación, categoría, precio
- **Autenticación**: Sistema de usuarios y organizadores
- **Escalabilidad**: Diseñado para manejar múltiples ciudades/regiones

## 🏗️ Arquitectura

- **Backend**: Go (Golang) con Gin framework
- **Base de Datos**: PostgreSQL
- **Documentación API**: Swagger/OpenAPI
- **Autenticación**: JWT
- **Validación**: Go validator
- **Migraciones**: SQL migrations con golang-migrate
- **Geolocalización**: Integración con servicios de Maps

## 📋 Plan de Desarrollo

### Fase 1: Configuración Base ✅
- [x] Proyecto Go inicializado
- [ ] Configuración de Docker para PostgreSQL
- [ ] Configuración de variables de entorno (.env)
- [ ] Configuración de base de datos
- [ ] Configuración de Gin framework
- [ ] Configuración de Swagger/OpenAPI
- [ ] Sistema de migraciones

### Fase 2: Modelos de Datos
- [ ] Entidad `User` (usuarios del sistema)
- [ ] Entidad `Organization` (organizadores de eventos)
- [ ] Entidad `Event` (eventos)
- [ ] Entidad `EventCategory` (categorías de eventos)
- [ ] Entidad `Location` (ubicaciones geográficas)
- [ ] Entidad `Booking` (reservas/inscripciones)
- [ ] Relaciones entre entidades
- [ ] Migración inicial de base de datos

### Fase 3: Autenticación y Autorización
- [ ] Módulo de autenticación
- [ ] JWT Strategy
- [ ] Middleware para rutas protegidas
- [ ] DTOs para login/registro
- [ ] Endpoints de autenticación
- [ ] Sistema de roles (usuario, organizador, admin)

### Fase 4: Gestión de Eventos
- [ ] CRUD completo para eventos
- [ ] Búsqueda por ubicación geográfica
- [ ] Filtros por categoría, fecha, precio
- [ ] Gestión de imágenes y multimedia
- [ ] Endpoints para eventos
- [ ] Documentación Swagger para eventos

### Fase 5: Sistema de Reservas
- [ ] CRUD para reservas/inscripciones
- [ ] Gestión de capacidad de eventos
- [ ] Confirmaciones y cancelaciones
- [ ] Notificaciones
- [ ] Endpoints para reservas

### Fase 6: Integración con Maps
- [ ] Endpoints para búsqueda geográfica
- [ ] Filtros por radio de distancia
- [ ] Optimización de consultas espaciales
- [ ] Integración con APIs de Maps

### Fase 7: Funcionalidades Avanzadas
- [ ] Búsqueda avanzada con múltiples filtros
- [ ] Estadísticas de eventos
- [ ] Sistema de recomendaciones
- [ ] Exportación de datos
- [ ] Auditoría de cambios
- [ ] Notificaciones push

### Fase 8: Documentación API
- [ ] Configuración completa de Swagger
- [ ] Documentación de todos los endpoints
- [ ] Ejemplos de uso
- [ ] Esquemas de respuesta
- [ ] Generación de archivo JSON de Swagger

### Fase 9: Testing
- [ ] Tests unitarios
- [ ] Tests de integración
- [ ] Tests e2e
- [ ] Cobertura de código

### Fase 10: Optimización y Producción
- [ ] Optimización de consultas
- [ ] Caché (Redis)
- [ ] Logging estructurado
- [ ] Monitoreo y métricas
- [ ] Configuración de producción

## 🚀 Instalación y Configuración

### Prerrequisitos
- Go (v1.21+)
- Docker y Docker Compose
- PostgreSQL

### Pasos de Instalación

1. **Clonar el repositorio**
```bash
git clone <repository-url>
cd koopi-backend
```

2. **Instalar dependencias**
```bash
go mod download
```

3. **Configurar variables de entorno**
```bash
cp .env.example .env
# Editar .env con tus configuraciones

# Variables de entorno para Docker:
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=koopi_user
DATABASE_PASSWORD=koopi_password
DATABASE_NAME=koopi_db
DATABASE_URL=postgresql://koopi_user:koopi_password@localhost:5432/koopi_db

# Configuración de JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# Configuración del servidor
SERVER_PORT=8080
SERVER_HOST=localhost
```

4. **Configurar servicios con Docker**
```bash
# Levantar servicios (PostgreSQL, Redis, etc.)
docker-compose up -d

# Esperar a que PostgreSQL esté listo
docker-compose logs -f postgres

# Ejecutar migraciones
make migrate-up
```

5. **Ejecutar la aplicación**
```bash
# Desarrollo
make run

# Producción
make build
make run-prod
```

## 🐳 Docker y Migraciones

### Configuración de Docker

El proyecto utiliza Docker para gestionar los servicios externos:

#### Servicios disponibles:
- **PostgreSQL**: Base de datos principal
- **Redis**: Caché (opcional)
- **Swagger UI**: Documentación API

#### Archivos de configuración:
- `docker-compose.yml` - Configuración de servicios
- `docker/postgres/` - Configuración específica de PostgreSQL
- `docker/swagger/` - Configuración de Swagger UI

### Sistema de Migraciones

El proyecto implementa un sistema de migraciones con golang-migrate:

#### Comandos de migración:
```bash
# Ejecutar migraciones pendientes
make migrate-up

# Revertir última migración
make migrate-down

# Crear nueva migración
make migrate-create name=migration_name

# Ver estado de migraciones
make migrate-status
```

## 📚 Documentación API

La documentación de la API estará disponible en:
- **Desarrollo**: http://localhost:8080/swagger/index.html
- **Producción**: https://tu-dominio.com/swagger/index.html
- **Docker**: http://localhost:8080/swagger/index.html

## 🗂️ Estructura del Proyecto

```
cmd/
├── server/              # Punto de entrada de la aplicación
└── migrate/             # Herramienta de migraciones

internal/
├── auth/                # Módulo de autenticación
├── events/              # Módulo de eventos
├── organizations/       # Módulo de organizaciones
├── users/               # Módulo de usuarios
├── bookings/            # Módulo de reservas
├── locations/           # Módulo de ubicaciones
├── common/              # Utilidades comunes
├── config/              # Configuraciones
├── database/            # Configuración de BD
├── middleware/          # Middlewares
└── handlers/            # Manejadores HTTP

pkg/
├── models/              # Modelos de datos
├── dto/                 # Data Transfer Objects
├── utils/               # Utilidades
└── validators/          # Validadores

migrations/              # Migraciones SQL
docker/                  # Configuración Docker
docs/                    # Documentación
```

## 📊 Modelos de Datos

### User (Usuario)
```go
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"`
    Name      string    `json:"name" gorm:"not null"`
    Role      string    `json:"role" gorm:"default:'user'"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Organization (Organización)
```go
type Organization struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"not null"`
    Description string    `json:"description"`
    OwnerID     uint      `json:"owner_id" gorm:"not null"`
    IsVerified  bool      `json:"is_verified" gorm:"default:false"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### Event (Evento)
```go
type Event struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" gorm:"not null"`
    Description string    `json:"description"`
    CategoryID  uint      `json:"category_id" gorm:"not null"`
    OrganizerID uint      `json:"organizer_id" gorm:"not null"`
    LocationID  uint      `json:"location_id" gorm:"not null"`
    StartDate   time.Time `json:"start_date" gorm:"not null"`
    EndDate     time.Time `json:"end_date" gorm:"not null"`
    Price       float64   `json:"price" gorm:"default:0"`
    Capacity    int       `json:"capacity"`
    IsActive    bool      `json:"is_active" gorm:"default:true"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### Location (Ubicación)
```go
type Location struct {
    ID        uint    `json:"id" gorm:"primaryKey"`
    Name      string  `json:"name" gorm:"not null"`
    Address   string  `json:"address" gorm:"not null"`
    City      string  `json:"city" gorm:"not null"`
    Latitude  float64 `json:"latitude" gorm:"not null"`
    Longitude float64 `json:"longitude" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Booking (Reserva)
```go
type Booking struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    EventID   uint      `json:"event_id" gorm:"not null"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    Status    string    `json:"status" gorm:"default:'pending'"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## 🔐 Endpoints Principales

### Autenticación
- `POST /auth/register` - Registro de usuario
- `POST /auth/login` - Inicio de sesión
- `POST /auth/refresh` - Renovar token
- `GET /auth/profile` - Obtener perfil

### Eventos
- `GET /events` - Listar eventos con filtros
- `POST /events` - Crear evento (organizadores)
- `GET /events/:id` - Obtener evento
- `PUT /events/:id` - Actualizar evento
- `DELETE /events/:id` - Eliminar evento
- `GET /events/nearby` - Eventos cercanos por ubicación
- `GET /events/search` - Búsqueda avanzada

### Organizaciones
- `GET /organizations` - Listar organizaciones
- `POST /organizations` - Crear organización
- `GET /organizations/:id` - Obtener organización
- `PUT /organizations/:id` - Actualizar organización
- `GET /organizations/:id/events` - Eventos de organización

### Reservas
- `GET /bookings` - Mis reservas
- `POST /bookings` - Crear reserva
- `GET /bookings/:id` - Obtener reserva
- `PUT /bookings/:id/cancel` - Cancelar reserva
- `GET /events/:id/bookings` - Reservas de evento (organizador)

### Ubicaciones
- `GET /locations` - Listar ubicaciones
- `POST /locations` - Crear ubicación
- `GET /locations/:id` - Obtener ubicación
- `GET /locations/search` - Buscar ubicaciones

## 🧪 Testing

```bash
# Tests unitarios
make test

# Tests de integración
make test-integration

# Cobertura
make test-coverage
```

## 📝 Scripts Disponibles (Makefile)

```bash
# Desarrollo
make run              # Ejecutar en desarrollo
make build            # Compilar para producción
make run-prod         # Ejecutar en producción

# Testing
make test             # Tests unitarios
make test-integration # Tests de integración
make test-coverage    # Tests con cobertura

# Base de datos y migraciones
make migrate-up       # Ejecutar migraciones
make migrate-down     # Revertir migraciones
make migrate-create   # Crear nueva migración
make migrate-status   # Estado de migraciones

# Documentación
make swagger          # Generar documentación Swagger

# Docker
make docker-up        # Levantar servicios
make docker-down      # Parar servicios
make docker-logs      # Ver logs
```

## 🤝 Contribución

1. Fork el proyecto
2. Crear una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abrir un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para detalles.

## 🔐 Roles y Permisos

### Roles en el Sistema:
- **USER**: Usuario básico
  - Puede ver eventos públicos
  - Puede crear reservas
  - Puede ver su perfil y reservas
- **ORGANIZER**: Organizador de eventos
  - Puede crear y gestionar eventos
  - Puede ver reservas de sus eventos
  - Puede gestionar su organización
- **ADMIN**: Administrador del sistema
  - Acceso completo a todos los recursos
  - Puede gestionar usuarios y organizaciones
  - Puede ver estadísticas del sistema

## 📞 Contacto

- **Desarrollador**: [Tu Nombre]
- **Email**: [tu-email@ejemplo.com]
- **Proyecto**: [https://github.com/usuario/koopi-backend]
