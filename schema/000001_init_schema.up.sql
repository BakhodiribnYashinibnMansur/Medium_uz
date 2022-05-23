
CREATE TABLE IF NOT EXISTS users (
   id  SERIAL      NOT NULL UNIQUE ,
   firstname VARCHAR(253) NOT NULL ,
   secondname VARCHAR (255) NOT NULL ,
   email VARCHAR(255) NOT NULL UNIQUE ,
   nickname VARCHAR(255) NOT NULL UNIQUE ,
   password_hash VARCHAR(255) NOT NULL ,
   interests TEXT [] NULL,
   bio VARCHAR(255) NULL,
   city VARCHAR (255) NOT NULL ,
   is_verified BOOLEAN DEFAULT FALSE,
   verification_date TIMESTAMP NULL,
   account_image_path VARCHAR (255) NULL,
   phone VARCHAR (255) NOT NULL,
   rating DOUBLE PRECISION DEFAULT 0.0,
   post_views_count INTEGER DEFAULT 0,
   follower_count INTEGER DEFAULT 0,
   following_count INTEGER DEFAULT 0,
   like_count INTEGER DEFAULT 0,
   is_super_user BOOLEAN DEFAULT FALSE,
   created_at TIMESTAMP DEFAULT (NOW()),
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS post (
id SERIAL NOT NULL  UNIQUE,
post_title VARCHAR(255) NOT NULL,
post_image_path VARCHAR(255)   NULL,
post_body TEXT NOT NULL,
post_views_count INTEGER NOT NULL DEFAULT 0,
post_like_count INTEGER NOT NULL DEFAULT 0,
post_rated DOUBLE PRECISION NULL DEFAULT 0.0,
post_vote INTEGER NOT NULL DEFAULT 0,
post_tags TEXT[] NOT NULL,
post_date TIMESTAMP NOT NULL DEFAULT (NOW()),
is_new BOOLEAN  NULL DEFAULT FALSE,
is_top_read BOOLEAN  NULL DEFAULT FALSE,
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS post_user (
id  SERIAL NOT NULL ,
post_author_id INT  REFERENCES users (id) ON DELETE CASCADE NOT NULL,
post_id INT     REFERENCES post (id) ON DELETE CASCADE NOT NULL,
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS liked_post (
id  SERIAL NOT NULL UNIQUE,
reader_id INTEGER   REFERENCES users(id) ON DELETE CASCADE NOT NULL,
post_id INTEGER REFERENCES post(id) ON DELETE CASCADE NOT NULL,
like_data TIMESTAMP DEFAULT (NOW()),
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS followers (
id  SERIAL NOT NULL UNIQUE,
reader_id INTEGER   REFERENCES users(id) ON DELETE CASCADE NOT NULL,
poster_id INTEGER   REFERENCES users(id) ON DELETE CASCADE NOT NULL,
following_data  TIMESTAMP DEFAULT (NOW()),
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS viewed_posts (
id  SERIAL NOT NULL UNIQUE,
reader_id INTEGER   REFERENCES users(id) ON DELETE CASCADE NOT NULL,
post_id INTEGER   REFERENCES post(id) ON DELETE CASCADE NOT NULL,
view_date TIMESTAMP DEFAULT (NOW()),
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS review_post (
id  SERIAL NOT NULL UNIQUE,
reader_id INTEGER   REFERENCES users(id) ON DELETE CASCADE NOT NULL,
post_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
review_data TIMESTAMP DEFAULT (NOW()),
commits VARCHAR(255) NOT NULL,
created_at TIMESTAMP  DEFAULT (NOW()),
updated_at TIMESTAMP NULL,
deleted_at TIMESTAMP NULL
);
