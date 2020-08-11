CREATE TABLE users (
   id UUID PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   access_token VARCHAR(2048),
   refesh_token VARCHAR(2048),
   active BOOLEAN NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);