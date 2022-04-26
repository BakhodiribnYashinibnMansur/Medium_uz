
CREATE TABLE IF NOT EXISTS users (
   id            SERIAL PRIMARY KEY     NOT NULL ,
   email VARCHAR (255) NOT NULL ,
   password_hash VARCHAR (255) NOT NULL ,
   firstname VARCHAR (255) NOT NULL UNIQUE,
   secondname VARCHAR (255) NOT NULL ,
   city VARCHAR (255) NOT NULL ,
   is_verified BOOLEAN DEFAULT FALSE,
   verification_date TIMESTAMP NULL,
   account_image_path VARCHAR (255) NULL,
   phone VARCHAR (255) NOT NULL,
   rating DOUBLE PRECISION DEFAULT 0.0,
   post_views INTEGER DEFAULT 0,
   is_super_user BOOLEAN DEFAULT FALSE,
   created_at TIMESTAMP DEFAULT (NOW ()),
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP NULL
);
