# Koopi - Sistema de Gestión de Inventarios

Koopi es una aplicación para gestionar inventarios de objetos organizados jerárquicamente. Permite registrar objetos en ubicaciones específicas dentro de una estructura anidada (ej: Casa → Habitación → Armario → Cajón).

## 🎯 Niveles de Gestión

### Nivel Personal (Privado)
- **Propósito**: Gestión personal de espacios y objetos
- **Acceso**: Solo el propietario
- **Ejemplo**: Tu casa, tu oficina personal, tus espacios privados
- **Funcionalidades**: CRUD completo, búsqueda personal, estadísticas privadas

### Nivel Organizacional (Compartido)
- **Propósito**: Gestión compartida de espacios y objetos
- **Acceso**: Múltiples usuarios con diferentes roles
- **Ejemplo**: Banda musical, empresa, organización, grupo de trabajo
- **Funcionalidades**: Roles y permisos, auditoría, gestión de miembros

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
- [ ] Entidad `User` (usuarios del sistema)
- [ ] Entidad `Organization` (organizaciones/grupos)
- [ ] Entidad `OrganizationMember` (miembros y roles)
- [ ] Entidad `Location` (ubicaciones: casa, habitación, armario, etc.)
- [ ] Entidad `Item` (objetos/artículos)
- [ ] Relaciones entre entidades
- [ ] Migración inicial de base de datos
- [ ] Scripts de migración para cambios de esquema

### Fase 3: Autenticación y Autorización
- [ ] Módulo de autenticación (AuthModule)
- [ ] JWT Strategy
- [ ] Guards para rutas protegidas
- [ ] DTOs para login/registro
- [ ] Endpoints de autenticación
- [ ] Sistema de roles y permisos
- [ ] Guards para verificación de permisos organizacionales

### Fase 4: Gestión de Organizaciones
- [ ] CRUD completo para organizaciones
- [ ] Gestión de miembros y roles
- [ ] Invitaciones a organizaciones
- [ ] Endpoints para organizaciones
- [ ] Documentación Swagger para organizaciones

### Fase 5: Gestión de Ubicaciones
- [ ] CRUD completo para ubicaciones (personal y organizacional)
- [ ] Estructura jerárquica (árbol de ubicaciones)
- [ ] Validación de jerarquía
- [ ] Separación por propietario (usuario u organización)
- [ ] Endpoints para ubicaciones
- [ ] Documentación Swagger para ubicaciones

### Fase 6: Gestión de Objetos
- [ ] CRUD completo para objetos (personal y organizacional)
- [ ] Asociación con ubicaciones
- [ ] Búsqueda y filtros
- [ ] Separación por propietario (usuario u organización)
- [ ] Endpoints para objetos
- [ ] Documentación Swagger para objetos

### Fase 7: Funcionalidades Avanzadas
- [ ] Búsqueda global de objetos (personal y organizacional)
- [ ] Filtros por ubicación
- [ ] Estadísticas de inventario
- [ ] Exportación de datos
- [ ] Importación de datos
- [ ] Auditoría de cambios (especialmente para organizaciones)
- [ ] Notificaciones de cambios

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
├── organizations/        # Módulo de organizaciones
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

### Organization (Organización)
```typescript
{
  id: string;
  name: string;
  description?: string;
  ownerId: string;       // Usuario propietario
  isPublic: boolean;     // Si es pública o privada
  createdAt: Date;
  updatedAt: Date;
}
```

### OrganizationMember (Miembro de Organización)
```typescript
{
  id: string;
  organizationId: string;
  userId: string;
  role: 'OWNER' | 'ADMIN' | 'MEMBER' | 'VIEWER';
  joinedAt: Date;
  invitedBy?: string;    // Usuario que invitó
}
```

### Location (Ubicación)
```typescript
{
  id: string;
  name: string;
  description?: string;
  parentId?: string;     // Ubicación padre
  ownerType: 'USER' | 'ORGANIZATION';
  ownerId: string;       // ID del usuario u organización
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
  ownerType: 'USER' | 'ORGANIZATION';
  ownerId: string;       // ID del usuario u organización
  assignedTo?: string;   // Usuario asignado (para organizaciones)
  createdAt: Date;
  updatedAt: Date;
}
```

## 🔐 Endpoints Principales

### Autenticación
- `POST /auth/register` - Registro de usuario
- `POST /auth/login` - Inicio de sesión
- `POST /auth/refresh` - Renovar token

### Organizaciones
- `GET /organizations` - Listar organizaciones del usuario
- `POST /organizations` - Crear organización
- `GET /organizations/:id` - Obtener organización
- `PUT /organizations/:id` - Actualizar organización
- `DELETE /organizations/:id` - Eliminar organización
- `GET /organizations/:id/members` - Listar miembros
- `POST /organizations/:id/members` - Invitar miembro
- `PUT /organizations/:id/members/:userId` - Actualizar rol
- `DELETE /organizations/:id/members/:userId` - Expulsar miembro

### Ubicaciones
- `GET /locations` - Listar ubicaciones (personales y organizacionales)
- `POST /locations` - Crear ubicación
- `GET /locations/:id` - Obtener ubicación
- `PUT /locations/:id` - Actualizar ubicación
- `DELETE /locations/:id` - Eliminar ubicación
- `GET /locations/:id/items` - Objetos en ubicación
- `GET /locations/personal` - Solo ubicaciones personales
- `GET /locations/organization/:orgId` - Ubicaciones de organización

### Objetos
- `GET /items` - Listar objetos (personales y organizacionales)
- `POST /items` - Crear objeto
- `GET /items/:id` - Obtener objeto
- `PUT /items/:id` - Actualizar objeto
- `DELETE /items/:id` - Eliminar objeto
- `GET /items/search` - Buscar objetos
- `GET /items/personal` - Solo objetos personales
- `GET /items/organization/:orgId` - Objetos de organización
- `PUT /items/:id/assign` - Asignar objeto a usuario (organizaciones)

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

## 🔐 Roles y Permisos

### Roles en Organizaciones:
- **OWNER**: Propietario de la organización
  - Puede gestionar miembros
  - Puede eliminar la organización
  - Acceso completo a todos los recursos
- **ADMIN**: Administrador
  - Puede gestionar miembros (excepto otros admins)
  - Puede gestionar ubicaciones y objetos
  - No puede eliminar la organización
- **MEMBER**: Miembro activo
  - Puede crear, editar y eliminar objetos
  - Puede ver ubicaciones y otros objetos
  - No puede gestionar miembros
- **VIEWER**: Solo lectura
  - Puede ver objetos y ubicaciones
  - No puede crear, editar o eliminar
  - Útil para consultas y auditoría

### Permisos por Nivel:
- **Personal**: Solo el propietario tiene acceso
- **Organizacional**: Según el rol asignado
- **Público**: Organizaciones marcadas como públicas (solo lectura)

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
