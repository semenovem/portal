-- Временное хранение аудита действий пользователей
CREATE TABLE IF NOT EXISTS audit.audits
(
  created_at timestamp default now()     NOT NULL,
  user_id    int REFERENCES people.users NOT NULL, -- действие этого пользователя
  operation  varchar                     NOT NULL, -- тип события
  action     varchar                     NOT NULL, -- действие (create,update, delete и тд)
  payload    jsonb                       NOT NULL, -- данные операции
  entity_id  int                         NOT NULL  -- ID затронутой сущности
);

