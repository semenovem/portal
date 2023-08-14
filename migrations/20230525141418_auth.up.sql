-- Авторизованные сессии
CREATE TABLE IF NOT EXISTS auth.sessions
(
  id         serial PRIMARY KEY,
  user_id    integer REFERENCES people.users NOT NULL,
  device_id  uuid                            NOT NULL,
  created_at timestamp default now()         NOT NULL,
  logout_at  timestamp default NULL          NULL,
  refresh_id uuid                            NOT NULL
);