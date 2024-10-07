CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    age INTEGER NOT NULL
);

INSERT INTO students (name, email, age) VALUES ('Saunderson McDuall', 'smcduall0@virginia.edu', 45);
INSERT INTO students (name, email, age) VALUES ('Shepherd Crandon', 'scrandon1@spiegel.de', 61);
INSERT INTO students (name, email, age) VALUES ('Dolorita MacIver', 'dmaciver2@ezinearticles.com', 44);
INSERT INTO students (name, email, age) VALUES ('Nance Giddens', 'ngiddens3@facebook.com', 29);
INSERT INTO students (name, email, age) VALUES ('Mill I''anson', 'mianson4@ehow.com', 25);

ALTER TABLE students REPLICA IDENTITY FULL;

CREATE ROLE debezium_user WITH REPLICATION LOGIN PASSWORD 'password';
CREATE PUBLICATION students_publication FOR TABLE students;
GRANT SELECT ON students TO debezium_user;
