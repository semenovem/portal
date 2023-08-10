CREATE TYPE people.roles_enum AS ENUM ('admin', 'guest', 'operator', 'boss');
CREATE TYPE people.statuses_enum AS ENUM ('active', 'inactive', 'blocked', 'registration');
CREATE TYPE people.additional_field_kind_enum AS ENUM ('email', 'tel', 'note');

-- Должности
CREATE TABLE IF NOT EXISTS people.positions
(
  id          serial PRIMARY KEY,
  title       varchar UNIQUE                                NOT NULL, -- название должности
  description varchar                                       NOT NULL, -- Описание
  parent_id   int REFERENCES people.positions default NULL  NULL,     -- руководитель
  created_at  timestamp                       default now() NOT NULL,
  deleted_at  timestamp                       default NULL  NULL
);

-- Пользователи
CREATE TABLE IF NOT EXISTS people.users
(
  id          serial PRIMARY KEY,
  firstname   varchar                            NOT NULL, -- Имя
  surname     varchar                            NOT NULL, -- Фамилия
  created_at  timestamp            default now() NOT NULL, -- дата создания записи
  updated_at  timestamp            default now() NOT NULL, -- обновление данных
  deleted_at  timestamp            default NULL  NULL,
  note        text                 default ''    NOT NULL, -- примечание
  position_id int REFERENCES people.positions    NOT NULL,
  status      people.statuses_enum default 'inactive',
  roles       people.roles_enum[]  default NULL  NULL
);

-- Имеют дополнительные поля
CREATE TABLE IF NOT EXISTS people.ext_users
(
  login       varchar UNIQUE                                 NOT NULL,
  email       varchar                           default NULL NULL,
  passwd_hash varchar(64)                       default NULL NULL, -- хэш пароля
  avatar      int REFERENCES media.avatars (id) default NULL NULL
) INHERITS (people.users);


-- дополнительные поля пользователя
CREATE TABLE IF NOT EXISTS people.user_additional_fields
(
  id         serial PRIMARY KEY,
  user_id    int REFERENCES people.users NOT NULL,
  value      varchar   default ''        NOT NULL,
  kind       people.additional_field_kind_enum  NOT NULL, -- тип поля
  sort       smallint  default 0         NOT NULL,
  deleted_at timestamp default NULL      NULL
);

-- --------------------------------
-- --------------------------------
-- --------------------------------
-- test data
-- --------------------------------
-- --------------------------------
-- --------------------------------


insert into people.positions (title, description)
values ('должность 1', 'описание должности 1'),
       ('водитель-экспедитор', 'описание должности 2')
on conflict do nothing;


insert into people.ext_users
  (firstname, surname, note, position_id, login)
values ('ivan', 'ivanov', 'note для пользователя', 1, 'login1'),
       ('oleg', 'olegovich', 'note2 для пользователя', 1, 'login2')
on conflict do nothing;

insert into people.users (firstname, surname, position_id)
values ('Петр', 'Петрович', 2)
on conflict do nothing;
