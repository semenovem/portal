-- ------------------------------------------------------------------------
-- фото / видео
-- ------------------------------------------------------------------------
-- картинки
CREATE TABLE IF NOT EXISTS media.images
(
  id          serial PRIMARY KEY,
  hash_sha256 varchar NOT NULL, -- хеш оригинала картинки
  preview     varchar NOT NULL, -- превью
  preview2    varchar NOT NULL, -- превью2
  note        text
--  хранение в S3
);

-- наборы картинок
CREATE TABLE IF NOT EXISTS media.image_boxes
(
  id        serial PRIMARY KEY,
  image_ids smallint[] NOT NULL -- список картинок, входящих в набор
);

-- набор фото автомобиля
CREATE TABLE IF NOT EXISTS media.vehicle_images
(
  id                    serial PRIMARY KEY,
  cabin_image_box_id    int REFERENCES media.image_boxes NULL, -- список картинок, входящих в набор
  front_image_box_id    int REFERENCES media.image_boxes NULL,
  left_image_box_id     int REFERENCES media.image_boxes NULL,
  right_image_box_id    int REFERENCES media.image_boxes NULL,
  back_image_box_id     int REFERENCES media.image_boxes NULL,
  odometer_image_box_id int REFERENCES media.image_boxes NULL,
  other_image_box_id    int REFERENCES media.image_boxes NULL
);

-- аватарки
CREATE TABLE IF NOT EXISTS media.avatars
(
  id   serial PRIMARY KEY,
  path varchar NOT NULL
--  хранение в S3
);
