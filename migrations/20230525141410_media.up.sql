-- ------------------------------------------------------------------------
-- фото / видео
-- ------------------------------------------------------------------------

-- типы файлов
CREATE TYPE media.file_kind AS ENUM ('jpg', 'png', 'webp', 'pdf', 'video');

-- картинки
CREATE TABLE IF NOT EXISTS media.files
(
  id          serial PRIMARY KEY,
  hash_sha256 varchar            NOT NULL, -- хеш оригинала картинки TODO указать точный размер
  note        text default NULL,
  deleted     bool default false NOT NULL,
  kind        media.file_kind    NOT NULL

--  хранение в S3
);

-- наборы картинок
CREATE TABLE IF NOT EXISTS media.files_boxes
(
  id       serial PRIMARY KEY,
  note     text default NULL,
  file_ids int[]              NOT NULL, -- список файлов, входящих в набор
  deleted  bool default false NOT NULL
);

-- набор фото автомобиля
CREATE TABLE IF NOT EXISTS media.vehicle_images
(
  id                    serial PRIMARY KEY,
  cabin_image_box_id    int REFERENCES media.files_boxes, -- список картинок, входящих в набор
  front_image_box_id    int REFERENCES media.files_boxes,
  left_image_box_id     int REFERENCES media.files_boxes,
  right_image_box_id    int REFERENCES media.files_boxes,
  back_image_box_id     int REFERENCES media.files_boxes,
  odometer_image_box_id int REFERENCES media.files_boxes,
  other_image_box_id    int REFERENCES media.files_boxes
);

-- аватарки
CREATE TABLE IF NOT EXISTS media.avatars
(
  id   serial PRIMARY KEY,
  path varchar NOT NULL
--  хранение в S3
);
