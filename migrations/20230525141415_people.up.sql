CREATE TYPE people.roles_enum AS ENUM ('super-admin', 'admin', 'audit', 'guest', 'operator', 'boss');
CREATE TYPE people.statuses_enum AS ENUM ('active', 'inactive', 'blocked', 'registration');
CREATE TYPE people.additional_field_kind_enum AS ENUM (
  'email-main', 'email-personal', 'tel-work', 'tel-personal', 'note');

-- Отделы
CREATE TABLE IF NOT EXISTS people.departments
(
  id          smallserial PRIMARY KEY,
  deleted     bool                              default false NOT NULL,
  title       varchar UNIQUE                                  NOT NULL, -- название департамента
  description varchar                                         NOT NULL, -- описание
  parent_id   int REFERENCES people.departments default NULL            -- руководитель
);

-- Должности
CREATE TABLE IF NOT EXISTS people.positions
(
  id          smallserial PRIMARY KEY,
  deleted     bool default false NOT NULL,
  title       varchar UNIQUE     NOT NULL, -- название должности
  description varchar            NOT NULL  -- описание
);

-- Пользователи
CREATE TABLE IF NOT EXISTS people.users
(
  id          serial PRIMARY KEY,
  deleted     bool                                  default false      NOT NULL,
  updated_at  timestamp                             default now()      NOT NULL,
  firstname   varchar(128)                                             NOT NULL, -- Имя
  surname     varchar(128)                                             NOT NULL, -- Фамилия
  patronymic  varchar(128)                          default ''         NOT NULL, -- Отчество
  status      people.statuses_enum                  default 'inactive' NOT NULL,
  roles       people.roles_enum[]                   default NULL,
  note        text                                  default NULL,                -- примечание
  avatar_id   int check ( avatar_id > 0)            default NULL,
  expired_at  timestamp                             default NULL,                -- УЗ активна до указанного времени

  login       varchar(128)
    constraint users_login_unique_constraint UNIQUE default NULL,
  passwd_hash varchar(40)                           default NULL,                -- хэш пароля
  props       jsonb                                 default NULL                 -- данные пользователя
);

-- Сотрудники компании
CREATE TABLE IF NOT EXISTS people.employees
(
  user_id     int PRIMARY KEY REFERENCES people.users NOT NULL,
  updated_at  timestamp default now()                 NOT NULL,
  position_id int REFERENCES people.positions         NOT NULL,
  dept_id     int REFERENCES people.departments       NOT NULL,
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
CREATE TABLE IF NOT EXISTS people.heads_of_depts
(
  dept_id     smallint references people.departments NOT NULL,
  employee_id int references people.employees        NOT NULL,
  kind        smallint default 0                     NOT NULL -- роль в руководстве (рук/зам)
);


--------------------------------------------------------------------------------
-------------------------------   init data   ----------------------------------
--------------------------------------------------------------------------------

insert into people.departments (id, title, description, parent_id)
values (1, 'Дирекция', 'управление всеми процессами в компании', null),
       (2, 'Кадры', 'набор персонала', 1),
       (3, 'Транспортный отдел', 'работа с транспортом', 1),
       (4, 'Клиентский отдел', 'работа с клиентами (организациями) по доставке', 1),
       (5, 'Оперативный центр', 'работа с водителями/грузчиками', 1),
       (6, 'АТИ', 'Заказы с площадки ати', 1),
       (7, 'Сборка', '', 1)
on conflict do nothing;

insert into people.positions (id, title, description)
values (1, 'Генеральный директор', 'описание должности 1'),
       (2, 'Заместитель Гендира', 'описание должности 2'),
       (3, 'Администратор ИТ', 'инфраструктура/по'),
       (4, 'Руководитель отдела транспорта', ''),
       (5, 'Руководитель АТИ', ''),
       (6, 'Оператор оперативного центра', 'разруливание проблем сборки/доставки'),
       (7, 'водитель-экспедитор', 'описание должности 2'),
       (8, 'экспедитор', 'описание должности 2'),
       (9, 'грузчик-экспедитор', 'описание должности 2'),
       (10, 'Оператор', ''),
       (11, 'грузчик', '')
on conflict do nothing;

