CREATE TABLE IF NOT EXISTS characters
(
  id SERIAL NOT NULL,
  guid UUID NOT NULL DEFAULT uuid_generate_v1(),
  name TEXT NOT NULL,
  role_id INT NOT NULL,
  power INT NOT NULL,
  wealth FLOAT NOT NULL,
  PRIMARY KEY (id)
);

CREATE UNIQUE INDEX name_unique_idx ON characters (LOWER(name));
ALTER TABLE characters ADD CONSTRAINT FK_characters_role_id FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE RESTRICT ON UPDATE CASCADE;