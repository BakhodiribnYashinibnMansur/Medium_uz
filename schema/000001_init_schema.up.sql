
CREATE TABLE  IF NOT EXISTS users (
   id            SERIAL PRIMARY KEY     NOT NULL ,
   email VARCHAR(255) NOT NULL ,
   password_hash VARCHAR(255) NOT NULL ,
   firstname VARCHAR(255) NOT NULL UNIQUE,
   secondname VARCHAR(255) NOT NULL ,
   city VARCHAR(255) NOT NULL ,
   isverified BOOLEAN DEFAULT FALSE,
   verification_date VARCHAR(255) NULL ,
   account_image_path VARCHAR(255) NULL,
   phone VARCHAR(255) NOT NULL,
   rating DOUBLE PRECISION DEFAULT 0.0,
   post_views INTEGER DEFAULT 0,
   issuperuser BOOLEAN DEFAULT FALSE
);
