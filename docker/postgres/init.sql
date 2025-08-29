-- Script de inicialización para PostgreSQL
-- Este script se ejecuta cuando el contenedor se inicia por primera vez

-- Crear extensión para UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Nota: Los datos iniciales se insertarán después de ejecutar las migraciones
-- desde la aplicación o usando el comando make migrate-up 