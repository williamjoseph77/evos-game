CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS roles
(
  id SERIAL NOT NULL,
  guid UUID NOT NULL DEFAULT uuid_generate_v1(),
  name TEXT NOT NULL,
  PRIMARY KEY (id)
);

CREATE INDEX idx_roles_guid ON roles (guid);

INSERT INTO roles (name) VALUES ('Wizard'), ('Elf');