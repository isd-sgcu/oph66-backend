CREATE TABLE attendee_checkins (
    "id" SERIAL NOT NULL,
    "user_id" INT NOT NULL,
    "faculty_code" VARCHAR(10) NOT NULL,
    "department_code" VARCHAR(10) NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_checkins_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    CONSTRAINT "fk_checkins_faculties" FOREIGN KEY ("faculty_code") REFERENCES "faculties" ("code"),
    CONSTRAINT "fk_checkins_departments" FOREIGN KEY ("department_code", "faculty_code") REFERENCES "departments" ("code", "faculty_code"),
    CONSTRAINT "unique_checkin" UNIQUE ("user_id", "faculty_code", "department_code")
);
