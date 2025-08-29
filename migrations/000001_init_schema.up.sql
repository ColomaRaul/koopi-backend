-- Crear extensión para UUID si no existe
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabla de usuarios
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de organizaciones
CREATE TABLE organizations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    owner_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    is_verified BOOLEAN DEFAULT false,
    website VARCHAR(255),
    phone VARCHAR(50),
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de categorías de eventos
CREATE TABLE event_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    icon VARCHAR(100),
    color VARCHAR(7),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de ubicaciones
CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(500) NOT NULL,
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de eventos
CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category_id INTEGER NOT NULL REFERENCES event_categories(id),
    organizer_id INTEGER NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    location_id INTEGER NOT NULL REFERENCES locations(id),
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    price DECIMAL(10,2) DEFAULT 0,
    capacity INTEGER,
    is_active BOOLEAN DEFAULT true,
    is_free BOOLEAN DEFAULT true,
    image_url VARCHAR(500),
    website VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de reservas
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(50) DEFAULT 'pending',
    quantity INTEGER DEFAULT 1,
    total_price DECIMAL(10,2),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Índices para optimizar consultas
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_organizations_owner_id ON organizations(owner_id);
CREATE INDEX idx_events_category_id ON events(category_id);
CREATE INDEX idx_events_organizer_id ON events(organizer_id);
CREATE INDEX idx_events_location_id ON events(location_id);
CREATE INDEX idx_events_start_date ON events(start_date);
CREATE INDEX idx_events_is_active ON events(is_active);
CREATE INDEX idx_bookings_event_id ON bookings(event_id);
CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_bookings_status ON bookings(status);

-- Índice espacial para ubicaciones (PostGIS si está disponible)
-- CREATE INDEX idx_locations_geom ON locations USING GIST (ST_SetSRID(ST_MakePoint(longitude, latitude), 4326));

-- Insertar categorías de eventos por defecto
INSERT INTO event_categories (name, description, icon, color) VALUES
('Concierto', 'Eventos musicales y conciertos', 'music', '#FF6B6B'),
('Deportes', 'Eventos deportivos y competiciones', 'sports', '#4ECDC4'),
('Cultural', 'Eventos culturales y artísticos', 'culture', '#45B7D1'),
('Gastronomía', 'Eventos gastronómicos y ferias de comida', 'food', '#96CEB4'),
('Tecnología', 'Eventos tecnológicos y conferencias', 'tech', '#FFEAA7'),
('Educación', 'Eventos educativos y talleres', 'education', '#DDA0DD'),
('Negocios', 'Eventos empresariales y networking', 'business', '#98D8C8'),
('Entretenimiento', 'Eventos de entretenimiento y ocio', 'entertainment', '#F7DC6F'),
('Mercado', 'Mercados y ferias', 'market', '#BB8FCE'),
('Otros', 'Otros tipos de eventos', 'other', '#85C1E9'); 