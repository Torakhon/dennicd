CREATE TABLE "archive"(
                          "id" SERIAL PRIMARY KEY NOT NULL,
                          "doctor_availability_id" INTEGER NOT NULL,
                          "start_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                          "patient_problem" TEXT NOT NULL,
                          "end_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                          "status" VARCHAR(255) NOT NULL CHECK ("status" IN ('attended', 'cancelled', 'no_show')),
                          "payment_type" VARCHAR(255) NOT NULL CHECK ("payment_type" IN ('cash', 'card', 'insurance')),
                          "payment_amount" DOUBLE PRECISION NOT NULL,
                          "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE,
                          "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);

INSERT INTO archive (start_time, patient_problem, end_time, status, payment_type, payment_amount, created_at)
VALUES
    ('09:00:12', 'Fever', '10:00:12', 'attended', 'cash', 50.00),
    ('10:30:12', 'Headache', '11:15:12', 'cancelled', 'card', 30.00),
    ('13:00:12', 'Sore throat', '13:30:12', 'attended', 'insurance', 80.00),
    ('15:30:12', 'Stomach ache', '16:00:12', 'attended', 'cash', 60.00),
    ('11:00:12', 'Back pain', '11:30:12', 'no_show', 'card', 0.00),
    ('14:30:12', 'Allergy', '15:00:12', 'attended', 'cash', 70.00),
    ('16:30:12', 'Cough', '17:00:12', 'attended', 'insurance', 90.00),
    ('09:30:12', 'Fever', '10:30:12', 'attended', 'cash', 50.00),
    ('12:00:12', 'Sore throat', '12:30:12', 'cancelled', 'card', 0.00),
    ('14:00:12', 'Headache', '14:45:12', 'attended', 'cash', 40.00);