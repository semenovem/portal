CREATE TYPE people.roles_enum AS ENUM ('super-admin', 'admin', 'audit', 'guest', 'operator', 'boss');
CREATE TYPE people.statuses_enum AS ENUM ('active', 'inactive', 'blocked', 'registration');
CREATE TYPE people.additional_field_kind_enum AS ENUM (
  'email-main', 'email-personal', 'tel-work', 'tel-personal', 'note');

-- Отделы
CREATE TABLE IF NOT EXISTS people.departments
(
  id          smallserial PRIMARY KEY,
  title       varchar UNIQUE                                  NOT NULL, -- название департамента
  description varchar                                         NOT NULL, -- описание
  parent_id   int REFERENCES people.departments default NULL,           -- руководитель
  deleted     bool                              default false NOT NULL
);

-- Должности
CREATE TABLE IF NOT EXISTS people.positions
(
  id          smallserial PRIMARY KEY,
  dept_id     int references people.departments             NOT NULL,
  title       varchar UNIQUE                                NOT NULL, -- название должности
  description varchar                                       NOT NULL, -- описание
  parent_id   int REFERENCES people.positions default NULL,           -- руководитель
  deleted     bool                            default false NOT NULL
);

-- Пользователи
CREATE TABLE IF NOT EXISTS people.users
(
  id          serial PRIMARY KEY,
  deleted     bool                                  default false      NOT NULL,
  firstname   varchar(128)                                             NOT NULL, -- Имя
  surname     varchar(128)                          default ''         NOT NULL, -- Фамилия
  status      people.statuses_enum                  default 'inactive' NOT NULL,
  roles       people.roles_enum[]                   default NULL,
  note        text                                  default NULL,                -- примечание
  avatar_id   int check ( avatar_id > 0)            default NULL,
  expired_at  timestamp                             default NULL,                -- УЗ активна до указанного времени
  updated_at  timestamp                             default now()      NOT NULL,

  login       varchar(128)
    constraint users_login_unique_constraint UNIQUE default NULL,
  passwd_hash varchar(40)                           default NULL,                -- хэш пароля
  props       jsonb                                 default NULL                 -- данные пользователя
);

-- Сотрудники компании
CREATE TABLE IF NOT EXISTS people.employees
(
  user_id     int PRIMARY KEY REFERENCES people.users NOT NULL,
  position_id int REFERENCES people.positions         NOT NULL,
  dept_id     int REFERENCES people.departments       NOT NULL,
  updated_at  timestamp default now()                 NOT NULL,
  worked_at   timestamp default now()                 NOT NULL, -- дата начала работы
  fired_at    timestamp default NULL,                           -- дата увольнения (последний день работы)
  constraint users_fired_before_work_constraint CHECK (employees.worked_at < employees.fired_at)
);

-- Наборы документов пользователей
CREATE TABLE IF NOT EXISTS people.user_media_boxes
(
  user_id     int PRIMARY KEY REFERENCES people.users NOT NULL,
  position_id int REFERENCES people.positions         NOT NULL,
  worked_at   timestamp default now()                 NOT NULL, -- дата начала работы
  fired_at    timestamp default NULL                  NULL      -- дата увольнения (последний день работы)
);

-- Руководители отделов
CREATE TABLE IF NOT EXISTS people.head_of_dept
(
  dept_id     smallint references people.departments NOT NULL,
  employee_id int references people.employees        NOT NULL
);


--------------------------------------------------------------------------------
-------------------------------   init data   ----------------------------------
--------------------------------------------------------------------------------

insert into people.departments (id, title, description, parent_id)
values (1, 'Дирекция', 'управление всеми процессами в компании', null),
       (2, 'Доставки/перевозки', 'работа с водителями/грузчиками', 1),
       (3, 'Кадры', 'набор персонала', 1),
       (4, 'АТИ', 'междугородние перевозки', 1),
       (5, 'Транспортный отдел', 'работа с транспортом', 1),
       (6, 'Сборка', '', 1)
on conflict do nothing;

insert into people.positions (id, dept_id, title, description, parent_id)
values (1, 1, 'Генеральный директор', 'описание должности 1', null),
       (2, 2, 'Заместитель Гендира', 'описание должности 2', null),
       (2, 1, 'Администратор ИТ', 'описание должности 2', null),
       (5, 2, 'Руководитель отдела транспорта', '', null),
       (7, 4, 'Руководитель АТИ', '', null),
       (2, 2, 'водитель-экспедитор', 'описание должности 2', null),
       (3, 2, 'экспедитор', 'описание должности 2', null),
       (4, 2, 'грузчик-экспедитор', 'описание должности 2', null),
       (8, 2, 'Оператор', '', null),
       (9, 2, 'грузчик', '', null)
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
