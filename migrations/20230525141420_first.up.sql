-- ------------------------------------------------------------------------
-- фото / видео
-- ------------------------------------------------------------------------
-- картинки
CREATE TABLE IF NOT EXISTS core.images
(
  id          serial PRIMARY KEY NOT NULL,
  hash_sha256 varchar            NOT NULL, -- хеш оригинала картинки
  preview     varchar            NOT NULL, -- превью
  preview2    varchar            NOT NULL, -- превью2
  note        text
);

-- наборы картинок
CREATE TABLE IF NOT EXISTS core.image_boxes
(
  id        serial PRIMARY KEY                     NOT NULL,
  image_ids smallint[] REFERENCES core.images (id) NOT NULL -- список картинок, входящих в набор
);

-- набор фото автомобиля
CREATE TABLE IF NOT EXISTS core.vehicle_images
(
  id                    serial PRIMARY KEY                       NOT NULL,
  cabin_image_box_id    integer REFERENCES core.image_boxes (id) NULL, -- список картинок, входящих в набор
  front_image_box_id    integer REFERENCES core.image_boxes (id) NULL,
  left_image_box_id     integer REFERENCES core.image_boxes (id) NULL,
  right_image_box_id    integer REFERENCES core.image_boxes (id) NULL,
  back_image_box_id     integer REFERENCES core.image_boxes (id) NULL,
  odometer_image_box_id integer REFERENCES core.image_boxes (id) NULL,
  other_image_box_id    integer REFERENCES core.image_boxes (id) NULL
);



-- ------------------------------------------------------------------------
-- Сотрудники
CREATE TABLE IF NOT EXISTS core.employees
(
  id              serial PRIMARY KEY NOT NULL,
  firstname       varchar            NOT NULL,                  -- Имя
  surname         varchar            NOT NULL,                  -- Фамилия
  email           varchar            NOT NULL,
  passwd_hash     varchar(64)        NOT NULL,                  -- хэш пароля
  created_at      timestamp GENERATED ALWAYS AS (now()) STORED, -- дата создания записи
  updated_at      timestamp          NOT NULL,                  -- обновление данных
  deleted_at      timestamp          NULL,                      -- обновление данных
  created_user_id serial REFERENCES core.employees (id),        -- кто создал запись
  note            text    default '' NOT NULL,                  -- примечание
  avatar          varchar default '' NOT NULL,
  position        varchar default ''
);

-- Подрядчики
CREATE TABLE IF NOT EXISTS core.contractors
(
  id              serial PRIMARY KEY NOT NULL,
  name            varchar            NOT NULL,                  -- Название
  note            text default ''    NOT NULL,                  -- примечание
  created_at      timestamp GENERATED ALWAYS AS (now()) STORED, -- дата создания записи
  updated_at      timestamp          NOT NULL,                  -- обновление данных
  deleted_at      timestamp          NULL,                      -- запись удалена
  created_user_id serial REFERENCES core.employees (id)
);


-- ------------------------------------------------------------------------
-- автотранспорт
-- ------------------------------------------------------------------------

-- Авто-транспортные средства
CREATE TABLE IF NOT EXISTS core.vehicles
(
  id              serial PRIMARY KEY NOT NULL,
  brand           varchar            NOT NULL,                  -- бренд
  model           varchar            NOT NULL,                  -- модель
  date_of_issue   date               NOT NULL,                  -- дата выпуска
  date_start_use  date               NOT NULL,                  -- дата начала эксплуатации
  date_end_use    date               NULL,                      -- дата завершения эксплуатации
-- todo использовать enum
  status          varchar(10)        NOT NULL,                  -- статус
  note            text default ''    NOT NULL,                  -- примечание
  number          varchar            NOT NULL,                  -- гос.номер автомобиля
  vin             varchar            NOT NULL,                  -- VIN номер автомобиля
  tonnage         smallint           NOT NULL,                  -- грузоподъемность в килограммах
  created_at      timestamp GENERATED ALWAYS AS (now()) STORED, -- дата создания записи
  updated_at      timestamp,                                    -- обновление данных
  created_user_id serial references core.employees (id)         -- кто создал запись
);


-- Техническое обслуживание (в сервисе или нашими силами)
CREATE TABLE IF NOT EXISTS core.maintenances
(
  id                 serial PRIMARY KEY                      NOT NULL,
  created_at         timestamp GENERATED ALWAYS AS (now()) STORED,     -- дата создания записи
  note               text default ''                         NOT NULL, -- примечание
  contractor_id      serial REFERENCES core.contractors (id) NULL,
  employee_id        serial REFERENCES core.employees (id)   NOT NULL,
  start_date         timestamp,                                        -- дата начала работ по обслуживанию
  end_date           timestamp,                                        -- дата завершения работ по обслуживанию
  employee_master_id serial REFERENCES core.employees (id)             -- ответственный за обслуживание сотрудник
);


-- Список работ по обслуживанию автомобиля + TODO таблицу с материалами для обслуживания
CREATE TABLE IF NOT EXISTS core.maintenance_items
(
  id         serial PRIMARY KEY NOT NULL,
  name       varchar            NOT NULL,                 -- наименование
  note       text default ''    NOT NULL,                 -- примечание
  type       varchar            NOT NULL,                 -- тип (единица измерения и тд)
  created_at timestamp GENERATED ALWAYS AS (now()) STORED -- дата создания записи
);

CREATE TABLE IF NOT EXISTS core.maintenance_do_items
(
  maintenance_id      integer REFERENCES core.maintenances      NOT NULL,
  maintenance_item_id integer REFERENCES core.maintenance_items NOT NULL,
  cost                real default 0                            NOT NULL -- стоимость
);
