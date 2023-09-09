CREATE TYPE people.roles_enum AS ENUM ('super-admin', 'admin', 'audit', 'guest', 'operator', 'boss');
CREATE TYPE people.statuses_enum AS ENUM ('active', 'inactive', 'blocked', 'registration');
CREATE TYPE people.additional_field_kind_enum AS ENUM (
  'email-main', 'email-personal', 'tel-work', 'tel-personal', 'note');

-- Должности
CREATE TABLE IF NOT EXISTS people.positions
(
  id          smallserial PRIMARY KEY,
  title       varchar UNIQUE NOT NULL,                      -- название должности
  description varchar        NOT NULL,                      -- описание
  parent_id   int REFERENCES people.positions default NULL, -- руководитель
  deleted     timestamp                       default NULL
);

-- Отделы
CREATE TABLE IF NOT EXISTS people.departments
(
  id          smallserial PRIMARY KEY,
  title       varchar UNIQUE NOT NULL,                        -- название департамента
  description varchar        NOT NULL,                        -- описание
  parent_id   int REFERENCES people.departments default NULL, -- руководитель
  deleted     timestamp                         default NULL
);

-- Пользователи
CREATE TABLE IF NOT EXISTS people.users
(
  id          serial PRIMARY KEY,
  deleted     bool                       default false      NOT NULL,
  firstname   varchar(128)                                  NOT NULL, -- Имя
  surname     varchar(128)               default ''         NOT NULL, -- Фамилия
  status      people.statuses_enum       default 'inactive' NOT NULL,
  roles       people.roles_enum[]        default NULL,
  note        text                       default NULL,                -- примечание
  avatar_id   int check ( avatar_id > 0) default NULL,
  expired_at  timestamp                  default NULL,                -- УЗ активна до указанного времени
  updated_at  timestamp                  default now()      NOT NULL, -- время последнего обновления

  login       varchar(128) UNIQUE        default NULL,
  passwd_hash varchar(40)                default NULL,                -- хэш пароля
  props       jsonb                      default NULL                 -- данные пользователя
);

-- Сотрудники компании
CREATE TABLE IF NOT EXISTS people.employees
(
  user_id     int PRIMARY KEY REFERENCES people.users                                             NOT NULL,
  position_id int REFERENCES people.positions                                                     NOT NULL,
  dept_id     int REFERENCES people.departments                                                   NOT NULL,
  worked_at   timestamp default now() check ( fired_at IS NULL OR worked_at < fired_at) NOT NULL, -- дата начала работы
  fired_at    timestamp default NULL check ( fired_at IS NULL OR fired_at > worked_at )                                         -- дата увольнения (последний день работы)
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

insert into people.positions (id, title, description, parent_id)
values (1, 'водитель', 'описание должности 1', null),
       (2, 'водитель-экспедитор', 'описание должности 2', null),
       (3, 'экспедитор', 'описание должности 2', null),
       (4, 'грузчик-экспедитор', 'описание должности 2', null),
       (5, 'Руководитель отдела транспорта', '', null),
       (6, 'Генеральный директор', '', null),
       (7, 'Руководитель АТИ', '', null),
       (8, 'Оператор', '', null),
       (9, 'грузчик', '', null)
on conflict do nothing;


insert into people.departments (id, title, description, parent_id)
values (1, 'Управление', '', null),
       (2, 'АТИ', 'междугородние перевозки', null),
       (3, 'Кадры', 'работа с персоналом', null),
       (4, 'Транспортный', 'работа с транспортом', null),
       (5, 'Сборка', '', null)
on conflict do nothing;


--------------------------------------------------------------------------------
-------------------------------   test data   ----------------------------------
--------------------------------------------------------------------------------

insert into people.users (firstname, surname, note, status, roles, login, passwd_hash)
values ('Петр', 'Петрович', '', 'active', '{super-admin}', 'login1',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b'),
       ('Иван', 'Сидорович', '', 'active', null, null, null),
       ('ivan', 'ivanov', 'note для пользователя', 'inactive', null, 'login3',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b'),
       ('oleg', 'olegovich', 'note2 для пользователя', 'active', null, 'login4',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b'),
       ('Макс', 'Масков', 'note(макс) для пользователя', 'active', null, 'login5',
        'ec95a5a1e2e7b82333340b5ec1db3e82e3a8ae9b')

on conflict do nothing;

insert into people.employees (user_id, position_id, dept_id, worked_at, fired_at)
values (3, 1, 1, '2023-07-12T15:38:30Z', now()),
       (4, 2, 2, '2022-08-12T15:38:30Z', null)
on conflict do nothing;
