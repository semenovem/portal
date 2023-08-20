CREATE TYPE people.roles_enum AS ENUM ('admin', 'guest', 'operator', 'boss');
CREATE TYPE people.statuses_enum AS ENUM ('active', 'inactive', 'blocked', 'registration');
CREATE TYPE people.additional_field_kind_enum AS ENUM (
  'email-main', 'email-personal', 'tel-work', 'tel-personal', 'note');

-- Должности
CREATE TABLE IF NOT EXISTS people.positions
(
  id          smallserial PRIMARY KEY,
  title       varchar UNIQUE                               NOT NULL, -- название должности
  description varchar                                      NOT NULL, -- Описание
  parent_id   int REFERENCES people.positions default NULL NULL,     -- руководитель
  deleted     timestamp                       default NULL NULL
);

-- Пользователи
CREATE TABLE IF NOT EXISTS people.users
(
  id          serial PRIMARY KEY,
  firstname   varchar                                         NOT NULL, -- Имя
  surname     varchar                           default ''    NOT NULL, -- Фамилия
  deleted     bool                              default false NOT NULL,
  note        text                              default ''    NOT NULL, -- примечание
  position    varchar                           default ''    NOT NULL,
  status      people.statuses_enum              default 'inactive',
  roles       people.roles_enum[]               default NULL  NULL,
  avatar      int REFERENCES media.avatars (id) default NULL  NULL,
  expired_at  timestamp                         default NULL  NULL,     -- УЗ активна до указанного времени

  login       varchar(128) UNIQUE               default NULL  NULL,
  passwd_hash varchar(40)                       default ''    NOT NULL  -- хэш пароля
);

-- Сотрудники компании
CREATE TABLE IF NOT EXISTS people.employees
(
  position_id int REFERENCES people.positions NOT NULL,
  worked_at   timestamp default now()         NOT NULL, -- дата начала работы
  fired_at    timestamp default NULL          NULL      -- дата увольнения (последний день работы)
) INHERITS (people.users);

-- Дополнительные поля пользователя
CREATE TABLE IF NOT EXISTS people.user_additional_fields
(
  id         serial PRIMARY KEY,
  user_id    int REFERENCES people.users       NOT NULL,
  value      varchar   default ''              NOT NULL,
  kind       people.additional_field_kind_enum NOT NULL, -- тип поля
  sort       smallint  default 0               NOT NULL,
  deleted_at timestamp default NULL            NULL
);

-- ----------------------------------------------------------------
-- test data
-- ----------------------------------------------------------------


insert into people.positions (title, description)
values ('должность 1', 'описание должности 1'),
       ('водитель-экспедитор', 'описание должности 2'),
       ('грузчик', '')
on conflict do nothing;


insert into people.employees
(firstname, surname, note, status, login, passwd_hash, position_id, worked_at, fired_at)
values ('ivan', 'ivanov', 'note для пользователя', 'inactive', 'login1',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b', 1, '2023-07-12T15:38:30Z', now()),

       ('oleg', 'olegovich', 'note2 для пользователя', 'active', 'login2',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b', 2, '2022-08-12T15:38:30Z', NULL)
on conflict do nothing;

insert into people.users (firstname, surname, position)
values ('Петр', 'Петрович', 2)
on conflict do nothing;
