ALTER TABLE attendee_checkins RENAME TO attendee_faculty_checkins;

CREATE TABLE attendee_central_checkins (
    "id" SERIAL NOT NULL,
    "user_id" INT NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_checkins_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
    CONSTRAINT "unique_attendee_central_checkins_user_id" UNIQUE ("user_id")
);
