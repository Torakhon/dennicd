CREATE TABLE "patients"(
                           "id" UUID PRIMARY KEY NOT NULL,
                           "first_name" VARCHAR(50) NOT NULL,
                           "last_name" VARCHAR(50) NOT NULL,
                           "birth_date" DATE NOT NULL,
                           "gender" VARCHAR(255) NOT NULL CHECK ("gender" IN ('male', 'female', 'other')),
                           "blood_group" VARCHAR(255) NOT NULL CHECK ("blood_group" IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-')),
                           "phone_number" VARCHAR(15) NOT NULL,
                           "address" VARCHAR(250) NOT NULL,
                           "city" VARCHAR(50) NOT NULL,
                           "country" VARCHAR(50) NOT NULL,
                           "patient_problem" TEXT NOT NULL,
                           "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE ,
                           "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);

INSERT INTO patients (id, first_name, last_name, birth_date, gender, blood_group, phone_number, address, city, country, patient_problem, created_at)
VALUES
    ( 'John', 'Doe', '1990-05-15', 'male', 'A+', '1234567890', '123 Main St', 'New York', 'USA', 'Fever'),
    ( 'Jane', 'Smith', '1985-09-20', 'female', 'B-', '0987654321', '456 Elm St', 'Los Angeles', 'USA', 'Headache'),
    ( 'Michael', 'Johnson', '1978-12-10', 'male', 'O-', '4567890123', '789 Oak St', 'Chicago', 'USA', 'Sore throat'),
    ( 'Emily', 'Brown', '1995-03-25', 'female', 'AB+', '7890123456', '987 Pine St', 'Houston', 'USA', 'Stomach ache'),
    ( 'William', 'Wilson', '1982-07-08', 'male', 'A-', '2345678901', '654 Cedar St', 'Phoenix', 'USA', 'Back pain'),
    ('Emma', 'Taylor', '1998-09-30', 'female', 'B+', '8901234567', '321 Birch St', 'Philadelphia', 'USA', 'Allergy'),
    ( 'Christopher', 'Martinez', '1975-11-05', 'male', 'O+', '5678901234', '210 Maple St', 'San Antonio', 'USA', 'Cough'),
    ( 'Sophia', 'Garcia', '1992-01-12', 'female', 'AB-', '6789012345', '543 Spruce St', 'San Diego', 'USA', 'Fever'),
    ( 'Daniel', 'Lopez', '1987-08-18', 'male', 'B-', '7890123456', '876 Oak St', 'Dallas', 'USA', 'Sore throat'),
    ( 'Olivia', 'Perez', '1993-06-22', 'female', 'A+', '8901234567', '123 Elm St', 'San Francisco', 'USA', 'Headache');