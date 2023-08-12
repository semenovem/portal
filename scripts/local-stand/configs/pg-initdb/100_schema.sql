DROP SCHEMA public;

-- тестовая БД
CREATE USER adm_db WITH PASSWORD 'airohZ9o';
CREATE SCHEMA core;

CREATE TABLE core.schema_migrations
(
  version bigint NOT NULL,
  dirty   bool   NOT NULL
);

ALTER TABLE core.schema_migrations
  OWNER TO adm_db;

ALTER ROLE adm_db SET search_path = core;
ALTER SCHEMA core OWNER TO adm_db;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA core TO adm_db;

-- users
CREATE SCHEMA people;
ALTER SCHEMA people OWNER TO adm_db;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA people TO adm_db;

-- auth
CREATE SCHEMA auth;
ALTER SCHEMA auth OWNER TO adm_db;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA auth TO adm_db;

-- vehicle
CREATE SCHEMA vehicle;
ALTER SCHEMA vehicle OWNER TO adm_db;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA auth TO adm_db;

-- vehicle
CREATE SCHEMA media;
ALTER SCHEMA media OWNER TO adm_db;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA media TO adm_db;


-- audit - для временного хранения данных аудита до момента их отправки в систему хранения
CREATE SCHEMA audit;
ALTER SCHEMA audit OWNER TO adm_db;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA audit TO adm_db;

