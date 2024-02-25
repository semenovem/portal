-- ------------------------------------------------------------------------
-- фото / видео
-- ------------------------------------------------------------------------

-- типы файлов
CREATE TYPE media.file_kind AS ENUM ('jpg', 'png', 'webp', 'pdf', 'video');

-- аватарки пользователей
CREATE TABLE IF NOT EXISTS media.avatars
(
  id      varchar(15) PRIMARY KEY
);

-- предварительно загруженные файлы
CREATE TABLE IF NOT EXISTS media.upload_files
(
  id      serial PRIMARY KEY,
  note    text default NULL,
  deleted bool default false NOT NULL,
  kind    media.file_kind    NOT NULL,
  s3_path varchar            NOT NULL -- путь сохранения в S3
);

-- предварительно наборы картинок
CREATE TABLE IF NOT EXISTS media.upload_boxes
(
  id              SERIAL PRIMARY KEY,
  note            TEXT default NULL,
  upload_file_ids INT[]              NOT NULL, -- список файлов, входящих в набор
  deleted         BOOL default false NOT NULL
);

-- файлы
CREATE TABLE IF NOT EXISTS media.files
(
  id          SERIAL PRIMARY KEY,
  hash_sha256 VARCHAR            NOT NULL,
  note        TEXT DEFAULT NULL,
  deleted     BOOL DEFAULT FALSE NOT NULL,
  kind        media.file_kind    NOT NULL

--  хранение в S3
);

-- наборы картинок
CREATE TABLE IF NOT EXISTS media.boxes
(
  id       SERIAL PRIMARY KEY,
  note     TEXT DEFAULT NULL,
  file_ids INT[]              NOT NULL, -- список файлов, входящих в набор
  deleted  BOOL DEFAULT FALSE NOT NULL
);

-- набор фото автомобиля
CREATE TABLE IF NOT EXISTS media.vehicle_images
(
  id                    SERIAL PRIMARY KEY,
  cabin_image_box_id    INT REFERENCES media.boxes, -- список картинок, входящих в набор
  front_image_box_id    INT REFERENCES media.boxes,
  left_image_box_id     INT REFERENCES media.boxes,
  right_image_box_id    INT REFERENCES media.boxes,
  back_image_box_id     INT REFERENCES media.boxes,
  odometer_image_box_id INT REFERENCES media.boxes,
  other_image_box_id    INT REFERENCES media.boxes
);
