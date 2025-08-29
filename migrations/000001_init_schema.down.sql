-- Eliminar índices
DROP INDEX IF EXISTS idx_bookings_status;
DROP INDEX IF EXISTS idx_bookings_user_id;
DROP INDEX IF EXISTS idx_bookings_event_id;
DROP INDEX IF EXISTS idx_events_is_active;
DROP INDEX IF EXISTS idx_events_start_date;
DROP INDEX IF EXISTS idx_events_location_id;
DROP INDEX IF EXISTS idx_events_organizer_id;
DROP INDEX IF EXISTS idx_events_category_id;
DROP INDEX IF EXISTS idx_organizations_owner_id;
DROP INDEX IF EXISTS idx_users_role;
DROP INDEX IF EXISTS idx_users_email;

-- Eliminar tablas en orden inverso (por las foreign keys)
DROP TABLE IF EXISTS bookings;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS locations;
DROP TABLE IF EXISTS event_categories;
DROP TABLE IF EXISTS organizations;
DROP TABLE IF EXISTS users;

-- Eliminar extensión
DROP EXTENSION IF EXISTS "uuid-ossp"; 