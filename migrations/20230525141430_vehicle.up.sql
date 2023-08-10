-- Подрядчики технического обслуживания
CREATE TABLE IF NOT EXISTS vehicle.contractors
(
  id         serial PRIMARY KEY,
  name       varchar                 NOT NULL,
  note       text      default ''    NOT NULL,
  created_at timestamp default now() NOT NULL,
  updated_at timestamp default now() NOT NULL,
  deleted_at timestamp default NULL  NULL
);


-- ------------------------------------------------------------------------
-- автотранспорт
-- ------------------------------------------------------------------------

-- Авто-транспортные средства
CREATE TABLE IF NOT EXISTS vehicle.vehicles
(
  id              serial PRIMARY KEY,
  brand           varchar                 NOT NULL, -- бренд
  model           varchar                 NOT NULL, -- модель
  date_of_issue   date                    NOT NULL, -- дата выпуска
  date_start_use  date                    NOT NULL, -- дата начала эксплуатации
  date_end_use    date                    NULL,     -- дата завершения эксплуатации
-- todo использовать enum
  status          varchar(10)             NOT NULL, -- статус
  note            text      default ''    NOT NULL, -- примечание
  number          varchar                 NOT NULL, -- гос.номер автомобиля
  vin             varchar                 NOT NULL, -- VIN номер автомобиля
  tonnage         smallint                NOT NULL, -- грузоподъемность в килограммах
  created_at      timestamp default now() NOT NULL, -- дата создания записи
  updated_at      timestamp default now() NOT NULL, -- обновление данных
  created_user_id int REFERENCES people.users       -- кто создал запись
);


-- Техническое обслуживание (в сервисе или нашими силами)
CREATE TABLE IF NOT EXISTS vehicle.maintenances
(
  id             serial PRIMARY KEY,
  vehicle_id     integer REFERENCES vehicle.vehicles,              -- дата создания записи
  created_at     timestamp default now()                 NOT NULL, -- дата создания записи
  note           text      default ''                    NOT NULL, -- примечание
  contractor_id  int REFERENCES vehicle.contractors (id) NULL,
  user_id        int REFERENCES people.users             NOT NULL,
  start_date     timestamp,                                        -- дата начала работ по обслуживанию
  end_date       timestamp,                                        -- дата завершения работ по обслуживанию
  user_master_id int REFERENCES people.users                       -- ответственный за обслуживание сотрудник
);


-- Список работ по обслуживанию автомобиля + TODO таблицу с материалами для обслуживания
CREATE TABLE IF NOT EXISTS vehicle.maintenance_items
(
  id         serial PRIMARY KEY,
  name       varchar                 NOT NULL, -- наименование
  note       text      default ''    NOT NULL, -- примечание
  type       varchar                 NOT NULL, -- тип (единица измерения и тд)
  created_at timestamp default now() NOT NULL  -- дата создания записи
);


CREATE TABLE IF NOT EXISTS vehicle.maintenance_use_items
(
  maintenance_id      integer REFERENCES vehicle.maintenances      NOT NULL,
  maintenance_item_id integer REFERENCES vehicle.maintenance_items NOT NULL,
  cost                real default 0                               NOT NULL -- стоимость
);
