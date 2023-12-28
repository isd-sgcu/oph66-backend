CREATE TABLE event_registrations (
    "user_id" INT NOT NULL,
    "schedule_id" INT NOT NULL,
    PRIMARY KEY ("user_id", "schedule_id")
);

ALTER TABLE "schedules" ADD COLUMN "current_attendee" INT NOT NULL DEFAULT 0;

ALTER SEQUENCE "users_id_seq" RESTART 10000;
