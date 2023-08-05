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
