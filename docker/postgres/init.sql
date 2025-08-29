-- Script de inicialización para PostgreSQL
-- Este script se ejecuta cuando el contenedor se inicia por primera vez

-- Crear extensión para UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Crear usuario admin por defecto (solo para desarrollo)
-- En producción, esto debería hacerse a través de la aplicación
INSERT INTO users (email, password, name, role) 
VALUES (
    'admin@koopi.com', 
    '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', -- password: password
    'Administrador',
    'admin'
) ON CONFLICT (email) DO NOTHING; 