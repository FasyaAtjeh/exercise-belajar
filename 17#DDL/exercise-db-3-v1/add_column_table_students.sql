-- TODO: answer here
ALTER TABLE students
ADD date_of_birth DATE NOT NULL,
ADD street VARCHAR(255),
ADD city VARCHAR(100),
ADD province VARCHAR(100),
ADD country VARCHAR(100),
ADD postal_code VARCHAR(50);