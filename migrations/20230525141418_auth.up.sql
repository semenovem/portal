-- Авторизованные сессии
CREATE TABLE IF NOT EXISTS auth.sessions
(
  id         serial PRIMARY KEY,
  user_id    int CHECK (user_id > 0) NOT NULL,
  device_id  uuid                    NOT NULL,
  created_at timestamp default now() NOT NULL,
  logouted   bool      default false NOT NULL, -- вышел из системы или был разлогинен
  refresh_id uuid                    NOT NULL
);
