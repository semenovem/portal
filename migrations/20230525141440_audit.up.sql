-- Временное хранение аудита действий пользователей
CREATE TABLE IF NOT EXISTS audit.audits
(
  id         uuid unique             NOT NULL,
  created_at timestamp default now() NOT NULL,
  user_id    int CHECK (user_id > 0) NOT NULL, -- действие этого пользователя
  operation  varchar                 NOT NULL, -- тип события
  action     varchar                 NOT NULL, -- действие (create,update, delete и тд)
  payload    jsonb                   NOT NULL, -- данные операции
  entity_id  int                     NOT NULL  -- ID затронутой сущности
);

-- Временное хранение аудита операций логина
CREATE TABLE IF NOT EXISTS audit.auth_audits
(
  id         uuid unique             NOT NULL,
  created_at timestamp default now() NOT NULL,
  user_id    int CHECK (user_id > 0) NOT NULL, -- действие этого пользователя
  code       varchar                 NOT NULL, -- тип события
  payload    jsonb                   NOT NULL  -- данные операции
);

