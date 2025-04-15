CREATE TABLE IF NOT EXISTS users (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  username text UNIQUE NOT NULL,
  password bytea NOT NULL,
  phone text UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS couples (
  id   BIGSERIAL PRIMARY KEY,
  user_id BIGSERIAL REFERENCES users(id),
  partner_id BIGSERIAL REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS wishes (
  id   BIGSERIAL PRIMARY KEY,
  title text NOT NULL,
  description text, 
  url text,
  price text,
  completed boolean DEFAULT(false),
  couple_id BIGINT REFERENCES couples(id),
  created_at TIMESTAMP DEFAULT NOW()
);