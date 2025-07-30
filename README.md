# Koopi - Sistema de Gestión de Inventarios

Koopi es una aplicación para gestionar inventarios de objetos organizados jerárquicamente. Permite registrar objetos en ubicaciones específicas dentro de una estructura anidada (ej: Casa → Habitación → Armario → Cajón).

## 🏗️ Arquitectura

- **Backend**: NestJS con TypeScript
- **Base de Datos**: PostgreSQL (Docker)
- **Documentación API**: Swagger/OpenAPI (Docker)
- **Autenticación**: JWT
- **Validación**: class-validator
- **Migraciones**: TypeORM Migrations

## 📋 Plan de Desarrollo

### Fase 1: Configuración Base ✅
- [x] Proyecto NestJS inicializado
- [x] CORS configurado
- [ ] Configuración de Docker para PostgreSQL
- [ ] Configuración de variables de entorno (.env)
- [ ] Configuración de TypeORM con migraciones
- [ ] Configuración de validación (class-validator, class-transformer)
- [ ] Configuración de Swagger/OpenAPI con Docker
- [ ] Sistema de migraciones automáticas

### Fase 2: Modelos de Datos
- [ ] Entidad `Location` (ubicaciones: casa, habitación, armario, etc.)
- [ ] Entidad `Item` (objetos/artículos)
- [ ] Entidad `User` (usuarios del sistema)
- [ ] Relaciones entre entidades
- [ ] Migración inicial de base de datos
- [ ] Scripts de migración para cambios de esquema

### Fase 3: Autenticación y Autorización
- [ ] Módulo de autenticación (AuthModule)
- [ ] JWT Strategy
- [ ] Guards para rutas protegidas
- [ ] DTOs para login/registro
- [ ] Endpoints de autenticación

### Fase 4: Gestión de Ubicaciones
- [ ] CRUD completo para ubicaciones
- [ ] Estructura jerárquica (árbol de ubicaciones)
- [ ] Validación de jerarquía
- [ ] Endpoints para ubicaciones
- [ ] Documentación Swagger para ubicaciones

### Fase 5: Gestión de Objetos
- [ ] CRUD completo para objetos
- [ ] Asociación con ubicaciones
- [ ] Búsqueda y filtros
- [ ] Endpoints para objetos
- [ ] Documentación Swagger para objetos

### Fase 6: Funcionalidades Avanzadas
- [ ] Búsqueda global de objetos
- [ ] Filtros por ubicación
- [ ] Estadísticas de inventario
- [ ] Exportación de datos
- [ ] Importación de datos

### Fase 7: Documentación API
- [ ] Configuración completa de Swagger
- [ ] Documentación de todos los endpoints
- [ ] Ejemplos de uso
- [ ] Esquemas de respuesta
- [ ] Generación de archivo JSON de Swagger

### Fase 8: Testing
- [ ] Tests unitarios
- [ ] Tests de integración
- [ ] Tests e2e
- [ ] Cobertura de código

### Fase 9: Optimización y Producción
- [ ] Optimización de consultas
- [ ] Caché (Redis)
- [ ] Logging
- [ ] Monitoreo
- [ ] Configuración de producción

## 🚀 Instalación y Configuración

### Prerrequisitos
- Node.js (v18+)
- Docker y Docker Compose
- npm o yarn

### Pasos de Instalación

1. **Clonar el repositorio**
```bash
git clone <repository-url>
cd koopi-backend
```

2. **Instalar dependencias**
```bash
npm install
```

3. **Configurar variables de entorno**
```bash
cp .env.example .env
# Editar .env con tus configuraciones

# Variables de entorno para Docker:
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USERNAME=koopi_user
DATABASE_PASSWORD=koopi_password
DATABASE_NAME=koopi_db
DATABASE_URL=postgresql://koopi_user:koopi_password@localhost:5432/koopi_db

# Configuración de migraciones
TYPEORM_MIGRATIONS_DIR=./migrations
TYPEORM_MIGRATIONS_RUN=true
```

4. **Configurar servicios con Docker**
```bash
# Levantar servicios (PostgreSQL, Redis, etc.)
docker-compose up -d

# Esperar a que PostgreSQL esté listo
docker-compose logs -f postgres

# Ejecutar migraciones
npm run migration:run
```

5. **Ejecutar la aplicación**
```bash
# Desarrollo
npm run start:dev

# Producción
npm run build
npm run start:prod
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

El proyecto implementa un sistema completo de migraciones con TypeORM:

#### Comandos de migración:
```bash
# Generar nueva migración
npm run migration:generate -- -n NombreMigracion

# Ejecutar migraciones pendientes
npm run migration:run

# Revertir última migración
npm run migration:revert

# Ver estado de migraciones
npm run migration:show
```

#### Flujo de trabajo para cambios en BD:
1. **Desarrollo**: Modificar entidades en el código
2. **Generar migración**: `npm run migration:generate -- -n DescripcionCambio`
3. **Revisar**: Verificar el archivo de migración generado
4. **Ejecutar**: `npm run migration:run`
5. **Commit**: Incluir migración en el control de versiones

#### Ejemplo de migración:
```typescript
// migrations/1234567890-AddUserTable.ts
export class AddUserTable1234567890 implements MigrationInterface {
    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`
            CREATE TABLE "user" (
                "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
                "email" character varying NOT NULL,
                "password" character varying NOT NULL,
                "name" character varying NOT NULL,
                "createdAt" TIMESTAMP NOT NULL DEFAULT now(),
                "updatedAt" TIMESTAMP NOT NULL DEFAULT now(),
                CONSTRAINT "PK_user" PRIMARY KEY ("id")
            )
        `);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`DROP TABLE "user"`);
    }
}
```

## 📚 Documentación API

La documentación de la API estará disponible en:
- **Desarrollo**: http://localhost:3001/api
- **Producción**: https://tu-dominio.com/api
- **Docker**: http://localhost:8080 (Swagger UI)

### Generación de archivo Swagger JSON
```bash
npm run swagger:generate
```

### Acceso a Swagger con Docker
```bash
# Swagger UI disponible en Docker
docker-compose up swagger-ui
```

## 🗂️ Estructura del Proyecto

```
src/
├── auth/                 # Módulo de autenticación
├── locations/            # Módulo de ubicaciones
├── items/               # Módulo de objetos
├── users/               # Módulo de usuarios
├── common/              # Utilidades comunes
├── config/              # Configuraciones
└── database/            # Configuración de BD y migraciones

docker/
├── postgres/            # Configuración PostgreSQL
├── redis/               # Configuración Redis (opcional)
└── swagger/             # Configuración Swagger UI

migrations/              # Migraciones de TypeORM
```

## 📊 Modelos de Datos

### Location (Ubicación)
```typescript
{
  id: string;
  name: string;
  description?: string;
  parentId?: string;     // Ubicación padre
  userId: string;        // Propietario
  createdAt: Date;
  updatedAt: Date;
}
```

### Item (Objeto)
```typescript
{
  id: string;
  name: string;
  description?: string;
  locationId: string;    // Ubicación donde está guardado
  userId: string;        // Propietario
  createdAt: Date;
  updatedAt: Date;
}
```

### User (Usuario)
```typescript
{
  id: string;
  email: string;
  password: string;
  name: string;
  createdAt: Date;
  updatedAt: Date;
}
```

## 🔐 Endpoints Principales

### Autenticación
- `POST /auth/register` - Registro de usuario
- `POST /auth/login` - Inicio de sesión
- `POST /auth/refresh` - Renovar token

### Ubicaciones
- `GET /locations` - Listar ubicaciones
- `POST /locations` - Crear ubicación
- `GET /locations/:id` - Obtener ubicación
- `PUT /locations/:id` - Actualizar ubicación
- `DELETE /locations/:id` - Eliminar ubicación
- `GET /locations/:id/items` - Objetos en ubicación

### Objetos
- `GET /items` - Listar objetos
- `POST /items` - Crear objeto
- `GET /items/:id` - Obtener objeto
- `PUT /items/:id` - Actualizar objeto
- `DELETE /items/:id` - Eliminar objeto
- `GET /items/search` - Buscar objetos

## 🧪 Testing

```bash
# Tests unitarios
npm run test

# Tests e2e
npm run test:e2e

# Cobertura
npm run test:cov
```

## 📝 Scripts Disponibles

```bash
# Desarrollo
npm run start:dev      # Desarrollo con hot reload
npm run build          # Compilar para producción
npm run start:prod     # Ejecutar en producción

# Testing
npm run test           # Tests unitarios
npm run test:e2e       # Tests e2e
npm run test:cov       # Tests con cobertura

# Base de datos y migraciones
npm run migration:run  # Ejecutar migraciones
npm run migration:generate  # Generar migración
npm run migration:revert    # Revertir última migración
npm run migration:show      # Mostrar estado de migraciones

# Documentación
npm run swagger:generate    # Generar JSON de Swagger

# Docker
docker-compose up -d        # Levantar servicios
docker-compose down         # Parar servicios
docker-compose logs         # Ver logs
```

## 🤝 Contribución

1. Fork el proyecto
2. Crear una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abrir un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para detalles.

## 🔄 Control de Versiones de Base de Datos

### Buenas Prácticas para Migraciones:

1. **Nunca modificar migraciones existentes** - Crear nuevas migraciones en su lugar
2. **Revisar migraciones antes de ejecutar** - Verificar que el SQL generado es correcto
3. **Probar migraciones en desarrollo** - Antes de aplicar en producción
4. **Incluir migraciones en commits** - Las migraciones son parte del código
5. **Documentar cambios complejos** - Añadir comentarios en migraciones importantes

### Comandos útiles para desarrollo:

```bash
# Ver migraciones pendientes
npm run migration:show

# Ejecutar migraciones en modo verbose
npm run migration:run -- --verbose

# Generar migración con timestamp específico
npm run migration:generate -- -n AddNewField -t 20240101000000

# Revertir múltiples migraciones
npm run migration:revert -- --count 3
```

## 📞 Contacto

- **Desarrollador**: [Tu Nombre]
- **Email**: [tu-email@ejemplo.com]
- **Proyecto**: [https://github.com/usuario/koopi-backend]
