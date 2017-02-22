CREATE SCHEMA animals;

CREATE SEQUENCE people.profile_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;

CREATE TABLE people.profile (
    id   integer NOT NULL DEFAULT nextval('people.profile_seq'::regclass),
    eid integer,
    name text,
    position text
);

ALTER TABLE ONLY people.profile
    ADD CONSTRAINT profile_pkey PRIMARY KEY (id);

