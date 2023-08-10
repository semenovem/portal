-- Авторизованные сессии
CREATE TABLE IF NOT EXISTS auth.sessions
(
  id         serial PRIMARY KEY,
  user_id    integer REFERENCES people.users NOT NULL,
  created_at timestamp default now()         NOT NULL,
  deleted_at timestamp default NULL          NULL
);

-- Refresh токены
CREATE TABLE IF NOT EXISTS auth.refresh
(
  id            serial PRIMARY KEY,
  session_id    integer REFERENCES auth.sessions NOT NULL,
  created_at    timestamp default now()          NOT NULL,
  valid_till_at timestamp                        NOT NULL
);
