-- Партнеры
CREATE TABLE IF NOT EXISTS people.partners
(
  user_id     int PRIMARY KEY REFERENCES people.users NOT NULL
);


-- Контакты сторонних людей
CREATE TABLE IF NOT EXISTS people.contacts
(
  id          smallserial PRIMARY KEY,
  title       varchar UNIQUE NOT NULL                      -- название должности
);
