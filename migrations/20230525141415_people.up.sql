CREATE TYPE people.roles_enum AS ENUM ('super-admin', 'admin', 'audit', 'guest', 'operator', 'boss');
CREATE TYPE people.statuses_enum AS ENUM ('active', 'inactive', 'blocked', 'registration');
CREATE TYPE people.additional_field_kind_enum AS ENUM (
  'email-main', 'email-personal', 'tel-work', 'tel-personal', 'note');

-- Должности
CREATE TABLE IF NOT EXISTS people.positions
(
  id          smallserial PRIMARY KEY,
  title       varchar UNIQUE NOT NULL,                      -- название должности
  description varchar        NOT NULL,                      -- Описание
  parent_id   int REFERENCES people.positions default NULL, -- руководитель
  deleted     timestamp                       default NULL
);

-- Пользователи
CREATE TABLE IF NOT EXISTS people.users
(
  id          serial PRIMARY KEY,
  firstname   varchar(128)                                  NOT NULL, -- Имя
  surname     varchar(128)               default ''         NOT NULL, -- Фамилия
  deleted     bool                       default false      NOT NULL,
  note        text                       default NULL,                -- примечание
  status      people.statuses_enum       default 'inactive' NOT NULL,
  roles       people.roles_enum[]        default NULL,
  avatar_id   int check ( avatar_id > 0) default NULL,
  expired_at  timestamp                  default NULL,                -- УЗ активна до указанного времени

  login       varchar(128) UNIQUE        default NULL,
  passwd_hash varchar(40)                default NULL,                -- хэш пароля
  props       jsonb                      default NULL                 -- данные пользователя
);

-- Сотрудники компании
CREATE TABLE IF NOT EXISTS people.employees
(
  user_id     int PRIMARY KEY REFERENCES people.users NOT NULL,
  position_id int REFERENCES people.positions         NOT NULL,
  worked_at   timestamp default now()                 NOT NULL, -- дата начала работы
  fired_at    timestamp default NULL                            -- дата увольнения (последний день работы)
);

-- Наборы документов пользователей
CREATE TABLE IF NOT EXISTS people.user_media_boxes
(
  user_id     int PRIMARY KEY REFERENCES people.users NOT NULL,
  position_id int REFERENCES people.positions         NOT NULL,
  worked_at   timestamp default now()                 NOT NULL, -- дата начала работы
  fired_at    timestamp default NULL                  NULL      -- дата увольнения (последний день работы)
);


--------------------------------------------------------------------------------
-------------------------------   init data   ----------------------------------
--------------------------------------------------------------------------------

insert into people.positions (title, description, parent_id)
values ('водитель', 'описание должности 1', null),
       ('водитель-экспедитор', 'описание должности 2', null),
       ('экспедитор', 'описание должности 2', null),
       ('грузчик-экспедитор', 'описание должности 2', null),
       ('Руководитель отдела транспорта', '', null),
       ('Генеральный директор', '', null),
       ('Руководитель АТИ', '', null),
       ('Оператор', '', null),
       ('грузчик', '', null)
on conflict do nothing;


--------------------------------------------------------------------------------
-------------------------------   test data   ----------------------------------
--------------------------------------------------------------------------------

insert into people.users (firstname, surname, note, status, login, passwd_hash)
values ('Петр', 'Петрович', '', 'active', null, null),
       ('Иван', 'Сидорович', '', 'active', 'login1',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b'),
       ('ivan', 'ivanov', 'note для пользователя', 'inactive', 'login2',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b'),
       ('oleg', 'olegovich', 'note2 для пользователя', 'active', 'login3',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b'),
       ('Макс', 'Масков', 'note(макс) для пользователя', 'active', 'login4',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b')

on conflict do nothing;

insert into people.employees (user_id, position_id, worked_at, fired_at)
values (3, 1, '2023-07-12T15:38:30Z', now()),
       (4, 2, '2022-08-12T15:38:30Z', NULL)
on conflict do nothing;
